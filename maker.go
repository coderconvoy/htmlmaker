package htmlmaker

type Tag struct {
	TType    string
	Attrs    map[string]string
	Children []*Tag
	innerT   string
}

func NewTag(kind string, s ...string) *Tag {
	res := &Tag{kind, make(map[string]string), make([]*Tag, 0), ""}
	i := 0
	for i = 0; i+1 < len(s); i += 2 {
		res.Attrs[s[i]] = s[i+1]
	}
	return res

}

func (self *Tag) addChildren(ts ...*Tag) {
	self.Children = append(self.Children, ts...)
}

func Childless(ttype string) bool {
	if ttype == "br" {
		return true
	}
	if ttype == "image" {
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

	if len(self.Children) > 0 {
		res += "\n"
	}
	for i := 0; i < len(self.Children); i++ {
		res += self.Children[i].toString(pre + " ")

	}
	if len(self.Children) > 1 {
		res += pre
	}
	res += "</" + self.TType + ">\n"

	return res

}