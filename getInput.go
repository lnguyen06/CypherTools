
func getInput() (toEncrypt bool, encoding string, message string) {
	fmt.Print("\nWelcome to the Cypher Tool!\n")

	reader := buffio.NewReader(os.Stdin)

	fmt.Print("Select operation (1/2):\n")
	fmt.Print("1. Encrypt.\n")
	fmt.Print("2. Decrypt.\n")
	op, _ := reader.ReadString('\n')
	operation := strings.TrimSpace(op)
	for {
		switch operation {
		case "1": 
			toEncrypt = true
		case "2":
			toEncrypt = false 
		default:
			fmt.Println("Wrong input! Please enter number 1 or 2")
		}
	}
	
	fmt.Print("\nSelect cypher (1/3):\n")
	fmt.Print("1. ROT13.\n")
	fmt.Print("2. Reverse.\n")
	fmt.Print("3. Vinegre.\n")
	cp, _ := reader.ReadString('\n')
	encoding := strings.TrimSpace(cp)
	for {
		switch encoding {
		case "1": 
			 = 
		case "2":
			toEncrypt = false 
		default:
			fmt.Println("Wrong input! Please enter number 1, 2 or 3")
	}
	}

	fmt.Print("\nEnter the message:")
	input, _ := reader.ReadString('\n')
	message := strings.TrimSpace(input)
	for len(message) == 0 {
		fmt.Println("Message could not be empty. Please enter the message!")
	}
}