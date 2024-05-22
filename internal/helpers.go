package internal

import (
	"fmt"
	"log"
	"os"
	"time"
)

func MessageFromUser(name string) string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	return fmt.Sprintf("\n[%s][%s]:", formattedTime, name)
}

func Welcome() string {
	image, err := os.ReadFile("assets/pattern.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(image)
}
