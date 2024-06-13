package storage

import (
	"fmt"
	"log"
	"time"
)

type Message struct {
	Id        string    `form:"id" json:"id"`
	Sender    string    `form:"sender" json:"sender"`
	Receiver  string    `form:"receiver" json:"receiver"`
	Text      string    `form:"text" json:"text"`
	CreatedAt time.Time `form:"createdAt" json:"createdAt"`
}

// Read query docs here :
// https://go.dev/doc/tutorial/database-access

func (msg Message) GetMessages(sender string) ([]Message, error) {
	// query DB
	data, err := dbHandler.db.Query("select * from chats where sender=?", sender)
	if err != nil {
		log.Fatal("Failed executing GetMessages query ", err)
		return nil, err
	}
	chats := []Message{}
	for data.Next() {
		msg := Message{}
		data.Scan(&msg.Id, &msg.Sender, &msg.Receiver, &msg.Text, &msg.CreatedAt)
		chats = append(chats, msg)
	}
	return chats, nil
}

func (msg Message) SetMessage() (int64, error) {
	fmt.Println("Inside SetMessage", msg)
	result, err := dbHandler.db.Exec("INSERT INTO chats ( sender, receiver, text, created_at) VALUES ( ?, ?, ?, ?)", msg.Sender, msg.Receiver, msg.Text, msg.CreatedAt)
	if err != nil {
		log.Println("failed adding message", err)
		return 0, err
	}
	fmt.Println("In between result ")
	id, err := result.LastInsertId()
	fmt.Println("ID of inserted chat", id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
