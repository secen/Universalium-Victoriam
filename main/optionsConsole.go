package main
import(
	"bufio"
	"fmt"
	"os"
	"strings"
)
func optionsConsoleShowIndex(){
	fmt.Println("Victorium Universalim is a complex simulation of the historical conditions of the Globe during the 14th to 19th centuries.");
	fmt.Println("It is advised that you either play the game in tutorial mode, or that you read this manual beforehand, so that you might be able to interact correctly and tactfully in the world of medieval history.")
	fmt.Println("Index: ")
	fmt.Println("1.Introduction")
	fmt.Println("2.User Interface")
	fmt.Println("3.Quick Tips")
	fmt.Println("4.Combat")
	fmt.Println("5.The Economy")
	fmt.Println("6.Trade")
	fmt.Println("7.Tax Income")
	fmt.Println("8.Administration")
	fmt.Println("9.Production")
	fmt.Println("10.Government")
	fmt.Println("11.Advanced Tips")
	fmt.Println("12. Recommended Reading")
	fmt.Println("0.Back")
}
func showOptions() {

}
func optionsConsoleprintBookRecomendations() {
	fmt.Println("12. RECOMENDED READING")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("The Art Of War By Sun Tsu")
	fmt.Println("Diplomacy By Henry Kissinger")
	fmt.Println("Rules For Radicals by Saul Alinsky")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}
func startOptions() {
	gameState=OPTIONS_MENU;
	optionsConsoleShowIndex();
	handleInputOptionsConsole();
}
func handleInputOptionsConsole() {
	for{
		if gameState!= OPTIONS_MENU {
			break;}
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare(text,"0") == 0{
			gameState = MAIN_MENU;
			main();
			break;
		}
		if strings.Compare(text,"12")==0{
			optionsConsoleprintBookRecomendations()
		}

		}
}
func execManual() {
	optionsConsoleShowIndex()
	handleInputOptionsConsole()
}