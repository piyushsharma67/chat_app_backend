package main

import (
	"fmt"
	"mainserver/server"
	"mainserver/store"
)

func main() {

	// myStore,err:=store.New(os.Getenv("DATABASE_URL"))
	myStore, err := store.New("postgres://postgres:password@192.168.1.7:5432/mydb?sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	s := server.New(myStore.Conn)

	s.Start()

	defer myStore.Close()
}
