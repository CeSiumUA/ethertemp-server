package models

import (
	"time"
)

type Dat struct {
	Temperature float32
	Humidity    float32
	Timestamp   time.Time
}
