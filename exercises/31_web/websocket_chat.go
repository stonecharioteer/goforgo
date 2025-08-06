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

// TODO: Chat message structure
type Message struct {
	/* define message fields: Username, Content, Timestamp */
}

// TODO: Chat room structure
type ChatRoom struct {
	/* define room fields: clients map, broadcast channel, register channel, unregister channel */
}

// TODO: Client connection structure
type Client struct {
	/* define client fields: chatroom pointer, connection pointer, send channel */
}

// WebSocket upgrader configuration
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: Allow connections from any origin (be more restrictive in production)
		/* return true for any origin */
	},
}

func main() {
	fmt.Println("=== WebSocket Chat Server ===")
	
	// TODO: Create chat room
	chatRoom := /* create new chat room instance */
	
	// TODO: Start chat room in goroutine
	/* start chat room run method in goroutine */
	
	// Setup routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", /* create websocket handler with chatRoom */
	
	fmt.Println("Chat server starting on :8082")
	fmt.Println("Open http://localhost:8082 in multiple browser tabs")
	
	log.Fatal(http.ListenAndServe(":8082", nil))
}

// TODO: Create new chat room
func newChatRoom() *ChatRoom {
	return &ChatRoom{
		/* initialize all fields */
	}
}

// TODO: Chat room run method (handles client management)
func (cr *ChatRoom) run() {
	for {
		select {
		case client := /* receive from register channel */:
			// TODO: Register new client
			/* add client to clients map */
			/* log new client connection */
			
			// Send welcome message to all clients
			welcome := Message{
				Username:  "System",
				Content:   "A new user joined the chat!",
				Timestamp: time.Now(),
			}
			/* convert welcome to JSON and send to broadcast channel */
			
		case client := /* receive from unregister channel */:
			// TODO: Unregister client
			if /* check if client exists in clients map */ {
				/* remove client from clients map */
				/* close client send channel */
				/* log client disconnection */
				
				// Send goodbye message
				goodbye := Message{
					Username:  "System", 
					Content:   "A user left the chat.",
					Timestamp: time.Now(),
				}
				/* convert goodbye to JSON and send to broadcast channel */
			}
			
		case message := /* receive from broadcast channel */:
			// TODO: Broadcast message to all clients
			for client := range /* iterate over clients */ {
				select {
				case /* send message to client send channel */:
				default:
					/* close client send channel */
					/* remove client from clients map */
				}
			}
		}
	}
}

// TODO: WebSocket handler
func handleWebSocket(chatRoom *ChatRoom, w http.ResponseWriter, r *http.Request) {
	// TODO: Upgrade HTTP connection to WebSocket
	conn, err := /* upgrade connection using upgrader */
	if /* check for error */ {
		/* log error */
		return
	}
	
	// TODO: Create new client
	client := &Client{
		/* initialize client fields */
	}
	
	// TODO: Register client
	/* send client to chatRoom register channel */
	
	// TODO: Start client goroutines
	/* start writePump in goroutine */
	/* start readPump in goroutine */
}

// TODO: Client read pump (handles incoming messages)
func (c *Client) readPump() {
	defer func() {
		/* send client to chatroom unregister channel */
		/* close websocket connection */
	}()
	
	// Set read limits and deadlines
	c.conn.SetReadLimit(512)
	c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})
	
	for {
		// TODO: Read message from WebSocket
		_, messageBytes, err := /* read message from connection */
		if /* check for error */ {
			/* log error if not normal closure */
			break
		}
		
		// TODO: Parse message
		var msg Message
		if err := /* unmarshal JSON message */; err != nil {
			/* log parse error and continue */
			continue
		}
		
		// TODO: Add timestamp and broadcast
		msg.Timestamp = time.Now()
		messageData, _ := json.Marshal(msg)
		/* send messageData to chatroom broadcast channel */
	}
}

// TODO: Client write pump (handles outgoing messages)
func (c *Client) writePump() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		/* stop ticker */
		/* close websocket connection */
	}()
	
	for {
		select {
		case message, ok := /* receive from client send channel */:
			/* set write deadline to 10 seconds from now */
			
			if /* check if channel was closed */ {
				/* write close message to connection */
				return
			}
			
			// TODO: Write message to WebSocket
			if err := /* write message to connection */; err != nil {
				/* log error */
				return
			}
			
		case /* receive from ticker channel */:
			/* set write deadline to 10 seconds from now */
			if err := /* write ping message to connection */; err != nil {
				/* log error */
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