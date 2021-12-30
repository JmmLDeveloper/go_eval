package main

import (
	"github.com/JmmLDeveloper/go_eval/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {


	fmt.Println("type a valid mathematic expression after the > and press enter to see the result!")
	fmt.Println("you can also type q to stop the program")


	loop:
	for {
		fmt.Print("> ")
		in := bufio.NewReader(os.Stdin)
		input,err := in.ReadString('\n')
		if err != nil {
			continue
		}
		input = strings.TrimSpace(input)
		switch input {
		case "":
			continue
		case "q":
			break loop
		default:
			result,err := utils.Evaluate(input)

			if err != nil {
				fmt.Printf("%v\n", err )
			} else {
				fmt.Printf("%v\n", result )
			}
		}
	}

	fmt.Print("until next time!")

}