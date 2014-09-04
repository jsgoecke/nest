package main

import (
	"../."
	"encoding/json"
	"fmt"
	"os"
	"time"
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

	for i := 0; i < 3; i++ {
		structures := <-structuresChan
		switch i {
		case 0:
			logEvent(structures, i)
			fmt.Println("Setting away status")
			err := structures["h68sn..."].SetAway(nest.Away)
			if err != nil {
				fmt.Printf("Error: %s - %d\n", err.Description, i)
				os.Exit(2)
			}
		case 2:
			logEvent(structures, i)
			fmt.Println("Setting ETA")
			err := structures["h68sn..."].SetETA("foobar-trip-id", time.Now().Add(10*time.Minute), time.Now().Add(30*time.Minute))
			if err != nil {
				fmt.Printf("Error: %s - %d\n", err.Description, i)
				os.Exit(2)
			}
			logEvent(structures, i)
		case 3:
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
