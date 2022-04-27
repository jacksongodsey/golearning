package names

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Baby struct {
	Rank   int64
	Name   string
	Births int64
}

func getinfoboys() []string {
	file, e := os.Open("boynames2015.txt")
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
	return slice
}

func Boysclass(out chan []Baby) {
	temp := make([]Baby, 0)
	slice := getinfoboys()
	for _, str := range slice {
		t := strings.Trim(str, "\n")
		s := strings.Split(t, " ")
		rank, err1 := strconv.ParseInt(s[0], 10, 64)
		if err1 != nil {
			fmt.Println("there seems to be an error")
		}
		name := s[1]
		births, err2 := strconv.ParseInt(s[2], 10, 64)
		if err2 != nil {
			fmt.Println("there seems to be an error")
		}
		pals := Baby{rank, name, births}
		temp = append(temp, pals)
	}
	out <- temp
	close(out)
}

func getinfogirls() []string {
	file, e := os.Open("girlnames2015.txt")
	if e != nil {
		panic(e)
	}
	slice := make([]string, 0)
	reader := bufio.NewReader(file)
	for i := 0; i < 1001; i++ {
		str, err := reader.ReadString('\n')
		slice = append(slice, str)
		if err == io.EOF {
			break
		}
	}
	return slice
}

func Girlclass(out2 chan []Baby) {
	temp := make([]Baby, 0)
	slice := getinfogirls()
	for _, str := range slice {
		t := strings.Trim(str, "\n")
		s := strings.Split(t, " ")
		rank, err1 := strconv.ParseInt(s[0], 10, 64)
		if err1 != nil {
			fmt.Println("there seems to be an error")
		}
		name := s[1]
		births, err2 := strconv.ParseInt(s[2], 10, 64)
		if err2 != nil {
			fmt.Println("there seems to be an error")
		}
		pals := Baby{rank, name, births}
		temp = append(temp, pals)
	}
	out2 <- temp
	close(out2)
}
