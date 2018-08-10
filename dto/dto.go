package dto

import (
	"github.com/jinzhu/gorm"
	"time"
)

type WorkDay struct {
	gorm.Model
	Tasks       []Task
	Description string
	WorkTime    time.Duration
}

type Task struct {
	gorm.Model
	Date      time.Time
	StartTime time.Time
	EndTime   time.Time
}
