package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	switch cypher {
	case "1" :
		ROT13(message)
	case "2" :
		Reverse(message)
	case "3" :
		Vinegre(message)
	default:
		fmt.Println("Wrong input! Please enter number 1, 2 or 3")
	}

	switch operation {
	case "1": 
		fmt.Println("\nEncrypted message using ")
	case "2" :
		fmt.Println("\nDecrypted message using")
	}


	
}

