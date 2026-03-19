package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReturnError(str string) error {
	if str == "" {
		return errors.New("invalid input")
	}
	return errors.New(str)
}

func ConvArray(slice []string) ([]int, error) {
	arrint := make([]int, len(slice))
	for i, c := range slice {
		var err error
		arrint[i], err = strconv.Atoi(string(c))
		if err != nil {
			return nil, ReturnError("")
		}
	}
	return arrint, nil

}

func ReConvArray(array []int) string {
	found := false
	var builder strings.Builder
	for i, v := range array {
		for j := 0; j < i; j++ {
			if v == array[j] {
				found = true
			}
		}
		if found {
			found = false
			continue
		}
		if i < len(array) && i != 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(strconv.Itoa(v))
	}
	return builder.String()
}

func SlicesHandler(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	if !scanner.Scan() {
		return "", ReturnError("")
	}
	slice_1 := strings.Fields(scanner.Text())
	if len(slice_1) == 0 {
		return "", ReturnError("Empty input")
	}
	slice_1_int, err := ConvArray(slice_1)
	if err != nil {
		return "", err
	}

	if !scanner.Scan() {
		return "", ReturnError("")
	}
	slice_2 := strings.Fields(scanner.Text())
	if len(slice_2) == 0 {
		return "", ReturnError("Empty input")
	}
	slice_2_int, err := ConvArray(slice_2)
	if err != nil {
		return "", err
	}

	out := []int{}

	for _, cl1 := range slice_1_int {
		for _, cl2 := range slice_2_int {
			if cl2 == cl1 {
				out = append(out, cl1)
				break
			}
		}
	}

	if len(out) == 0 {
		return "", ReturnError("Empty intersection")
	}

	return ReConvArray(out), nil
}

func main() {
	if str, err := SlicesHandler(os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Print(str)
	}
}
