package main

import "fmt"
import "encoding/json"

func main()  {
     fmt.Println("NANOUPDATE")
     somestring := []string{"hello", "world", "yoooo"}
     something, err := json.Marshal(somestring)
     if err != nil {
         panic(err)
     }
     fmt.Println(string(something))
}
