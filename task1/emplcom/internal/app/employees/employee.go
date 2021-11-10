package employees

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID          uuid.UUID
	Name	 	string
	Position	string
	Env			[]uuid.UUID
	CreateAt	time.Time
	UpdateAt	time.Time
	DeleteAt	*time.Time
}