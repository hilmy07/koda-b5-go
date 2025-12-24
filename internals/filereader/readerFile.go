package filereader

import (
	"fmt"
)

func ReaderFile() {
	fmt.Println("TEST 1: valid file")
	content, err := ReadFile("test.txt")

	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("content:", content)
	}

	fmt.Println("---------------")

	fmt.Println("TEST 2: invalid file")
	_, err = ReadFile("tidak_ada.txt")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("---------------")

	fmt.Println("TEST 3: directory path")
	_, _ = ReadFile("./filereader")
}