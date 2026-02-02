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
	fmt.Println("version: ", a.Version())

	fmt.Println("-----------生成v7：")
	b, err := uuid.NewV7()
	if err != nil {
		fmt.Println("生成v7失败：", err)
		return
	}
	fmt.Println(b.String())
	fmt.Println(b.URN())
	fmt.Println("version: ", b.Version())
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

func Test11237(t *testing.T) {
	var updated bool
	fmt.Println(updated)
}

func TestTimeDuration(t *testing.T) {
	start := time.Now()
	time.Sleep(2*time.Second + 500*time.Millisecond) // 模拟处理时间
	end := time.Now()

	duration := end.Sub(start)

	fmt.Printf("Duration: %d\n", duration) //%d可以打印数字
	fmt.Printf("%0.2f seconds\n", duration.Seconds())

	s := duration.String()
	fmt.Println(s)
}

func isPrime(n int) (bool, int) {
	if n <= 1 {
		return false, 0
	}
	if n <= 3 {
		return true, 0 // 2 和 3 是质数
	}
	if n%2 == 0 {
		return false, 2 // 排除能被 2 整除的
	}
	if n%3 == 0 {
		return false, 3 // 排除能被 3 整除的
	}
	// 检查形如 6k ± 1 的因数
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 {
			return false, i
		}

		if n%(i+2) == 0 {
			return false, i + 2
		}
	}

	return true, 0
}

func TestIsPrime(t *testing.T) {
	n := 259
	r, f := isPrime(259)
	fmt.Printf("%v的判断结果是：%v , 它的因数是：%v\n", n, r, f)
	//
}

// 测试一下函数返回的是接口类型的时候，还能不能用类型断言转换成具体类型
type IA interface {
	F() int
}
type type1 struct{}
type type2 struct{}

var _ IA = (*type1)(nil)
var _ IA = (*type2)(nil)

func (*type1) F() int {
	return 3
}
func (*type2) F() int {
	return 5
}
func newt1() IA {
	return &type1{}
}

func TestInterfaceAndTypeAssertion(t *testing.T) {
	// 测试一下函数返回的是接口类型的时候，还能不能用类型断言转换成具体类型
	a := newt1()
	a, ok := a.(*type2)
	fmt.Printf("对a的断言结果: %v \n", ok)
}

// func TestFailedTest(t *testing.T) {
// 	a := 3
// 	assert.Equal(t, 5, a, "a应该等于5")
// }
