package main

import "campfirereads/internal/server"

func main() {
	srv, err := server.New("UI")
	if err != nil {
		panic(err)
	}

	if err := srv.Start(); err != nil {
		panic(err)
	}
}
