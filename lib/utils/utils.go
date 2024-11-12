package utils

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

func NewULID() ulid.ULID {
	ulidValue, err := ulid.New(ulid.Timestamp(time.Now()), rand.Reader)
	if err != nil {
		fmt.Println(err.Error()) //nolint: forbidigo
	}

	return ulidValue
}
