package main

import( 
	"fmt"
	"bufio"
	"os"
	"github.com/ceyrich/pokedex/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := repl.CleanInput(input)
		fmt.Println("Your command was:", cleaned[0])
	}
}