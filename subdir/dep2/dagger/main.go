package main

// Dep2 is a two-level-deep module dependency
type Dep2 struct{}

func (m *Dep2) Fn() string {
	return "hi from dep2"
}
