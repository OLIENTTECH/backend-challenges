package ulid

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func NewULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.Reader, 0)

	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
