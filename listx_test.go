package listx

import (
	"testing"
)

func TestX(t *testing.T) {
	myList := New()

	myList.LPush(3)
	t.Log(myList.Len())

	t.Log(myList.LIndex(0))
	t.Log(myList.LPop())
	t.Log(myList.Len())
}

func TestListx_LRange(t *testing.T) {

	myList := New()

	for i := 0; i < 10000000; i++ {
		myList.RPush(i)
	}

	//fmt.Println(myList.LRange(1, 5))
}

func TestNew(t *testing.T) {
	count := 8000

	f := make([]interface{}, count)
	for i := 0; i < count; i++ {
		v := New()
		f[i] = v
		v.LPush("lskdjf")
	}
}
