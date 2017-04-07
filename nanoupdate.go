package main

import "fmt"
import "encoding/json"
import "io/ioutil"

func main()  {
	fmt.Println("NANOUPDATE")
	config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	var config_data map[string]interface{}

	if err := json.Unmarshal(config, &config_data); err != nil {
		panic(err)
	}
	fmt.Println(config_data)
}
