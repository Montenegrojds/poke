package main
import 
	("fmt"
	"bufio"
	"os"
)

func main(){

fmt.Println("wellcome to PokedexCLI")
scanner := bufio.NewScanner(os.Stdin)


for {
	fmt.Print("Pokedex > ")	
	if !scanner.Scan(){
		break
	}
	input := scanner.Text()
	fmt.Println(input)
}}
