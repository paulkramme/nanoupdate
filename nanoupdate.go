package main

import "fmt"
import "encoding/json"
import "io"
import "io/ioutil"
import "os"
import "net/http"

// version struct is used to hold information from json configuration files.
type version struct {
	Name          string
	Major         int
	Minor         int
	Patch         int
	Base_download string
	Bin_location  string
	Conf_location string
	Httptype      string
}

// fromjson parses json from an string and returnes to an array.
func fromjson(src string, v interface{}) error {
	return json.Unmarshal([]byte(src), v)
}

// download downloads a file from a specific location URL "link" to a specific disk location "path".
func download(link string, path string) (err error) {
	out, err := os.Create(path)
	if err != nil {
		return
	}
	defer out.Close()
	resp, err := http.Get(link)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return
	}
	return
}

// info prints information from the version struct.
func (v version) info() {
	fmt.Printf("Name: %s\nVersion: %d.%d.%d\nServer: %s\n", v.Name, v.Major, v.Minor, v.Patch, v.Base_download)
}

func main() {
	config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var ver version
	err = fromjson(string(config), &ver)
	if err != nil {
		panic(err)
	}
	ver.info()
	link := fmt.Sprintf("%s://%s/%d_%d/latest/%s*", ver.Httptype, ver.Base_download, ver.Major, ver.Minor, ver.Name)
	fmt.Println(link)
	os.Rename(ver.Bin_location, fmt.Sprintf("%s_old", ver.Bin_location))
	err = download(link, ver.Bin_location)
	if err != nil {
		fmt.Println("Download failed, more information in crash report.")
		os.Rename(fmt.Sprintf("%s_old", ver.Bin_location), ver.Bin_location)
		fmt.Println("Reverted old version. Crashing...")
		panic(err)
	}
	os.Remove(fmt.Sprintf("%s_old", ver.Bin_location))
}
