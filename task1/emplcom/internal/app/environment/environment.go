package environment

import (
	"time"

	"github.com/google/uuid"
)

type Environment struct {
	ID          uuid.UUID
	Name	 	string
	CreateAt	time.Time
	UpdateAt	time.Time
	DeleteAt	*time.Time
}
