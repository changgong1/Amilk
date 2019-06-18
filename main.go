package main

import (
	"net/http"
	"text/template"

	"./server/gosrc/mlogger"
)

func main() {
	Auth()
	mlogger.LogerPrint("%s", "server start.")
	http.ListenAndServe("127.0.0.1:8080", nil)
	mlogger.LogerPrint("%s", "succeed server")
}

// Auth is the router
func Auth() {
	http.Handle("/js/", http.FileServer(http.Dir("webapp")))
	http.Handle("/css/", http.FileServer(http.Dir("webapp")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpServer V1"))
	})
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/bye", sayBye)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./webapp/index.html")
	if err != nil {
		mlogger.LogerPrint("s%", "获取路经失败")
		return
	}

	t.Execute(w, nil)
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye, this is v1 httpSerever"))
}
