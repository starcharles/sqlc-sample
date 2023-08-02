package value_object

type ID struct {
	value int
}

func (i ID) Validate() error {
	return nil
}

func (i ID) Value() int {
	return i.value
}

func NewID(value int) (*ID, error) {
	id := ID{value: value}

	err := id.Validate()

	if err != nil {
		return nil, err
	}

	return &id, nil
}
