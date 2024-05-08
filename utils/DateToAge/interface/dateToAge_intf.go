package interface_dateToAge

import "time"

type IDateToAge interface {
	DateTOAge(createdAt *time.Time) *string
}
