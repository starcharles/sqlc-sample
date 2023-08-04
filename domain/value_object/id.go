package value_object

import (
	"encoding/json"
	"strconv"
)

type ID struct {
	value int
}

func (i ID) Value() int {
	return i.value
}

func (i ID) Validate() error {
	return nil
}

func (i ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.value)
}

func (i *ID) UnmarshalParam(src string) error {
	val, err := strconv.Atoi(src)
	if err != nil {
		return err
	}
	i.value = val
	return nil
}

func NewID(value int) (*ID, error) {
	id := ID{value: value}

	err := id.Validate()

	if err != nil {
		return nil, err
	}

	return &id, nil
}
