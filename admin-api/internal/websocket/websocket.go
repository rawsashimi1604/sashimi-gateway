package websocket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// WebSocketServer will hold the active connections and the broadcast channel
type WebSocketServer struct {
	clients map[*websocket.Conn]bool
	mutex   sync.Mutex
}

// NewWebSocketServer creates a new WebSocketServer
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (server *WebSocketServer) HandleClient(w http.ResponseWriter, r *http.Request) {

	type Message struct {
		Message string `json:"message"`
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Info().Msg("error occured when upgrading")
		return
	}
	defer ws.Close()

	// Register the new client
	server.mutex.Lock()
	server.clients[ws] = true
	server.mutex.Unlock()

	log.Info().Msg("New client connected: " + ws.RemoteAddr().String())

	// Listen for new messages from the client
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Info().Msgf("error: %v", err)
			delete(server.clients, ws)
			break
		}
		log.Info().Msg(ws.RemoteAddr().String() + ": " + msg.Message)
		server.BroadcastMessage([]byte("from server"))
	}
}

func (server *WebSocketServer) BroadcastMessage(message []byte) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for client := range server.clients {
		err := client.WriteJSON(map[string]interface{}{"message": message})
		if err != nil {
			log.Info().Msgf("error: %v", err)
			client.Close()
			delete(server.clients, client)
		}
	}
}
