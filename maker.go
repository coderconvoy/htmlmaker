package htmlmaker

import "strings"

type Attr struct {
	Name string
	Val  string
}

type Tag struct {
	TType    string
	Attrs    []Attr
	Children []*Tag
	Inner    string
}

func NewTag(kind string, s ...string) *Tag {
	res := &Tag{kind, []Attr{}, make([]*Tag, 0), ""}
	res.AddAttrs(s...)
	return res
}

func NewParent(kind string, children []*Tag, s ...string) *Tag {
	res := &Tag{kind, []Attr{}, children, ""}
	res.AddAttrs(s...)
	return res
}

func NewTextTag(kind string, inner string, s ...string) *Tag {
	res := &Tag{kind, []Attr{}, make([]*Tag, 0), inner}
	res.AddAttrs(s...)
	return res
}

func NewText(s string) *Tag {
	return &Tag{
		TType: "text",
		Inner: s,
	}
}

// NewPage takes the standard requests, for a page, and returns a page object
//Optional string Params:
//title - The page title
//css - coma separated links to css,
//js - coma separated links to js,
//return the top page object, and the body
func NewPage(ss ...string) (*Tag, *Tag) {
	//fill params
	title, css, js := "", "", ""
	if len(ss) > 0 {
		title = ss[0]
	}
	if len(ss) > 1 {
		css = ss[1]
	}
	if len(ss) > 2 {
		js = ss[2]
	}

	//create bases
	dt := NewTag("!DOCTYPE", "html")
	mh := NewTag("html")
	head := NewTag("head")
	body := NewTag("body")
	mh.AddChildren(head, body)
	head.AddChildren(
		NewTextTag("title", title),
		NewTag("meta", "charset", "utf-8"),
	)
	for _, s := range strings.Split(css, ",") {
		head.AddChildren(NewTag("link", "rel", "stylesheet", "type", "text/css", "href", s))
	}

	for _, _ = range js {
	}

	return NewParent("page", []*Tag{dt, mh}), body

}

//AddAttrs is a function for the super lazy.
//Simply strings are paired as k-v, Odds at the end, become unstrung
func (t *Tag) AddAttrs(s ...string) {
	ls := len(s)

	for i := 0; i+1 < ls; i += 2 {
		added := false
		for _, v := range t.Attrs {
			if v.Name == s[i] {
				v.Val = s[i+1]
				added = true
				break
			}
		}
		if !added {
			t.Attrs = append(t.Attrs, Attr{s[i], s[i+1]})
		}

	}
	if ls%2 == 1 {
		t.Attrs = append(t.Attrs, Attr{s[len(s)-1], ""})
	}

}

func (self *Tag) AddChildren(ts ...*Tag) {
	self.Children = append(self.Children, ts...)
}

func Childless(ttype string) bool {
	ttype = strings.ToLower(ttype)
	childless := []string{"br", "img", "meta", "!doctype"}
	for _, s := range childless {
		if s == ttype {
			return true
		}
	}
	return false
}
func (self *Tag) String() string {
	return self.toString("")
}

func (self *Tag) toString(pre string) string {
	res := ""
	pre2 := pre
	if self.TType != "page" {
		res = pre + "<" + self.TType
		for _, v := range self.Attrs {
			if v.Val == "" {
				res += " " + v.Name
				continue
			}
			res += " " + v.Name + "=" + "\"" + v.Val + "\""
		}
		res += ">"
		pre2 = pre + " "
	}
	if Childless(self.TType) {
		return res + "\n"
	}

	res += self.Inner

	if len(self.Children) > 0 {
		res += "\n"

		for i := 0; i < len(self.Children); i++ {
			res += self.Children[i].toString(pre2)
		}
		res += pre
	}
	if self.TType != "page" {
		res += "</" + self.TType + ">\n"
	}

	return res

}
