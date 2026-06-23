package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		next := scanner.Scan()
		if next != true {
			fmt.Errorf("There is no token to scan")
			continue
		}
		token := scanner.Text()
		text := cleanInput(token)
		fmt.Println("Your command was:", text[0])
	}
}