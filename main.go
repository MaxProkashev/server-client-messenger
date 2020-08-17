package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var (
	addr = flag.String("addr", ":9000", "http service address")

	// DB
	dbURI = "postgres://yzjjajvmbsktuh:a77133c5481237604f6f0eadbfd2f50e8523e34f4050e99b8f08ac299bb49bab@ec2-34-238-26-109.compute-1.amazonaws.com:5432/da9rj6924msh94"
	db, _ = sql.Open("postgres", dbURI)
)

func main() {
	// предварительная настройка порт, бд
	flag.Parse()

	createTable(db, "all_users") // таблица user
	createTable(db, "chats")     // таблица chat
	createTable(db, "messages")  // таблица message

	http.HandleFunc("/users/add", newUser)
	http.HandleFunc("/users/show", showUser)

	http.HandleFunc("/chats/add", newChat)
	http.HandleFunc("/chats/get", getChat)

	http.HandleFunc("/messages/add", newMess)
	http.HandleFunc("/messages/get", getMess)

	http.HandleFunc("/drop", drop)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func drop(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method GET not allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Println("[OK] Drop table")
	dropTable(db, "all_users")
	dropTable(db, "chats")
	dropTable(db, "messages")

	createTable(db, "all_users") // таблица user
	createTable(db, "chats")     // таблица chat
	createTable(db, "messages")  // таблица message
	//http.ServeFile(w, r, "home.html")
}
