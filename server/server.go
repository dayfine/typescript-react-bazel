package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/dayfine/typescript-react-bazel/server/handlers"
	"github.com/dayfine/typescript-react-bazel/server/session"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	// go globalSessions.GC()
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}

func main() {
	http.HandleFunc("/", handlers.SayhelloName)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/upload", handlers.Upload)
	http.HandleFunc("/count", countHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
