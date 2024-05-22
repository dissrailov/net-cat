package internal

import (
	"bufio"
	"log"
	"net"
	"strings"
)

var (
	clients  = make(map[string]net.Conn)
	leaving  = make(chan message)
	messages = make(chan message)
)

type message struct {
	text       string
	clientName string
}

func Handle(conn net.Conn) {
	// Запрашиваем имя клиента
	pattern := Welcome()
	_, err := conn.Write([]byte(pattern + "\n" + "[Enter your name]: "))
	if err != nil {
		log.Printf("Error writing to client: %s", err)
		return
	}

	// Считываем имя клиента
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading from client: %s", err)
		return
	}
	clientName := strings.TrimSpace(string(buffer[:n]))

	for {
		// Проверяем, существует ли уже клиент с таким именем
		if _, ok := clients[clientName]; ok {
			conn.Write([]byte("Sorry, this name is already taken\n"))
			_, err := conn.Write([]byte("[Enter your name]: "))
			if err != nil {
				log.Printf("Error writing to client: %s", err)
				return
			}
			n, err := conn.Read(buffer)
			if err != nil {
				log.Printf("Error reading from client: %s", err)
				return
			}
			clientName = strings.TrimSpace(string(buffer[:n]))
		} else {
			break // Выходим из цикла, если имя уникально
		}
	}

	// Добавляем клиента в карту клиентов
	clients[clientName] = conn

	leaving <- newMessage((clientName)+" has joined our chat..."+"\n", clientName)
	// Отслеживаем ввод сообщений от клиента
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- newMessage(MessageFromUser(clientName)+input.Text(), clientName)
	}
	// Клиент отключился
	delete(clients, clientName)
	leaving <- newMessage((clientName)+" has left our chat..."+"\n", clientName)
	conn.Close()
}

func newMessage(msg string, clientName string) message {
	return message{
		text:       msg,
		clientName: clientName,
	}
}
