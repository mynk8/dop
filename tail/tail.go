package tail

import (
	"bufio"
	"log"
	"os"
	"fmt"
)

func Tail(filename string, n int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("error reading file")
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) > n {
			lines = lines[1:]
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

