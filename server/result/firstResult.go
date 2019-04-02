// the first result

package result

import (
	"encoding/json"
	"net/http"
)

// FirstResult the first result
// Level = 0, 1, 2, 3, 4 info, error, warn
type FirstResult struct {
	Level int
	Code  string
	Msg   string
}

// Init initial a FirstResult
func (R *FirstResult) Init(level int, Code, Msg string) {
	R.Level = level
	R.Code = Code
	R.Msg = Msg
}

// Ret is returning the FirstResult
func (R *FirstResult) Ret(w http.ResponseWriter, r *http.Request) {
	s, _ := json.Marshal(R)
	w.Write(s)
	return
}
