package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Read Files
	//
	buf := make([]byte, 10)

	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i++ {
		println("Data:", d, string(buf[i]))
	}

	f, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f))

	// Read Folders

	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}

	defer dir.close()

	fileInfo, err := dir.ReadDir()
	for _, fi := range fileInfo {
		fmt.Println(fi.Name())
	}

	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("Filename:", fileInfo.Name())
	fmt.Println("File Or Folder:", fileInfo.IsDir())
	fmt.Println("File Size:", fileInfo.Size())
	fmt.Println("File Mode:", fileInfo.Mode())
	fmt.Println("File Modified At:", fileInfo.ModTime())

	// Create File
	f, err := os.Create("example_2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("hello")
	f.WriteString("Golang")

	bytes := []byte("Hello Golang")
	f.Write(bytes)

	// Read and Write to another file
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("example_3.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewReader(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		e := writer.WriteByte(b)
		if err != nil {
			panic(e)
		}
	}
	writer.Flush()
	fmt.Println("Success...")

	// Delete file

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()
	err := os.Remove("example_2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File Deleted...")
}
