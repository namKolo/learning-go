package main

import (
	"fmt"
	"learning-go/pattern/singleton/db"
)

func main() {
	repo := db.Repository()
	repo.Set("hello", "my friend")
	fmt.Println(repo.Get("hello"))
}
