package chat

import (
	"encoding/json"
)

func encode[T any](object T) ([]byte, error) {
	bytes, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func decode[T any](input []byte, obj T) error {
	err := json.Unmarshal(input, obj)
	if err != nil {
		return err
	}
	return nil
}
