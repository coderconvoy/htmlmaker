package htmlmaker

import "strings"

func QSelect(name string, options ...string) *Tag {
	ops := []*Tag{}
	id := ""
	for _, v := range options {
		if strings.HasPrefix(v, "#") {
			id = strings.TrimPrefix(v, "#")
			continue
		}
		ops = append(ops, NewTextTag("option", v, "value", v))
	}
	res := NewParent("select", ops, "id", id, "name", name)
	if id != "" {
		res.SetAttr("id", id)
	}

	return res
}

func QInput(ttype, name string, options ...string) *Tag {
	return NewTag("input", append(options, "name", name, "type", ttype)...)
}

func QSubmit(text string) *Tag {
	return NewTag("input", "type", "submit", "text", text, "value", text)
}

func QForm(action string, chids []*Tag, options ...string) *Tag {
	return NewParent("form", chids, append(options, "method", "post", "action", action)...)
}
