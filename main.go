package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"log"
)

func main(){
	var running bool
	var admin bool
	var pass_input string
	var processresult bool
	running = true

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to RANDRACAR's library system!")
	fmt.Print("If you are an admin, please input your password here: ")
	scanner.Scan()
	
	pass_input = scanner.Text()
	hashedinput := encrypt_string(pass_input)
	admin = check_password(hashedinput)

	for running == true {
		if admin == true { 
			fmt.Println("Welcome admin! Your options: ")
			fmt.Println("1 - Search for book")
			fmt.Println("2 - Add a book")
			fmt.Println("3 - Remove a book")
			fmt.Println("4 - Quit")
			fmt.Print("Input your option here: ")
			scanner.Scan()
			input, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			processresult = process_input(input, admin)
			if processresult == false {
				break;
			}
		} else {
			fmt.Println("Welcome user! Your options: ")
			fmt.Println("1 - Search for book")
			fmt.Println("2 - Quit")
			fmt.Print("Input your option here: ")
			scanner.Scan()
			input, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			processresult = process_input(input, admin)
			if processresult == false {
				break;
			}
		}
	}
}