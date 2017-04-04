package main

import "fmt"
import "encoding/json"

func main()  {
     fmt.Println("NANOUPDATE")
     somestring := []string{"hello", "world", "yoooo"}
     something, _ := json.Marshal(somestring)
     fmt.Println(string(something))
}
