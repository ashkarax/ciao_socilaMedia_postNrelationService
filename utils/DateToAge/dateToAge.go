package datetoage

import (
	"fmt"
	"time"

	interface_dateToAge "github.com/ashkarax/ciao_socilaMedia_postNrelationService/utils/DateToAge/interface"
)

type DateToAge struct{}

func NewDateToAgeUtil() interface_dateToAge.IDateToAge {
	return &DateToAge{}
}

func (dta *DateToAge) DateTOAge(createdAt *time.Time) *string {
	var resp string
	currentTime := time.Now()
	duration := currentTime.Sub(*createdAt)

	minutes := int(duration.Minutes())
	hours := int(duration.Hours())
	days := int(duration.Hours() / 24)
	months := int(duration.Hours() / 24 / 7)

	if minutes < 60 {
		resp = fmt.Sprintf("%d mins ago", minutes)
		return &resp
	} else if hours < 24 {
		resp = fmt.Sprintf("%d hrs ago", hours)
		return &resp
	} else if days < 30 {
		resp = fmt.Sprintf("%d dy ago", days)
		return &resp
	} else {
		resp = fmt.Sprintf("%d weks ago", months)
		return &resp
	}

	resp = fmt.Sprint(createdAt)
	return &resp
}
