package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cakeatthewalk-cmyk/pokedexcli/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var cl = pokeapi.NewClient()
	var cfg = config{client: cl, commands: getAvailableCommands()}
	cfg.pokemon = make(map[string]pokeapi.Pokemon)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		lineInput := scanner.Text()

		if strings.TrimSpace(lineInput) != "" {
			allTheWords := cleanInput(lineInput)

			cC, ok := getAvailableCommands()[allTheWords[0]]
			if ok {
				cC.callback(&cfg, allTheWords[1:])
			} else {
				fmt.Println("Unknown command")

			}
		}
	}
}
