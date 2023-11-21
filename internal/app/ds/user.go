package ds

import (
	"drones/internal/app/role"

	"github.com/google/uuid"
)

type User struct {
	UUID uuid.UUID `gorm:"type:uuid;unique"`
	Name string    `json:"name"`
	Role role.Role `sql:"type:string;"`
	Pass string
}
