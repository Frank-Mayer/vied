package outline

type Identifier string

func MakeIdentifier(value string) (Identifier, error) {
	return Identifier(value), nil
}
