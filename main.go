package main
import 
	("fmt"
	"bufio"
	"os"
	"strings"
)

type cliCommand struct{
	name	    string
	description string
	callback	func() error
}

func getCommands() map[string]cliCommand{
		commands := map[string]cliCommand{
		"exit":{
			name:	"exit",
			description: "Exit the Pokedex",
			callback: commandExit,
			},
			"help":{
			name: 	"help",
			description: "Display a help message",
			callback:	commandHelp, 
		},
	}
	return commands
}

func main(){

fmt.Println("Welcome to PokedexCLI")
scanner := bufio.NewScanner(os.Stdin)
commands :=  getCommands()

for {
	fmt.Print("Pokedex > ")	
	if !scanner.Scan(){
		break
	}

	input := scanner.Text()
	words := strings.Split(input," ")
	command := strings.ToLower(words[0])
	cmd, bol := commands[command]
	if bol{
		cmd.callback()
	}else{
		fmt.Println("Unknow command")
	}
	}
}
