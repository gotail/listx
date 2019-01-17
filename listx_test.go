package listx

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestList_Del(t *testing.T) {
	count := 800000
	f := make([]interface{}, count)
	v := New()
	for i := 0; i < count; i++ {
		f[i] = v
		v.RPush("lskdjf-" + strconv.Itoa(i))
	}

	start := time.Now()
	for i := 0; i < 5000; i++ {
		//t.Log(v.Len(), v.LIndex(0))
		v.Del(0)
		//t.Log(v.Len(), v.LIndex(0))
	}
	t.Log(v.Len(), v.LIndex(0))
	t.Log(time.Since(start))

}

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

	//start := time.Now()
	//for i := 0; i < 5000000; i++ {
	//	myList.LRange(1, 5)
	//}
	//t.Log(time.Since(start))

	x := myList.LRange(5, 10)
	fmt.Println(x)
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
