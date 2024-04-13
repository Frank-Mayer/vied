package outline

type Package struct {
	Name  Identifier
	Types []*Type
	Funcs []*Func
}
