package main

import (
	"fmt"
	"log"
	"os/user"

	"github.com/howeyc/fsnotify"
)

func main() {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	//Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("err:", err)
			}
		}
	}()

	histFile := fmt.Sprintf("%s/.bash_history", usr.HomeDir)
	err = watcher.Watch(histFile)
	if err != nil {
		log.Fatal(err)
	}

	// Hang so program doesn't close
	<-done

	/* .. stuff goes here ... */
	watcher.Close()
}
