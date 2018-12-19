package listx

import (
	"fmt"
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

	for i := 0; i < 1000; i++ {
		myList.RPush(i)
	}

	fmt.Println(myList.LRange(0, 5))
}
