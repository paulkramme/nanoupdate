package main

import "fmt"
import "encoding/json"
import "io/ioutil"

type version struct {
	Name string
	Major int
	Minor int
	Patch int
	Base_download string
}

func fromjson(src string, v interface{}) error {
	return json.Unmarshal([]byte(src), v)
}

func (v version) info() int {
	fmt.Printf("Name: %s\nVersion: %d.%d.%d\nServer: %s\n", v.Name, v.Major, v.Minor, v.Patch, v.Base_download)
	return 0
}

func main() {
	config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var ver version
	jsonerr := fromjson(string(config), &ver)
	if jsonerr != nil {
		panic(jsonerr)
	}
	ver.info()
}
