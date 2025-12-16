package uuidutil

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidUUID = errors.New("invalid uuid")
var ErrBinary = errors.New("invalid uuid binary")

func ParseToBinary(id string) ([]byte, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrInvalidUUID
	}

	b, err := u.MarshalBinary()
	if err != nil {
		return nil, ErrBinary
	}

	return b, nil
}
