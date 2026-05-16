package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Main Function
//Prompt the user to enter two integers.
//Read the integers from the user input.
//Compute the sum of the two integers.
//Display the result to the user.
//this comment represents a change for GIT

func main() {

	var sA, sB string
	var iA, iB int
	var errA, errB error

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the first integer here:")
	scanner.Scan()
	sA = strings.TrimSpace(scanner.Text())
	fmt.Print("Enter the second integer here:")
	scanner.Scan()
	sB = strings.TrimSpace(scanner.Text())

	iA, errA = strconv.Atoi(sA)
	iB, errB = strconv.Atoi(sB)

	if errA != nil || errB != nil {
		fmt.Println("Error: inputs not integers!")
	} else {
		fmt.Printf("%d + %d = %d", iA, iB, (iA + iB))
	}
}