package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	saysomething := doctor.Intro()
	fmt.Println(saysomething)

	for {
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)

		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}

}

// func sayhello(saysomthing string) {
// 	fmt.Println(saysomthing)
// }
