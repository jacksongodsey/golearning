package names

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type BabyBoy struct {
	Rank   int64
	Name   string
	Births int64
}

func getinfoboys() []string {
	file, e := os.Open("../boynames2015.txt")
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

func Boysclass() []BabyBoy {
	temp := make([]BabyBoy, 0)
	slice := getinfoboys()
	for _, str := range slice {
		t := strings.Trim(str, "\n")
		s := strings.Split(t, " ")
		rank, err1 := strconv.ParseInt(s[0], 10, 64)
		if err1 != nil {
			fmt.Println("there seems to have been an error")
		}
		name := s[1]
		births, err2 := strconv.ParseInt(s[2], 10, 64)
		if err2 != nil {
			fmt.Println("there seems to have been an error")
		}
		pals := BabyBoy{rank, name, births}
		temp = append(temp, pals)
	}
	return temp
}

type BabyGirl struct {
	Rank   int64
	Name   string
	Births int64
}

func getinfogirl() []string {
	file, e := os.Open("../girlnames2015.txt")
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

func Girlclass() []BabyGirl {
	temp := make([]BabyGirl, 0)
	slice := getinfogirl()
	for _, str := range slice {
		t := strings.Trim(str, "\n")
		s := strings.Split(t, " ")
		rank, err1 := strconv.ParseInt(s[0], 10, 64)
		if err1 != nil {
			fmt.Println("there seems to have been an error")
		}
		name := s[1]
		births, err2 := strconv.ParseInt(s[2], 10, 64)
		if err2 != nil {
			fmt.Println("there seems to have been an error")
		}
		pals := BabyGirl{rank, name, births}
		temp = append(temp, pals)
	}
	return temp
}
