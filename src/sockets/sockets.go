package sockets

import (
	"net"
	"bufio"
	"strings"
	"encoding/json"
	"log"
)

var (
	PrintsLog bool = false
)

type msgSimple struct {
	Id int `json:"id"`
}

type Socket struct {

	closed bool
	reader *bufio.Reader

	ActorName string
	Conn net.Conn

	InterpretClosure func(int, map[string]interface{})bool
}

func (socket *Socket) IsClosed() bool {
	return socket.closed
}

func (socket *Socket) Log(message string, format ...interface{}) {

	// if !PrintsLog {
	// 	return
	// }

	log.Printf(message, format...)
	log.Println()

	// fmt.Printf("- %v %v: ", socket.ActorName, socket.Conn.RemoteAddr())
	// fmt.Printf(message, format...)
	// fmt.Println()
}

func (socket *Socket) ReadMessage() {

	if socket.closed {
		return
	}

	if socket.reader == nil {
		socket.reader = bufio.NewReader(socket.Conn)
	}

	bytes, _, err := socket.reader.ReadLine()

	if err != nil {

		socket.closed = true
		return
	}

	message := strings.Trim(string(bytes), "\n\r\t ")
	log.Printf("Received JSON: %v.\n", message)

	socket.ParseJSON(message)
}

func (socket *Socket) ParseJSON(message string) {

	byt := []byte(message)
	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		socket.Log("Could not interpret JSON: %v.", message)
		return
	}

	id := data["id"]

	if id == nil {
		socket.Log("Data received without an ID: %v.", data)
		return
	}

	if !socket.InterpretClosure(int(id.(float64)), data) {
		socket.Log("Failed to interpret message with ID %v.", id)
	}
}

func (socket *Socket) SendMessage(data interface{}) bool {

	var bytes []byte
	var err error

	if bytes, err = json.Marshal(data); err != nil {
		return false
	}

	message := string(bytes)
	socket.Conn.Write([]byte(message + "\n"))

	return true
}

func (socket *Socket) SendMessageId(id int) bool {
	return socket.SendMessage(msgSimple {
		Id: id,
	})
}

func New(conn net.Conn) *Socket {

	socket := &Socket{}
	socket.ActorName = "Socket"
	socket.Conn = conn

	return socket
}