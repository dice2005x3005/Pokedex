package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"github.com/dice2005x3005/Pokedex/internal/pokecache"
)

func repl(c *Config, cache *pokecache.Cache, u *User) {
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
		if len(text) == 0 {
			continue
		}
		commandName := text[0]
		loc := []string{}
		if len(text) > 1 {
			loc = text[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(c, cache, u, loc...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}


func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	cut := strings.Fields(text)
	return cut
}

type cliCommand struct {
	name string
	description string
	callback func(c* Config, cache *pokecache.Cache, u *User, loc ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name: "map",
			description: "Show 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Show the previous 20 locations",
			callback: commandMapBack,
		},
		"explore": {
			name: "explore",
			description: "Show the pokemon you can encounter in that location",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Let you capture a pokemon by name",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Let you see the stats of a pokemon you have catched",
			callback: commandInspect,
		},
	}
}