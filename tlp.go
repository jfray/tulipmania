package main

import (
	"log"

	"github.com/howeyc/fsnotify"
)

func main() {
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

	err = watcher.Watch("/Users/jfray/.bash_history")
	if err != nil {
		log.Fatal(err)
	}

	// Hang so program doesn't close
	<-done

	/* .. stuff goes here ... */
	watcher.Close()
}
