package test

import (
	"fmt"
	"net/rpc"
	"testing"
)
import "../../ymload"

type A struct {
	a int
}
type B struct {
	c int
	A
}

func TestYmlLoad(t *testing.T) {
	d := ymload.Format("./test.yml")
	t.Log(d)
	dx := "asdasdasd"
	fmt.Println(dx[:2], dx[3:])

}

type Integer int

func (a *Integer) Add(b Integer) Integer { return *a + b }
func TestA(t *testing.T) {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	fmt.Println(sum)
}
func TestB(t *testing.T) {
	//x:=[]int{1,2,3,4,5}
	//ptr:=unsafe.Pointer(&x)
	//fmt.Println((*[1<<10]byte)(ptr))
	fmt.Println(010 >> 1)
	rpc.Register()
}
