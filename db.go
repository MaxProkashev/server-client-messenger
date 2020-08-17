package main

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
)

// создание таблицы
func createTable(db *sql.DB, name string) {
	if name == "all_users" {
		_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + name + " (id TEXT PRIMARY KEY, username TEXT, CreatedAt INT);")
		if err != nil {
			log.Fatalf("[X] Could not create %s table. Reason: %s", name, err.Error())
		} else {
			log.Printf("[OK] Create %s table", name)
		}
	} else if name == "chats" {
		_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + name + " (id TEXT PRIMARY KEY, name TEXT, users TEXT, CreatedAt INT);")
		if err != nil {
			log.Fatalf("[X] Could not create %s table. Reason: %s", name, err.Error())
		} else {
			log.Printf("[OK] Create %s table", name)
		}
	} else if name == "messages" {
		_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + name + " (id TEXT PRIMARY KEY, chat TEXT, author TEXT, text TEXT, CreatedAt INT);")
		if err != nil {
			log.Fatalf("[X] Could not create %s table. Reason: %s", name, err.Error())
		} else {
			log.Printf("[OK] Create %s table", name)
		}
	} else {
		log.Printf("[ERR] Wrong %s table DB", name)
	}
}

// Удаление таблицы
func dropTable(db *sql.DB, name string) {
	_, err := db.Exec("DROP TABLE " + name + ";")
	if err != nil {
		log.Fatalf("[X] Could not drop %s table. Reason: %s", name, err.Error())
	} else {
		log.Printf("[OK] Drop %s table", name)
	}
}

func createUser(db *sql.DB, user User) {
	_, err := db.Exec("INSERT INTO all_users (id,username,CreatedAt) VALUES ('" + user.ID + "', '" + user.Username + "', " + strconv.Itoa(user.CreatedAt) + ");")
	if err != nil {
		log.Fatalf("[X] Could not create new user. Reason: %s", err.Error())
	} else {
		log.Printf("[OK] New user %s", user.ID)
	}
}
func createChat(db *sql.DB, chat Chat) {
	users := ""
	for _, i := range chat.Users {
		users = users + i + ","
	}
	_, err := db.Exec("INSERT INTO chats (id,name,users,CreatedAt) VALUES ('" + chat.ID + "', '" + chat.Name + "','" + users + "' , " + strconv.Itoa(chat.CreatedAt) + ");")
	if err != nil {
		log.Fatalf("[X] Could not create new chat. Reason: %s", err.Error())
	} else {
		log.Printf("[OK] New chat %s", chat.ID)
	}
}
func createMess(db *sql.DB, mess Message) {
	_, err := db.Exec("INSERT INTO messages (id,chat,author,text,CreatedAt) VALUES ('" + mess.ID + "', '" + mess.Chat + "','" + mess.Author + "','" + mess.Text + "' , " + strconv.Itoa(mess.CreatedAt) + ");")
	if err != nil {
		log.Fatalf("[X] Could not create new message. Reason: %s", err.Error())
	} else {
		log.Printf("[OK] New message %s", mess.ID)
	}
}

func getAllUser(db *sql.DB) (allUser []User) {
	var singUser User
	rows, err := db.Query("SELECT id,username,CreatedAt FROM all_users;")
	defer rows.Close()
	if err != nil {
		log.Fatalf("[X] Could not select. Reason: %s", err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&singUser.ID, &singUser.Username, &singUser.CreatedAt)
			allUser = append(allUser, singUser)
		}
	}
	return allUser
}

func getUserChat(db *sql.DB, idUser string) (Chats []Chat) {
	var singCh Chat
	var users string
	rows, err := db.Query("SELECT id,name,users,CreatedAt FROM chats;")
	defer rows.Close()
	if err != nil {
		log.Fatalf("[X] Could not select. Reason: %s", err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&singCh.ID, &singCh.Name, &users, &singCh.CreatedAt)
			for _, i := range strings.Split(users, ",") {

				if i == idUser {
					singCh.Users = strings.Split(users, ",")
					Chats = append(Chats, singCh)
				}
			}
		}
	}
	return Chats
}

func getMessChat(db *sql.DB, idChat string) (Messages []Message) {
	var singMess Message
	rows, err := db.Query("SELECT id,chat,author,text,CreatedAt FROM messages WHERE chat = '" + idChat + "';")
	defer rows.Close()
	if err != nil {
		log.Fatalf("[X] Could not select. Reason: %s", err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&singMess.ID, &singMess.Chat, &singMess.Author, &singMess.Text, &singMess.CreatedAt)
			Messages = append(Messages, singMess)
		}
	}
	return Messages
}
