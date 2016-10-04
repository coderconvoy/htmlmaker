package htmlmaker

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

//AddAttrs is a function for the super lazy.
//Simply strings are paired as k-v, odds at the end become innerHTML
func (self *Tag) AddAttrs(s ...string) {
	ls := len(s)
	for i := 0; i+1 < ls; i += 2 {
		self.Attrs[s[i]] = s[i+1]
	}
	if ls%2 == 1 {
		self.Inner = s[ls-1]
	}

}

func (self *Tag) AddChildren(ts ...*Tag) {
	self.Children = append(self.Children, ts...)
}

func Childless(ttype string) bool {
	if ttype == "br" {
		return true
	}
	if ttype == "img" {
		return true
	}
	return false

}
func (self *Tag) String() string {
	return self.toString("")
}

func (self *Tag) toString(pre string) string {
	res := pre + "<" + self.TType
	for k, v := range self.Attrs {
		res += " " + k + "=" + "\"" + v + "\""
	}
	res += ">"
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
	res += "</" + self.TType + ">\n"

	return res

}
