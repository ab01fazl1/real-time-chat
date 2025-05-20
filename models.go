package main

type User struct {
	Username string `json:"username"`
}

type Message struct {
	Id        string `json:"id" redis:"id"`
	User      User   `json:"user" redis:"user"`
	Content   string `json:"content" redis:"content"`
	CreatedAt string `json:"created_at" redis:"created_at"`
	RoomId    string `json:"room_id" redis:"room_id"`
}

type CreateMessageRequest struct {
	User    User   `json:"user"`
	Content string `json:"content"`
	RoomId  string `json:"room_id"`
}

type CreateRoomRequest struct {
	Name string `json:"name"`
}

type Room struct {
	Id        string `json:"id" redis:"id"`
	Name      string `json:"name" redis:"name"`
	CreatedAt string `json:"created_at" redis:"created_at"`
}
