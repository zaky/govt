// vtIpReport - fetches information about a given IP from VirusTotal.
//  vtIpReport -ip=8.8.8.8
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/williballenthin/govt"
)

var apikey string
var apiurl string
var ip string

// init - initializes flag variables.
func init() {
	flag.StringVar(&apikey, "apikey", os.Getenv("VT_API_KEY"), "Set environment variable VT_API_KEY to your VT API Key or specify on prompt")
	flag.StringVar(&apiurl, "apiurl", "https://www.virustotal.com/vtapi/v2/", "URL of the VirusTotal API to be used.")
	flag.StringVar(&ip, "ip", "", "ip sum of a file to as VT about.")
}

// check - an error checking function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()
	if ip == "" {
		fmt.Println("-ip=<ip> missing!")
		os.Exit(1)
	}
	c, err := govt.New(govt.SetApikey(apikey), govt.SetUrl(apiurl))
	check(err)

	r, err := c.GetIpReport(ip)
	check(err)
	j, err := json.MarshalIndent(r, "", "    ")
	check(err)
	fmt.Printf("IP Report: ")
	os.Stdout.Write(j)
}
