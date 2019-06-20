package common
import (
	"time"
	_ "fmt"
	"strconv"
)

//StrToTime string to timt.Time
func StrToTime(t string) (time.Time, error) {
	tmtmp, err := time.Parse("2006-01-02", t)
	return tmtmp, err
}

// Decimal 4
func Decimal(value float32) float32 {
	v := strconv.FormatFloat(float64(value), 'E', -1, 32)
	val, _ := strconv.ParseFloat(v, 32)
	return float32(val)
	// value, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	// return value

	// return math.Trunc(value*1e4+0.5) * 1e-4
}