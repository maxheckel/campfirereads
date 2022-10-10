package main

import "campfirereads/internal/server"

func main() {
	srv, err := server.New("API")
	if err != nil {
		panic(err)
	}

	if err := srv.Start(); err != nil {
		panic(err)
	}
}
