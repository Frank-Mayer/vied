package outline

type (
	Type struct {
		Name  Identifier
		Value TypeValue
	}

	TypeValue interface {
		Name() string
	}
)
