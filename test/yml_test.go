package test

import (
	"fmt"
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
