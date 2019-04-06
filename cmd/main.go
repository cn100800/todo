package main

import "os"

func main() {
	f, _ := os.OpenFile("t", os.O_CREATE|os.O_RDWR, 0644)
	f.Truncate(4)
	f.Seek(6, 0)
	f.WriteString("9")
}
