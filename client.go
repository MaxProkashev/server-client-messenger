package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// User ..
type User struct {
	ID        string
	Username  string
	CreatedAt int
}

// Chat ..
type Chat struct {
	ID        string
	Name      string
	Users     []string
	CreatedAt int
}

// Message ..
type Message struct {
	ID        string
	Chat      string
	Author    string
	Text      string
	CreatedAt int
}

// создает нового юзера и записывает его в бд
// рандомный id по uuid v.4
func newUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users/add" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method GET not allowed", http.StatusMethodNotAllowed)
		return
	}
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Misspelled request", http.StatusMisdirectedRequest)
		return
	}
	user := User{
		ID:        uuid.New().String(),
		Username:  username,
		CreatedAt: int(time.Now().Unix()),
	}
	createUser(db, user)
	fmt.Fprintf(w, "New user ID: %s", user.ID)

	//http.ServeFile(w, r, "home.html")
}

// выводит список всех юзеров
func showUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/users/show" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method GET not allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Println("[OK] Show user request")
	fmt.Fprintf(w, "All Users ID:\n\n")
	for _, user := range getAllUser(db) {
		fmt.Fprintln(w, "ID: "+user.ID+"\nUsername: "+user.Username+"\nCreate time: "+time.Unix(int64(user.CreatedAt), 0).Format(time.UnixDate)+"\n")
	}
	//http.ServeFile(w, r, "home.html")
}

func newChat(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chats/add" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method GET not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")
	users := r.URL.Query().Get("users")
	if name == "" || users == "" {
		http.Error(w, "Misspelled request", http.StatusMisdirectedRequest)
		return
	}
	chat := Chat{
		ID:        uuid.New().String(),
		Name:      name,
		Users:     strings.Split(users, ","),
		CreatedAt: int(time.Now().Unix()),
	}
	createChat(db, chat)
	fmt.Fprintf(w, "New chat ID: %s", chat.ID)

	//http.ServeFile(w, r, "home.html")
}

func getChat(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chats/get" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method GET not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := r.URL.Query().Get("user")
	if user == "" {
		http.Error(w, "Misspelled request", http.StatusMisdirectedRequest)
		return
	}
	log.Println("[OK] Show user chat request")
	fmt.Fprintf(w, "All Chats User ID "+user+":\n\n")
	for _, chat := range getUserChat(db, user) {
		fmt.Fprintln(w, "ID: "+chat.ID+"\nName: "+chat.Name)
		fmt.Fprintln(w, "Users:")
		for _, user := range chat.Users {
			fmt.Fprintln(w, user)
		}
		fmt.Fprintln(w, "Create time: "+time.Unix(int64(chat.CreatedAt), 0).Format(time.UnixDate)+"\n")
	}

	//http.ServeFile(w, r, "home.html")
}
func getMess(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/messages/get" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method GET not allowed", http.StatusMethodNotAllowed)
		return
	}

	chat := r.URL.Query().Get("chat")
	if chat == "" {
		http.Error(w, "Misspelled request", http.StatusMisdirectedRequest)
		return
	}
	log.Println("[OK] Show message chat request")
	fmt.Fprintf(w, "All Message Chat ID "+chat+":\n\n")
	for _, mess := range getMessChat(db, chat) {
		fmt.Fprintln(w, "ID: "+mess.ID+"\nChat: "+mess.Chat+"\nAuthor: "+mess.Author+"\nText: "+mess.Text+"\nCreate time: "+time.Unix(int64(mess.CreatedAt), 0).Format(time.UnixDate)+"\n")
	}

	//http.ServeFile(w, r, "home.html")
}

func newMess(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/messages/add" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method GET not allowed", http.StatusMethodNotAllowed)
		return
	}

	chat := r.URL.Query().Get("chat")
	author := r.URL.Query().Get("author")
	text := r.URL.Query().Get("text")
	if chat == "" || author == "" || text == "" {
		http.Error(w, "Misspelled request", http.StatusMisdirectedRequest)
		return
	}
	message := Message{
		ID:        uuid.New().String(),
		Chat:      chat,
		Author:    author,
		Text:      text,
		CreatedAt: int(time.Now().Unix()),
	}
	createMess(db, message)
	fmt.Fprintf(w, "New message ID: %s", message.ID)

	//http.ServeFile(w, r, "home.html")
}
