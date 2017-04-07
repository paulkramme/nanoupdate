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

func main() {
	config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(config))

	var ver version

	jsonerr := fromjson(string(config), &ver)
	if jsonerr != nil {
		panic(jsonerr)
	}

	fmt.Printf("%s %d.%d.%d\n", ver.Name, ver.Major, ver.Minor, ver.Patch)
}
