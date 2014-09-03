package main

import (
	"../."
	"encoding/json"
	"fmt"
)

const (
	ClientID          = "<client-id>"
	State             = "STATE"
	ClientSecret      = "<client-secret>"
	AuthorizationCode = "<authorization-code> - https://developer.nest.com/documentation/how-to-auth"
	Token             = "<token>"
)

func main() {
	client := nest.New(ClientID, State, ClientSecret, AuthorizationCode)
	client.Token = Token
	structures, err := client.Structures()
	if err != nil {
		fmt.Println(err)
	}
	logEvent(structures, 1)
	for _, structure := range structures {
		err := structure.SetAway(nest.Away)
		if err != nil {
			fmt.Println(err)
		}
	}
	structures, err = client.Structures()
	if err != nil {
		fmt.Println(err)
	}
	logEvent(structures, 2)
}

func logEvent(structures map[string]*nest.Structure, cnt int) {
	fmt.Printf(">>>>>%d<<<<<\n", cnt)
	data, _ := json.MarshalIndent(structures, "", "  ")
	fmt.Println(string(data))
	fmt.Printf(">>>>>%d<<<<<\n", cnt)
}
