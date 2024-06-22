package exception

import "fmt"

type EntityNotFoundException struct {
	Entity string
	ID     int
}

func (e EntityNotFoundException) Error() string {
	return fmt.Sprintf("Entity %s with ID %d not found", e.Entity, e.ID)
}
