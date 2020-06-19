package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
var gameState = MAIN_MENU;
func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Victorium Universalim")
	fmt.Println("---------------------")
	fmt.Println("Type 'help' for help")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		inputHandler(text);
	}

}