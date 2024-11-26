package migrations

import "fmt"

var (
	ErrNoChange = fmt.Errorf("no migrations to apply")
)
