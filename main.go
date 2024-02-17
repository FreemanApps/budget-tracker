package main

import (
	"context"
	"fmt"
	"os"
	""
)

func main() {
	fmt.Println("Hello, World!")
	component := hello("John")
	component.Render(context.Background(), os.Stdout)
}