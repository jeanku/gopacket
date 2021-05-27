package gosql

import "fmt"

// say Hi to someone
func SayHi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func Where(id string, value string) {
	println("call where", id, value)
}