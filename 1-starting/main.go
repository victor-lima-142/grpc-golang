package main

import (
	"exemplopb/src/pb/users"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func readData() {
	var user1 users.User
	dados, err := os.ReadFile("dados.txt")
	if err != nil {
		panic(err)
	}
	err = proto.Unmarshal(dados, &user1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user: %+v\n", user1)
}
func main() {
	user1 := users.User{
		Id:       1,
		Name:     "John Doe",
		Email:    "john.doe@mail.com",
		Password: "123456",
	}

	out, err := proto.Marshal(&user1)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("dados.txt", out, 0644)
	if err != nil {
		panic(err)
	}
	readData()
}
