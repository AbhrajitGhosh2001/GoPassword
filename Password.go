import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Password Generator!")
	fmt.Println("Choose an option:")
	fmt.Println("1. Generate a random password")
	fmt.Println("2. Generate a custom password")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("Enter the length of the random password:")
		lengthInput, _ := reader.ReadString('\n')
		lengthInput = strings.TrimSpace(lengthInput)
		length := 0
		fmt.Sscanf(lengthInput, "%d", &length)

		password := generateRandomPassword(length)
		if password == "" {
			fmt.Println("Failed to generate password. Please try again.")
			return
		}

		fmt.Println("Generated password:", password)
		fmt.Println("Password Strength:")
		fmt.Println(checkPasswordStrength(password))

	case "2":
		fmt.Println("Enter the length of the custom password:")
		lengthInput, _ := reader.ReadString('\n')
		lengthInput = strings.TrimSpace(lengthInput)
		length := 0
		fmt.Sscanf(lengthInput, "%d", &length)

		fmt.Println("Enter the characters you want to include in the password:")
		characters, _ := reader.ReadString('\n')
		characters = strings.TrimSpace(characters)

		password := generateCustomPassword(length, characters)
		if password == "" {
			fmt.Println("Failed to generate password. Please try again.")
			return
		}

		fmt.Println("Generated password:", password)
		fmt.Println("Password Strength:")
		fmt.Println(checkPasswordStrength(password))

	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}
