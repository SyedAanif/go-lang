package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while reading file")
		panic(err)
	}
	defer f.Close() // close the file stream

	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("Error while getting file-info")
		panic(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("File size:", fileInfo.Size())
	fmt.Println("File permission:", fileInfo.Mode())
	fmt.Println("File modified at:", fileInfo.ModTime())

	// read file
	// store in buffer
	buff := make([]byte, fileInfo.Size())

	_, err = f.Read(buff)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buff))

	// simpler way, but loads whole data into memory
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data)) // byte to string conversion

	// reads folders
	dir, err := os.Open("../")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	info, err := dir.ReadDir(0)
	if err != nil {
		panic(err)
	}

	for _, fi := range info {
		fmt.Println(fi.Name())
	}

	// create a file
	f, err = os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	f.WriteString("hello")
	f.WriteString(" world") // append

	bytes := []byte("\nhello golang") // byte slice
	f.Write(bytes)

	// dont load everything in memory, stream read and write
	srcFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}

	defer srcFile.Close()

	destFile, err := os.Create("destination.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		err = writer.WriteByte(b)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush() // this writes the data from buffer into destination
	fmt.Println("file copied via streaming!!!")

	// delete a file
	err = os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("file deleted!!!")
}
