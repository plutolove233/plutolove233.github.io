package main

// 定义一个接口，并分别用值接受者和指针接受者实现该接口，比较这两者的不同
type I interface {
	F()
}

type T1 struct {
}

func (t T1) F() {
	println("T1.F")
}

type T2 struct {
}

func (t *T2) F() {
	println("T2.F")
}

func checkType(i I) {
	switch i.(type) {
	case T1:
		println("T1")
	case *T2:
		println("T2")
	}

	v, ok := i.(T1)
	if ok {
		v.F()
	}
}

func main() {
	var i I
	i = T1{}
	i.F()
	checkType(i)
	i = &T2{}
	i.F()
	checkType(i)
}