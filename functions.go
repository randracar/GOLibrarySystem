package main

/* TODO:
- Arrumar o quit
- fazer a main function
*/

import(
	"bufio"
	"crypto/md5"
	"io/ioutil"
	"encoding/hex"
	"log"
	"fmt"
	"strings"
	"embed"
	"os"
)

//go:embed hash.txt
var hashfile embed.FS


var booksfile string

func encrypt_string(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func check_password(p string) bool {
	data, err := hashfile.ReadFile("hash.txt")

	if err != nil {
		log.Fatal(err)
	}

	encodedContent := string(data)
	if p == encodedContent { 
		return true
	} else {
		return false
	}
}

func split_bookstring(s string) []string {
	res := strings.Split(s, "|")
	return res
}

func search_book() {
	booksfile = "data/books.txt"
	scanner := bufio.NewScanner(os.Stdin)
	var s string
	readFile, err := os.Open(booksfile)
	var book Book

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Input a book title, author or ISBN to search for: ")
	scanner.Scan()

	s = scanner.Text()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), s){
			var booktext string = fileScanner.Text()
			bookinfo := split_bookstring(booktext)
			book.Title = bookinfo[0]
			book.Author = bookinfo[1]
			book.Isbn = bookinfo[2]
			fmt.Println("Found the book you were looking for!")
			fmt.Print("Book title: ", book.Title, "\n")
			fmt.Print("Book author: ", book.Author, "\n")
			fmt.Print("Book ISBN: ", book.Isbn, "\n")
		} else {
			fmt.Println("This is not the book...")
		}
	}
	fmt.Println("Went through the whole book list. If nothing was found, then there isn't a book with those search parameters in the library.")
}

func remove_book() {
	booksfile = "data/books.txt"
	readFile, err := os.Open(booksfile)
	scanner := bufio.NewScanner(os.Stdin)
	filescanner := bufio.NewScanner(readFile)
	var bookinfo string
	var string_append string
	string_append = ""

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Input a book title, author or ISBN to remove the book: ")
	scanner.Scan()
	bookinfo = scanner.Text()

	for filescanner.Scan() {
		if strings.Contains(filescanner.Text(), bookinfo) {
			// dont append the text into the string, so do nothing
			fmt.Println("Removing: ", filescanner.Text())
		} else {
			if string_append == "" {
				string_append = filescanner.Text()
			} else {
				string_append = fmt.Sprintf("%s\n%s", string_append, filescanner.Text())
			}
		}
	}

	data := []byte(string_append)
	err = ioutil.WriteFile(booksfile, data, 0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully removed the book from the list of books.")
	}
}

func add_book() {
	booksfile = "data/books.txt"
	scanner := bufio.NewScanner(os.Stdin)
	var book Book
	
	fmt.Print("Please input the title of the book: ")
	scanner.Scan()
	book.Title = scanner.Text()
	fmt.Print("\nPlease input the author of the book: ")
	scanner.Scan()
	book.Author = scanner.Text()
	fmt.Print("\nPlease input the book ISBN: ")
	scanner.Scan()
	book.Isbn = scanner.Text()

	bookstring := fmt.Sprintf("\n%s|%s|%s", book.Title, book.Author, book.Isbn)
	fmt.Println(bookstring)

	content, err := ioutil.ReadFile(booksfile)

	if err != nil {
		log.Fatal(err)
	}

	encodedContent := string(content)

	var data1 string
	data1 = fmt.Sprintf("%s%s", encodedContent, bookstring)

	data := []byte(data1)

	err = ioutil.WriteFile(booksfile, data, 0)

	if err != nil {
		log.Fatal(err)
	}
}

func process_input(i int, admin bool) bool {
	if admin == true {
		if i == 1 {
			// search book
			search_book()
		} else if i == 2 {
			// add book
			add_book()
		} else if i == 3 {
			// remove book
			remove_book()
		} else {
			// quit
			fmt.Println("Quitting the program...")
			return false
		}
	} else { 
		if i == 1 {
			// search book
		} else {
			// quit
			fmt.Println("Quitting the program...")
			return false
		}
	}
	return true
}

type Book struct {
	Title string
	Author string
	Isbn string
}

