package minitask6

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		file.Close()
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return data, nil
}

func SafeReadFile(path string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi panic:", r)
			fmt.Println("continue...")
		}
	}()

	data, err := ReadFile(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Isi file:")
	fmt.Println(string(data))
}
