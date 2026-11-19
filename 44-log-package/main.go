package main

import (
	"errors"
	"fmt"
	"log"
)


func loadConfig() error {
	return errors.New("Config file not found")
}

// func connect() error {
// 	return errors.New(" database not reachable")
// }

func loadUser() error {
    return errors.New("user data missing")
}

func main() {
	log.Println("starting application")
	log.Println("loading configurations")
	log.Println("application running")
	//
	fmt.Println()
	//
	err := loadConfig()
	if err != nil {
		log.Println("Error:", err)
		// return
	}
	// log.Println("Configuration loaded sucessfullly")

	//
	fmt.Println()
	//

	// err2 := connect()
	// if err2 != nil {
	// 	log.Fatal("connection failed.", err2)
	// }
	// log.Println("database connection sucessful")


	log.Println("Starting service")

    err3 := loadUser()
    if err3 != nil {
        log.Fatal("Critical failure, shutting down:", err3)
    }

    log.Println("Service started successfully")
}