package main

import (
	"fmt"
	"mainserver/server"
	"mainserver/store"
	"os"
)

func main(){

	myStore,err:=store.New(os.Getenv("DATABASE_URL"))

	if err!=nil{
		fmt.Println(err)
	}

	s:=server.New(myStore.Conn)

	s.Start()
	
	defer myStore.Close()
}