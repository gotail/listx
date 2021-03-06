package listx

import (
	"sync"
)

const (
	halfCap            = 100
	expansionThreshold = 10
	expansionCount     = 1000
)

type List struct {
	mutex sync.RWMutex
	left  uint32
	right uint32
	start uint32
	end   uint32
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

func (list *List) LPop() (p interface{}) {
	if p = list.Data[list.start]; p != nil {
		list.mutex.Lock()
		defer list.mutex.Unlock()

		list.Data[list.start] = nil
		list.start = list.start + 1
		list.expansion()
	}
	return p
}

func (list *List) RPop() (p interface{}) {
	if p = list.Data[list.end-1]; p != nil {
		list.mutex.Lock()
		defer list.mutex.Unlock()

		list.Data[list.end] = nil
		list.end = list.end - 1
		list.expansion()
	}
	return p
}

func (list *List) Del(index uint32) int {
	if list.start+index > list.end || index < 0 || list.start == list.end {
		return -1 // 超出索引
	}

	index = list.start + index
	list.Data = append(list.Data[:index-1], list.Data[index:]...)

	list.end = list.end - 1
	return 1
}

func (list *List) Len() uint32 {
	return list.end - list.start
}

func (list *List) LRange(s, e uint32) []interface{} {
	if s < 0 || s >= e {
		return nil
	}

	s += list.start
	e += list.start

	if s < list.start {
		s = list.start
	}
	if e > list.end {
		e = list.end
	}

	return list.Data[s:e:e]
}

func (list *List) LIndex(index uint32) interface{} {
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
