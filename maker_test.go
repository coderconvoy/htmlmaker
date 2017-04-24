package htmlmaker

import (
	"fmt"
	"testing"
)

func test_new(t *testing.T) {
	a := NewTag("button", "id", "btn_hello")
	a.AddChildren(NewTag("img", "id", "img_4", "class", "maker"))
	a.AddChildren(NewTag("button", "POOOO"))
	fmt.Println(a)
}

func Test_Page(t *testing.T) {
	fmt.Println("PAGE")
	p, b := NewPage("gofish", "s/poo/g.css,poopopp.css")
	fmt.Println(p)
	fmt.Println("BODY")
	fmt.Println(b)

}
