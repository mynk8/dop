package head

import (
	"log"
	"os"
	"bufio"
	"fmt"
)

func Head(filename string, n int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() && count < n {
		fmt.Println(scanner.Text())
		count++
	}
}


