package main

import "github.com/maxheckel/campfirereads/internal/server"

func main() {
	srv, err := server.NewAPI()
	if err != nil {
		panic(err)
	}

	if err := srv.Start(); err != nil {
		panic(err)
	}
}
