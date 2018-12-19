package listx

import "sync"

const (
	halfCap            = 100000
	expansionThreshold = 8000
	expansionCount     = 100000
)

type List struct {
	mutex sync.RWMutex
	left  int
	right int
	start int
	end   int
	Data  []interface{}
}

func New() *List {
	return &List{
		mutex: sync.RWMutex{},
		left:  halfCap,
		right: halfCap,
		start: halfCap - 1,
		end:   halfCap - 1,
		Data:  make([]interface{}, halfCap*2),
	}
}

func (list *List) LPush(element interface{}) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	list.start = list.start - 1
	list.Data[list.start] = element
	list.expansion()
}

func (list *List) RPush(element interface{}) {
	list.mutex.Lock()
	defer list.mutex.Unlock()

	list.Data[list.end] = element

	list.end = list.end + 1
	list.expansion()
}

func (list *List) LPop() interface{} {
	v := list.Data[list.start]
	if v == nil {
		return nil
	}

	list.mutex.Lock()

	list.Data[list.start] = nil
	list.start = list.start + 1
	list.mutex.Unlock()
	list.expansion()
	return v
}

func (list *List) RPop() interface{} {

	v := list.Data[list.end-1]
	if v == nil {
		return nil
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()

	list.Data[list.end] = nil
	list.end = list.end - 1

	list.expansion()
	return v
}

func (list *List) Len() int {
	return list.end - list.start
}

func (list *List) LRange(s, e int) []interface{} {
	if s < 0 || e > list.start+list.end || s >= e {
		return nil
	}
	return list.Data[s+list.start : e+list.start]
}

func (list *List) LIndex(index int) interface{} {
	if index < 0 || list.start+index > list.end {
		return nil
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()

	return list.Data[list.start+index]
}

func (list *List) expansion() {
	if list.start <= expansionThreshold {
		newLeft := make([]interface{}, expansionCount)
		newLeft = append(newLeft, list.Data...)
		list.Data = newLeft
		list.start = expansionCount + list.start - 1
		list.end = expansionCount + list.end - 1
		list.left = expansionCount + list.left
	} else if list.right-list.end <= expansionThreshold {
		newRight := make([]interface{}, expansionCount)
		newRight = append(list.Data, newRight...)
		list.Data = newRight
		list.right = expansionCount + list.right
	}
}
