package listx

import "sync"

const (
	halfCap            = 100000
	expansionThreshold = 8000
	expansionCount     = 100000
)

type listx struct {
	mutex sync.RWMutex
	left  int
	right int
	start int
	end   int
	data  []interface{}
}

func New() *listx {
	return &listx{
		mutex: sync.RWMutex{},
		left:  halfCap,
		right: halfCap,
		start: halfCap - 1,
		end:   halfCap - 1,
		data:  make([]interface{}, halfCap*2),
	}
}

func (list *listx) LPush(element interface{}) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.start = list.start - 1
	list.data[list.start] = element
	list.expansion()
}

func (list *listx) RPush(element interface{}) {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	list.data[list.end] = element

	list.end = list.end + 1
	list.expansion()
}

func (list *listx) LPop() interface{} {
	v := list.data[list.start]
	if v == nil {
		return nil
	}

	list.mutex.Lock()

	list.data[list.start] = nil
	list.start = list.start + 1
	list.mutex.Unlock()
	list.expansion()
	return v
}

func (list *listx) RPop() interface{} {

	v := list.data[list.end-1]
	if v == nil {
		return nil
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()

	list.data[list.end] = nil
	list.end = list.end - 1

	list.expansion()
	return v
}

func (list *listx) Len() int {
	return list.end - list.start
}

func (list *listx) LRange(s, e int) []interface{} {
	if s < 0 || e > list.start+list.end || s >= e {
		return nil
	}
	return list.data[s+list.start : e+list.start]
}

func (list *listx) LIndex(index int) interface{} {
	if index < 0 || list.start+index > list.end {
		return nil
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()

	return list.data[list.start+index]
}

func (list *listx) expansion() {
	if list.start <= expansionThreshold {
		newLeft := make([]interface{}, expansionCount)
		newLeft = append(newLeft, list.data...)
		list.data = newLeft
		list.start = expansionCount + list.start - 1
		list.end = expansionCount + list.end - 1
		list.left = expansionCount + list.left
	} else if list.right-list.end <= expansionThreshold {
		newRight := make([]interface{}, expansionCount)
		newRight = append(list.data, newRight...)
		list.data = newRight
		list.right = expansionCount + list.right
	}
}
