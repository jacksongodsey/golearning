package playgame

import (
	"fmt"
	"github.com/jacksongodsey/golearning/A4mk3/names"
	"math/rand"
	"os"
	"time"
)

func Playgame() {
	out := make(chan []names.Baby)
	go names.Boysclass(out)
	boyoutput := <-out
	out2 := make(chan []names.Baby)
	go names.Girlclass(out2)
	girloutput := <-out2
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
		fmt.Printf("Incorrect. There were %v girls named %v, and %v boys named %v.\n", girloutput[random_number].Births, girloutput[random_number].Name, boyoutput[random_number].Births, boyoutput[random_number].Name)
		fmt.Println("Would you like to play again (y/n)?")
		user_input := ""
		fmt.Scanln(&user_input)
		for {
			if user_input == "y" {
				Playgame()
			}
			if user_input == "n" {
				fmt.Println("Goodbye!")
				os.Exit(0)
			} else {
				fmt.Println("Please enter yes or no (y/n).")
				fmt.Scanln(&user_input)
			}
		}
	}
	if user_guess != right_guess {
		fmt.Printf("Incorrect. There were %v girls named %v, and %v boys named %v.\n", girloutput[random_number].Births, girloutput[random_number].Name, boyoutput[random_number].Births, boyoutput[random_number].Name)
		fmt.Println("Would you like to play again (y/n)?")
		user_input := ""
		fmt.Scanln(&user_input)
		for {
			if user_input == "y" {
				Playgame()
			}
			if user_input == "n" {
				fmt.Println("Goodbye!")
				os.Exit(0)
			} else {
				fmt.Println("Please enter yes or no (y/n).")
				fmt.Scanln(&user_input)
			}
		}
	}
}
