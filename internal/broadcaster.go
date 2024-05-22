package internal

import (
	"strings"
	"time"
)

func Broadcaster() {
	currentTime := time.Now()

	for {
		select {
		case msg := <-messages:
			for user, conn := range clients {
				msg.text = strings.ReplaceAll(msg.text, "\n", "")
				if len(msg.text) < 1 {
					break
				}
				if msg.clientName == conn.RemoteAddr().String() {
					continue
				}
				if msg.clientName != user {
					conn.Write([]byte("\n" + msg.text + "\n"))
				}
				conn.Write([]byte("[" + currentTime.Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ":"))

			}
		case msg := <-leaving:
			for user, conn := range clients {
				if msg.clientName != user {
					conn.Write([]byte("\n" + msg.text))
				}
				conn.Write([]byte("[" + currentTime.Format("2006-01-02 15:04:05") + "]" + "[" + user + "]" + ":"))

			}
		}
	}
}
