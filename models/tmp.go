package models

import (
	"time"
)

type Tmp struct {
	Temperature float32
	Timestamp   time.Time
}
