package common
import (
	"time"
)

//StrToTime string to timt.Time
func StrToTime(t string) (time.Time, error) {
	tmtmp, err := time.Parse("2006-01-02", t)
	return tmtmp, err
}