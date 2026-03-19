package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type WordFreq struct {
	word string
	freq int
}

func ErrMassage(message string) error {
	if message == "" {
		return errors.New("invalid input")
	}
	return errors.New(message)
}

func DebugPrint(pairs []WordFreq) {
	for i := 0; i < len(pairs)-1; i++ {
		fmt.Printf("[%d]-[%s][%d]\n", i, pairs[i].word, pairs[i].freq)
	}
}
func SortLexicographically(pairs []WordFreq, i, j int) bool {
	if pairs[i].freq == pairs[j].freq {
		return pairs[i].word < pairs[j].word
	}
	return pairs[i].freq > pairs[j].freq
}

func SortWordCount(WordCountMap map[string]int, globalcount int) string {
	pairs := make([]WordFreq, 0, len(WordCountMap))
	for w, f := range WordCountMap {
		pairs = append(pairs, WordFreq{w, f})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return SortLexicographically(pairs, i, j)
	})
	var builder strings.Builder
	for i := 0; i < globalcount && i < len(pairs); i++ {
		builder.WriteString(pairs[i].word)
		if i < globalcount-1 && i < len(pairs)-1 {
			builder.WriteString(" ")
		}
	}
	return builder.String()
}

func HandlerWord(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)

	if !scanner.Scan() {
		err := ErrMassage("")
		return "", err
	}
	line1 := scanner.Text()
	WordArray := strings.Fields(line1)
	WordCountMap := make(map[string]int)
	for _, v := range WordArray {
		WordCountMap[v]++
	}
	if !scanner.Scan() {
		return "", errors.New("invalid input")
	}
	numStr := scanner.Text()
	if numStr == "" {
		return "", errors.New("invalid input")
	}
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return "", errors.New("invalid input")
	}
	if len(WordCountMap) == 0 {
		return "", nil
	}
	return SortWordCount(WordCountMap, num), nil

}

func main() {

	if str, err := HandlerWord(os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	} else {
		fmt.Print(str)
	}
}
