package common
import (
	"testing"
	"fmt"
)
func TestStrToTime(t *testing.T) {
	a, err := StrToTime("2016-09-09")
	fmt.Print(a, err)
}

func TestDecimal(t *testing.T) {
	fmt.Print(Decimal(142313.31242134))
}