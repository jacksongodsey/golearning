package getword

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
)

func Getword(out chan string) {
	file, e := os.Open("words.txt")
	if e != nil {
		panic(e)
	}
	slice := make([]string, 0)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		slice = append(slice, str)
	}
	random_number := rand.Intn(2999)
	word := slice[random_number]
	fmt.Println(word)
	out <- word
	close(out)
}
