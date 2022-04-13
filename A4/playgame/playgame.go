package main

import (
	"A4/names"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func playgame() {
	boyoutput := names.Boysclass()
	girloutput := names.Girlclass()
	rand.Seed(time.Now().UnixNano())
	random_number := rand.Intn(998)
	right_guess := 0
	if girloutput[random_number].Births > boyoutput[random_number].Births {
		right_guess = 1
	} else {
		right_guess = 2
	}
	user_guess := 0
	fmt.Printf("In 2015, was the name %v (1) or %v (2) more popular (enter 1 or 2)?\n", girloutput[random_number].Name, boyoutput[random_number].Name)
	fmt.Scanln(&user_guess)
	if user_guess == right_guess {
		fmt.Println("Correct!")
		fmt.Println("Would you like to play again (y/n)?")
		user_input := ""
		fmt.Scanln(&user_input)
		if user_input == "y" {
			playgame()
		}
		if user_input == "n" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
	if user_guess != right_guess {
		fmt.Printf("Incorrect. There were %v girls named %v, and %v boys named %v.\n", girloutput[random_number].Births, girloutput[random_number].Name, boyoutput[random_number].Births, boyoutput[random_number].Name)
		fmt.Println("Would you like to play again (y/n)?")
		user_input := ""
		fmt.Scanln(&user_input)
		if user_input == "y" {
			playgame()
		}
		if user_input == "n" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}

func main() {
	playgame()
}
