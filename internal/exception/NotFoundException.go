package exception

import "fmt"

type NotFoundException struct {
	Entity string
	ID     int
}

func (e NotFoundException) Error() string {
	return fmt.Sprintf("%s with ID %d not found", e.Entity, e.ID)
}
