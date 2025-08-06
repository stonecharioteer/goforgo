// websocket_chat.go
// Learn WebSocket implementation for real-time communication

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	
	"github.com/gorilla/websocket"
)

// Chat message structure
type Message struct {
	Username  string    `json:"Username"`
	Content   string    `json:"Content"`
	Timestamp time.Time `json:"Timestamp"`
}

// Chat room structure
type ChatRoom struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// Client connection structure
type Client struct {
	chatroom *ChatRoom
	conn     *websocket.Conn
	send     chan []byte
}

// WebSocket upgrader configuration
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin (be more restrictive in production)
		return true
	},
}

func main() {
	fmt.Println("=== WebSocket Chat Server ===")
	
	// Create chat room
	chatRoom := newChatRoom()
	
	// Start chat room in goroutine
	go chatRoom.run()
	
	// Setup routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(chatRoom, w, r)
	})
	
	fmt.Println("Chat server starting on :8082")
	fmt.Println("Open http://localhost:8082 in multiple browser tabs")
	
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// Create new chat room
func newChatRoom() *ChatRoom {
	return &ChatRoom{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Chat room run method (handles client management)
func (cr *ChatRoom) run() {
	for {
		select {
		case client := <-cr.register:
			// Register new client
			cr.clients[client] = true
			fmt.Printf("Client connected. Total clients: %d\n", len(cr.clients))
			
			// Send welcome message to all clients
			welcome := Message{
				Username:  "System",
				Content:   "A new user joined the chat!",
				Timestamp: time.Now(),
			}
			welcomeData, _ := json.Marshal(welcome)
			cr.broadcast <- welcomeData
			
		case client := <-cr.unregister:
			// Unregister client
			if _, ok := cr.clients[client]; ok {
				delete(cr.clients, client)
				close(client.send)
				fmt.Printf("Client disconnected. Total clients: %d\n", len(cr.clients))
				
				// Send goodbye message
				goodbye := Message{
					Username:  "System", 
					Content:   "A user left the chat.",
					Timestamp: time.Now(),
				}
				goodbyeData, _ := json.Marshal(goodbye)
				cr.broadcast <- goodbyeData
			}
			
		case message := <-cr.broadcast:
			// Broadcast message to all clients
			for client := range cr.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(cr.clients, client)
				}
			}
		}
	}
}

// WebSocket handler
func handleWebSocket(chatRoom *ChatRoom, w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	
	// Create new client
	client := &Client{
		chatroom: chatRoom,
		conn:     conn,
		send:     make(chan []byte, 256),
	}
	
	// Register client
	chatRoom.register <- client
	
	// Start client goroutines
	go client.writePump()
	go client.readPump()
}

// Client read pump (handles incoming messages)
func (c *Client) readPump() {
	defer func() {
		c.chatroom.unregister <- c
		c.conn.Close()
	}()
	
	// Set read limits and deadlines
	c.conn.SetReadLimit(512)
	c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})
	
	for {
		// Read message from WebSocket
		_, messageBytes, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket read error: %v", err)
			}
			break
		}
		
		// Parse message
		var msg Message
		if err := json.Unmarshal(messageBytes, &msg); err != nil {
			log.Printf("JSON parse error: %v", err)
			continue
		}
		
		// Add timestamp and broadcast
		msg.Timestamp = time.Now()
		messageData, _ := json.Marshal(msg)
		c.chatroom.broadcast <- messageData
	}
}

// Client write pump (handles outgoing messages)
func (c *Client) writePump() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			
			// Write message to WebSocket
			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("WebSocket write error: %v", err)
				return
			}
			
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("WebSocket ping error: %v", err)
				return
			}
		}
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
	<title>WebSocket Chat</title>
	<style>
		body { font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; }
		#messages { height: 300px; border: 1px solid #ccc; overflow-y: scroll; padding: 10px; margin: 10px 0; }
		#input { width: 70%; padding: 10px; }
		#send { padding: 10px 20px; }
		.message { margin: 5px 0; }
		.username { font-weight: bold; color: #333; }
		.timestamp { color: #666; font-size: 12px; }
	</style>
</head>
<body>
	<h1>WebSocket Chat Room</h1>
	<div>
		Username: <input type="text" id="username" placeholder="Enter your name" value="User">
	</div>
	<div id="messages"></div>
	<div>
		<input type="text" id="input" placeholder="Type a message..." onkeypress="if(event.key==='Enter') sendMessage()">
		<button id="send" onclick="sendMessage()">Send</button>
	</div>

	<script>
		const ws = new WebSocket('ws://localhost:8082/ws');
		const messages = document.getElementById('messages');
		const input = document.getElementById('input');
		const usernameInput = document.getElementById('username');

		ws.onmessage = function(event) {
			const message = JSON.parse(event.data);
			const div = document.createElement('div');
			div.className = 'message';
			div.innerHTML = '<span class="username">' + message.Username + '</span>: ' + 
							message.Content + ' <span class="timestamp">(' + 
							new Date(message.Timestamp).toLocaleTimeString() + ')</span>';
			messages.appendChild(div);
			messages.scrollTop = messages.scrollHeight;
		};

		function sendMessage() {
			const message = {
				Username: usernameInput.value || 'Anonymous',
				Content: input.value,
			};
			ws.send(JSON.stringify(message));
			input.value = '';
		}

		ws.onopen = function() {
			console.log('Connected to chat server');
		};

		ws.onclose = function() {
			console.log('Disconnected from chat server');
		};
	</script>
</body>
</html>`
	
	fmt.Fprint(w, html)
}