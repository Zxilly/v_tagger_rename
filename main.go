package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func getFiles(filePath string) []os.FileInfo {
	resultFiles := make([]os.FileInfo, 0)
	files, _ := ioutil.ReadDir(filePath)
	for _, i := range files {
		if !i.IsDir() {
			resultFiles = append(resultFiles, i)
		}
	}
	return resultFiles
}

func rename(file os.FileInfo) (bool, error) {
	if !HasSuffix(file.Name(), ".mp4") {
		return false, errors.New("not mp4")
	}

	fileObject, err := os.Open(file.Name())
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 200*1024)

	_, err = fileObject.Read(buffer)

	if err != nil {
		log.Fatal(err)
	}

	fileObject.Close()

	hasher := md5.New()
	hasher.Write(buffer)
	hash := hex.EncodeToString(hasher.Sum(nil))

	filename := hash + ".mp4"

	err = os.Rename("./"+file.Name(), "./"+filename)

	if err != nil {
		panic(err)
	}

	return true, nil
}

func main() {
	fmt.Println("开始处理")
	count := 0
	for _, i := range getFiles("./") {
		_, err := rename(i)
		if err == nil {
			count++
		}
	}
	fmt.Printf("处理了 %v 个文件\n", count)
	time.Sleep(3 * time.Second)
	fmt.Println("quit")
}
