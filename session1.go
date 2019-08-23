package main

import (
	"fmt"
)

func printByte(chuoi string) {
	for i := 0; i <= len(chuoi); i++ {
		fmt.Printf("%x ", chuoi[i])
	}
}
func printChar(chuoi string) {
	for i := 0; i <= len(chuoi); i++ {
		fmt.Printf("%c ", chuoi[i])
	}
}
func main() {
	lamnn := "ahihi"
	printByte(lamnn)

	printChar(lamnn)
}
