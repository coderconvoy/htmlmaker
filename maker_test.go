package htmlmaker

import (
	"fmt"
	"testing"
)

func Test_new(t *testing.T) {
	a := NewTag("button", "id", "btn_hello")
	a.AddChildren(NewTag("img", "id", "img_4", "class", "maker"))
	a.AddChildren(NewTag("button", "POOOO"))
	fmt.Println(a)
}
