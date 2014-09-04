package main

import (
	"../."
	"encoding/json"
	"fmt"
	"os"
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
	structuresChan := make(chan map[string]*nest.Structure)
	go func() {
		client.StructuresStream(func(structures map[string]*nest.Structure, err error) {
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			structuresChan <- structures
		})
	}()

	for i := 0; i < 2; i++ {
		structures := <-structuresChan
		fmt.Println(structures["h68snN..."])
		switch i {
		case 0:
			logEvent(structures, i)
			fmt.Println("Setting away status")
			err := structures["h68snN..."].SetAway(nest.Home)
			if err != nil {
				fmt.Printf("Error: %s - %d\n", err.Description, i)
				os.Exit(2)
			}
		case 1:
			logEvent(structures, i)
			break
		}
	}
}

func logEvent(structures map[string]*nest.Structure, cnt int) {
	fmt.Printf(">>>>>%d<<<<<\n", cnt)
	data, _ := json.MarshalIndent(structures, "", "  ")
	fmt.Println(string(data))
	fmt.Printf(">>>>>%d<<<<<\n", cnt)
}
