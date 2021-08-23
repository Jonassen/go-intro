package main

import (
	"fmt"
	"music/music"
	"music/server"
	"music/storage"
)

func main() {
	memStore := storage.NewMemStore()
	musicService := music.NewService(memStore)

	serverDeps := server.ServerDependencies{
		MusicService: *musicService,
	}
	err := server.Start(serverDeps)
	if err != nil {
		fmt.Println("Server ended with error: ", err)
	}
	fmt.Println("Good bye!")
}
