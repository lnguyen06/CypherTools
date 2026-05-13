package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	toEncrypt, encoding, message := getInput()
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

	


	
}

