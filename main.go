package main

import (
	// "encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

// ws connection settings and config
var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan CreateMessageRequest)
	mutex     sync.Mutex
)

func main() {
	go handleMessages()

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// routes
	// r.GET("/", func(c *gin.Context) {
		// c.HTML(http.StatusOK, "home.html", nil)
	// })

	r.POST("/createroom", CreateRoom)

	r.GET("/rooms/:id", GetRoom)

	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})

	r.Run(":8080")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	for {
		var msg CreateMessageRequest
		err := conn.ReadJSON(&msg)
		// CreateMessage(c, msg)
		if err != nil {
			mutex.Lock()
			delete(clients, conn)
			mutex.Unlock()
			break
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func CreateMessage(c *gin.Context, msg CreateMessageRequest) {
	message := Message{
		Id:        CreateUniqueId(),
		User:      msg.User,
		Content:   msg.Content,
		CreatedAt: time.Now().Format(time.RFC3339),
		RoomId:    msg.RoomId,
	}
	var redisClient = GetDb()
	err := redisClient.HSet(c, "room:"+message.RoomId+"messages", message).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create message"})
		return
	}
}

func CreateUniqueId() string {
	return uuid.New().String()
}

func CreateRoom(c *gin.Context) {
	var createRoomRequest CreateRoomRequest
	if err := c.ShouldBindJSON(&createRoomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// generate a unique ID for the room
	code := CreateUniqueId()

	// create room struct
	room := Room{
		Id:        code,
		Name:      createRoomRequest.Name,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	// Here you would typically save the room to a database
	var redisClient = GetDb()
	err := redisClient.HSet(c, "room:"+room.Id, room).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create room"})
		return
	}

	// For this example, we'll just return it
	c.JSON(http.StatusOK, room)
}

func GetRoom(c *gin.Context) {
	roomId := c.Param("id")

	// check if room exists
	var redisClient = GetDb()
	_, err := redisClient.HGetAll(c, "room:"+roomId).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}
	
	// get room messages
	messages, err := redisClient.HGetAll(c, "room:"+roomId+"messages").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages"})
		return
	}

	c.HTML(http.StatusOK, "room.html", messages)
}
