package main

import (
	"context"
	"fmt"

	"github.com/frzifus/winsize"
)

func main() {
	ws, err := winsize.Get()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current size: %s\n", ws)

	watcher := winsize.NewWatcher()

	go func() {
		if err := watcher.Run(context.Background()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("subscribe")
	for ws := range watcher.Subscribe() {
		fmt.Printf("Current size: %s\n", &ws)
	}
}
