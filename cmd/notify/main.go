package main

import (
	"log"
	"os"
	"strings"

	gotinotifier "github.com/TiFlux/go-ti-notifier"
)

func main() {
	notification := strings.Join(os.Args[1:], " ")

	note := gotinotifier.NewNotification(notification)

	note.Title = "Notify"
	note.Message = "teste"

	err := note.Push()

	//If necessary, check error
	if err != nil {
		log.Println("Uh oh! Error with Notify: ", err)
	}
}
