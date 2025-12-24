package filereader

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) (string, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
			fmt.Println("continue...")
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer func() {
		fmt.Println("file closed")
		file.Close()
	}()

	// info, err := file.Stat()
	// info, err := os.Open(file.Name())
	// if err != nil {
	// 	panic("gagal membaca info file")
	// }
	// if info.IsDir() {
	// 	panic("path mengarah ke directory, bukan file")
	// }

	// data, err := os.ReadFile(path)
	data, err := io.ReadAll(file)
	if err != nil {
		panic("gagal membaca konten file")
	}

	return string(data), nil
}
