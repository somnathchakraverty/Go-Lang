package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func createWrite() {
	fmt.Printf("\n Current Unix Time: %v\n", time.Now().Unix())

	time.Sleep(50 * time.Second)
	current := time.Now()
	fileName := current.String()
	file, err := os.Create(fileName + ".txt")
	if err != nil {
		log.Fatal("Error while creating file", err.Error())
	}
	defer file.Close()
	var fileContent string
	fileContent = " "

	for i := 0; i < 100; i++ {
		fileContent = fileContent + "Hello World with value of i : " + strconv.Itoa(i) + " \n"
	}

	fmt.Fprintf(file, fileContent)

	fmt.Printf("\n Current Unix Time: %v\n", time.Now().Unix())

}
