package file_test

import (
	"log"
	"os"
	"sync"
	"testing"

	file "github.com/238Studio/child-nodes-file-service"
)

func TestCreateDir(t *testing.T) {
	err := file.CreateDir("./test")
	if err != nil {
		log.Fatal(err)
	}
}

func TestAppendFile(t *testing.T) {
	filePath := "./test/meta.txt"

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = file.AppendFile("./test/test1.txt", content)
	if err != nil {
		log.Fatal(err)
	}
}

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

	f, err := file.GetFile("./test/test2.txt")
	if err != nil {
		log.Fatal(err)
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		err = f.WriteFileByOffset(secondHalf, int64(secondHalfOffset))
		if err != nil {
			log.Fatal(err)
		}
		waitGroup.Done()
	}()
	go func() {
		err = f.WriteFileByOffset(firstHalf, 0)
		if err != nil {
			log.Fatal(err)
		}
		waitGroup.Done()
	}()

	err = f.CloseFile()
	if err != nil {
		t.Log(err)
	}
}

func TestOffset1(t *testing.T) {
	filePath := "./test/meta.txt"

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 分割成两截
	firstHalf := content[:len(content)/2]
	secondHalf := content[len(content)/2:]
	secondHalfOffset := len(content) / 2

	f, err := file.GetFile("./test/test3.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = f.WriteFileByOffset(secondHalf, int64(secondHalfOffset))
	if err != nil {
		log.Fatal(err)
	}

	err = f.WriteFileByOffset(firstHalf, 0)
	if err != nil {
		log.Fatal(err)
	}

	err = f.CloseFile()
	if err != nil {
		t.Log(err)
	}
}

func TestAppendString(t *testing.T) {
	content := "114514"
	err := file.AppendString("./test/test3.txt", content)
	if err != nil {
		log.Fatal(err)
	}
}

func TestDeleteFile(t *testing.T) {
	err := file.DeleteFile("./test/test2.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func TestDeleteDir(t *testing.T) {
	err := file.DeleteDir("./test")
	if err != nil {
		log.Fatal(err)
	}
}

func TestIsPathExist(t *testing.T) {
	log.Println(file.IsPathExist("./test"))
	log.Println(file.IsPathExist("./test/meta.txt"))
}
