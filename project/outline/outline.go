package outline

type Outline struct {
	Packages []*Package
}

func Init(root string) (*Outline, error) {
	return &Outline{
		Packages: make([]*Package, 0),
	}, nil
}
