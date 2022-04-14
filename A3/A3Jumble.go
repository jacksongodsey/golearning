package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func openfile() []string {
	file, e := os.Open("words.txt")
	if e != nil {
		panic(e)
	}

	slice := make([]string, 0)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}
		slice = append(slice, str)
	}
	return slice
}

func build_dict() map[string][]string {
	dictionary := make(map[string][]string)
	slice := openfile()
	for _, str := range slice {
		f := strings.Trim(str, " ")
		t := strings.Split(f, "")
		sort.Strings(t)
		w := strings.Join(t, "")
		//l := w
		dictionary[w] = append(dictionary[w], f)
		//if reflect.DeepEqual(dictionary[w], l) {
		//	dictionary[w] = append(dictionary[w], f)
		//} else {
		//	dictionary[w] = append(dictionary[w], f)
		//}
	}
	return dictionary
}

func unscrambled(user string) []string {
	output := build_dict()
	temp := strings.Split(user, "")
	sort.Strings(temp)
	sorted := strings.Join(temp, "")
	final := output[sorted]
	return final
}

func playgame() {
	user_input := ""
	fmt.Print("Please enter a jumble to solve: ")
	fmt.Scanln(&user_input)
	done := unscrambled(user_input)
	fmt.Printf("%v unscrambles to %v\n", user_input, done)

	hardcoded := make([]string, 0)
	hardcoded = append(hardcoded, "enost", "asewes", "aaabcs")
	for i := range hardcoded {
		done = unscrambled(hardcoded[i])
		fmt.Printf("%v unscrambles to %v\n", hardcoded[i], done)
	}
}

func main() {
	playgame()
}
