package solutions

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadAocInput(day string) ([]string, error) {
	data, err := ioutil.ReadFile("input/input-" + day)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string and print to screen
	text := string(data)
	return strings.Split(text, "\n"), nil
}

func csvRow2intArray(s string) ([]int, error) {
	entries := strings.Split(s, ",")
	numbers := make([]int, 0, len(entries))
	for i, entry := range entries {
		n, err := strconv.Atoi(entry)
		if err != nil {
			return nil, fmt.Errorf("%v: position %v, char '%v'", err, i, n)
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

func removeEmpty(s []string) []string {
	r := make([]string, 0)
	for _, entry := range s {
		if len(strings.Trim(entry, " ")) == 0 {
			continue
		}
		r = append(r, entry)
	}
	return r
}
