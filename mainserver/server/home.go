package server

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"mainserver/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader instance for WebSocket connection
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for this example, adjust if needed
		return true
	},
}

// Compute the Sec-WebSocket-Accept header value
func computeSecWebSocketAccept(key string) string {
	// The WebSocket GUID defined in the protocol
	const wsGUID = "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	// Concatenate Sec-WebSocket-Key and the GUID
	data := key + wsGUID
	// Perform SHA-1 hash on the concatenated string
	hasher := sha1.New()
	hasher.Write([]byte(data))
	hash := hasher.Sum(nil)
	// Return the base64-encoded SHA-1 hash
	return base64.StdEncoding.EncodeToString(hash)
}

func (s *Server) Home(g *gin.Context) {
	// Extract the Sec-WebSocket-Key from the request
	websocketKey := g.Request.Header.Get("Sec-WebSocket-Key")
	token := g.Request.Header.Get("Authorization")

	_, err := utils.ValidateToken(token)
	fmt.Println("error is", err)

	if err != nil {
		utils.ResponseFormatter(g, http.StatusForbidden, false, nil, utils.ErrorUserForbidden)
		return
	}
	if websocketKey == "" {
		utils.ResponseFormatter(g, http.StatusBadRequest, false, nil, fmt.Errorf("missing Sec-WebSocket-Key"))
		return
	}

	// Compute the Sec-WebSocket-Accept header
	secWebSocketAccept := computeSecWebSocketAccept(websocketKey)
	fmt.Println(websocketKey, secWebSocketAccept)

	// Set the required WebSocket headers and upgrade the connection
	g.Writer.Header().Set("Upgrade", "websocket")
	g.Writer.Header().Set("Connection", "Upgrade")
	g.Writer.Header().Set("Sec-WebSocket-Accept", secWebSocketAccept)

	// Send HTTP 101 Switching Protocols response
	g.Writer.WriteHeader(http.StatusSwitchingProtocols)

	// Upgrade the connection to WebSocket
	conn, err := upgrader.Upgrade(g.Writer, g.Request, nil)
	if err != nil {
		utils.ResponseFormatter(g, http.StatusInternalServerError, false, nil, err)
		return
	}
	defer conn.Close()

	// Echo back received messages (basic example)
	for {
		messageType, message, err := conn.ReadMessage()
		fmt.Println(messageType, string(message))
		if err != nil {
			// Handle WebSocket errors (e.g., connection closed)
			break
		}
		// Send the message back to the client
		if err := conn.WriteMessage(messageType, message); err != nil {
			break
		}
	}
}
