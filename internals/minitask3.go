package internals

import "fmt"

func Slice() {
	// slice awal
	data := []int{50, 75, 66, 20, 32, 90}

	// cari index angka 66
	index := -1
	for i, v := range data {
		if v == 66 {
			index = i
			break
		}
	}

	// sisipkan 88 setelah 66
	if index != -1 {
		data = append(data[:index+1], append([]int{88}, data[index+1:]...)...)
	}

	// tampilkan isi slice
	for i, v := range data {
		fmt.Printf("Index %d = %d\n", i, v)
	}
}