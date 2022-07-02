package troc

type View struct {
	data    map[int][]interface{}
	deleted map[int]bool
}

func (view View) String() string {
	return ""
}
