func getInput() (toEncrypt bool, encoding string, message string) {
	fmt.Print("\nWelcome to the Cypher Tool!\n")

	reader := buffio.NewReader(os.Stdin)

	fmt.Print("Select operation (1/2):\n")
	fmt.Print("1. Encrypt.\n")
	fmt.Print("2. Decrypt.\n")
	op, _ := reader.ReadString('\n')
	operation := strings.TrimSpace(op)

	fmt.Print("\nSelect cypher (1/3):\n")
	fmt.Print("1. ROT13.\n")
	fmt.Print("2. Reverse.\n")
	fmt.Print("3. Vinegre.\n")
	cp, _ := reader.ReadString('\n')
	cypher := strings.TrimSpace(cp)

	fmt.Print("\nEnter the message:")
	message, _ := reader.ReadString('\n')
	encoding := strings.TrimSpace(input)
}