package ms

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func NewUUID() string {
	t := time.Unix(1000000, 0)
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return fmt.Sprint(ulid.MustNew(ulid.Timestamp(t), entropy))
}

func NewRequestId() string {
	t := time.Unix(10000000, 0)
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return fmt.Sprint(ulid.MustNew(ulid.Timestamp(t), entropy))
}
