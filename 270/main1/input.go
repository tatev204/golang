package main1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputArray(N int) []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Numbers:")
	str, err := reader.ReadString('\n')
	CheckError(err)

	var values []int
	str = strings.TrimSpace(str)
	lines := strings.Fields(str)
	for _, line := range lines {
		val, err := strconv.Atoi(line)
		CheckError(err)
		values = append(values, val)
	}
	if len(values) != N {
		panic("Overflow")
	}
	return values
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
