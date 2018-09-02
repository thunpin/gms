package jwt

import "time"

type Obj struct {
	Id   int64
	Name string
	Date time.Time
}

func newObj() *Obj {
	return &Obj{1, "name", time.Now()}
}
