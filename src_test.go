package file

import (
	"log"
	"os"
	"sync"
	"testing"
)

func TestOffset(t *testing.T) {
	filePath := "./test/meta.txt"

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 分割成两截
	firstHalf := content[:len(content)/2]
	secondHalf := content[len(content)/2:]
	secondHalfOffset := len(content) / 2

	f, err := GetFile("./test/test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		err = writeFileByOffset(f.File, secondHalf, int64(secondHalfOffset))
		if err != nil {
			log.Fatal(err)
		}
		waitGroup.Done()
	}()
	go func() {
		err = writeFileByOffset(f.File, firstHalf, 0)
		if err != nil {
			log.Fatal(err)
		}
		waitGroup.Done()
	}()
}
