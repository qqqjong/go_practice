package main

import "fmt"

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			fmt.Println("for")
			return 1

		case a[i] < b[i]:
			return -1
		}

	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func main() {
	a := Compare([]byte{'A', 'A', 'A', 'A'}, []byte{'A', 'A', 'B'})
	fmt.Println(a)
}