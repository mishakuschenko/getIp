package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	//"github.com/go-chi/chi/v5"
)

type User struct {
	IP string
}

func main() {
	//r := chi.NewRouter()

	http.HandleFunc("/myIp", homeHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	userIp := getUserIp(r)
	Me := User{
		IP: userIp,
	}
	temp, _ := template.ParseFiles("./IP.html")
	temp.Execute(w, Me)

}

func getUserIp(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		parts := strings.Split(ip, ",")
		return strings.TrimSpace(parts[0])
	}

	return strings.Split(r.RemoteAddr, ":")[0]
}
