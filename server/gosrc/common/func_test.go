package common
import (
	"testing"
	"fmt"
)
func TestStrToTime(t *testing.T) {
	a, err := StrToTime("2016-09-09")
	fmt.Println(a, err)
}