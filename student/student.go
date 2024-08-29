package student

type Student struct {
	Name string
}

func New(name string) *Student {
	return &Student{Name: name}
}
