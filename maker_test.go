package htmlmaker

import (
	"fmt"
	"testing"
)

func Test_new(t *testing.T) {
	a := NewTag("button", "id", "btn_hello")
	a.addChildren(NewTag("img", "id", "img_4", "class", "maker"))
	a.addChildren(NewTag("button", "POOOO"))
	fmt.Println(a)
}
