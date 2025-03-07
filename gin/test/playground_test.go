package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

// golang默认的time.Time时间解析格式是RFC 3339，除非你再自定义type并实现json.Unmarshaler接口
func TestTimeInit(t *testing.T) {

	defer func() {
		rmsg := recover()
		fmt.Println(rmsg)
	}()

	type Timejson struct {
		T time.Time
	}

	tt := Timejson{}

	s := `{"t":"2023-01-02T18:07:05+05:00"}`

	c := []byte(s)
	err := json.Unmarshal(c, &tt)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(c))
	fmt.Println(tt)

}
func TestUuid(t *testing.T) {
	a := uuid.New()
	fmt.Println(a.String())
	fmt.Println(a.URN())
	// fmt.Println(a, a.Version())
}

func TestRangeCreatingVar(t *testing.T) {
	for range 5 {
		var a int
		fmt.Println(a+1, &a)
	}
}

// range切片的时候，在range语句这里就已经确定了要怎么循环了，后面大括号里再改切片都不会影响循环的方式
func TestRangetest(t *testing.T) {

	a := []int{1, 2, 3, 4, 5}

	for i, v := range a {
		fmt.Println(&i, i, "|", &v, v)

		i += 3
		a = a[:len(a)-1]

		fmt.Println(&i, i, "|", &v, v)
		fmt.Println("---")
	}

	fmt.Println(a)
}
