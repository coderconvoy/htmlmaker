package htmlmaker

import "strings"

type Tag struct {
	TType    string
	Attrs    map[string]string
	Children []*Tag
	Inner    string
}

func NewTag(kind string, s ...string) *Tag {
	res := &Tag{kind, make(map[string]string), make([]*Tag, 0), ""}
	res.AddAttrs(s...)
	return res
}

func NewParent(kind string, children []*Tag, s ...string) *Tag {
	res := &Tag{kind, make(map[string]string), children, ""}
	res.AddAttrs(s...)
	return res
}

func NewTextTag(kind string, inner string, s ...string) *Tag {
	res := &Tag{kind, make(map[string]string), make([]*Tag, 0), inner}
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
func (self *Tag) AddAttrs(s ...string) {
	ls := len(s)
	for i := 0; i+1 < ls; i += 2 {
		self.Attrs[s[i]] = s[i+1]
	}
	if ls%2 == 1 {
		self.Attrs[s[len(s)-1]] = ""
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
	if self.TType != "page" {
		res = pre + "<" + self.TType
		for k, v := range self.Attrs {
			if v == "" {
				res += " " + k
				continue
			}
			res += " " + k + "=" + "\"" + v + "\""
		}
		res += ">"
	}
	if Childless(self.TType) {
		return res + "\n"
	}

	res += self.Inner

	if len(self.Children) > 0 {
		res += "\n"

		for i := 0; i < len(self.Children); i++ {
			res += self.Children[i].toString(pre + " ")

		}
		res += pre
	}
	if self.TType != "page" {
		res += "</" + self.TType + ">\n"
	}

	return res

}
