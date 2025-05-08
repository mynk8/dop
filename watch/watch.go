package watch

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"time"
	"io"
)

func WatchWithTimestamps(filename string) {
	var offset int64
	for {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		file.Seek(offset, 0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05"), scanner.Text())
		}
		offset, _ = file.Seek(0, io.SeekCurrent)
		file.Close()
		time.Sleep(1 * time.Second)
	}
}
