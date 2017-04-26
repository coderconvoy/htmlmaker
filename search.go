package htmlmaker

type TagFilter func(*Tag) bool

//GetFirst takes a filter and returns the first (Depth first) Element to fulfil it,
func (t *Tag) GetFirst(f TagFilter, maxD int) *Tag {
	if f(t) {
		return t
	}
	if maxD == 0 { //negs go full depth -- careful of loops
		return nil
	}
	for _, c := range t.Children {
		res := c.GetFirst(f, maxD-1)
		if res != nil {
			return res
		}
	}
	return nil
}

func (t *Tag) GetAll(f TagFilter, maxD int) []*Tag {
	res = []*Tag{}
	if f(t) {
		res = append(res, t)
	}
	if maxD == 0 {
		return res
	}
	for _, c := range t.Children {
		cres := c.GetAll(f, maxD-1)
		res = append(res, cres...)
	}
	return res
}

func (t *Tag) GetElementById(id string, md int) *Tag {
	return t.GetFirst(func(t *Tag) bool {
		for _, a := range t.Attrs {
			if a.Name == "id" && a.Val == id {
				return true
			}
		}
		return false
	}, md)
}
