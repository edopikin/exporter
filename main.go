package exporter

import(
	"go_classes/go_classes/Interfaceclass"
	"fmt"
	"encoding/json"
	"encoding/xml"
	//"encoding/yaml"
	"os"
)

// Export feed
func Export(u interfaceclass.SocialMedia, filename string, format string) error {
	filename = filename+"."+format // fbdata, txt ==> fbdata.txt
	f, err:= os.OpenFile(filename, os.O_CREATE | os.O_WRONLY, 0755)
	if err != nil{
		return err
	}

	if format == "txt"{
		for _, fd := range u.Feed() {
			n, err := f.Write([]byte(fd + "\n"))
			if err != nil{
				return err
			}
			fmt.Printf("Wrote %d bytes\n", n)
		}
	}else if format == "json"{
		for n, fd := range u.Feed() {
			js := map[int]string{n:fd}
			json, _ := json.Marshal(js)
			n, err := f.Write([]byte(string(json)))
			if err != nil{
				return err
			}
			fmt.Printf("Wrote %d bytes\n", n)
		}
	}else if format == "xml"{
		for _, fd := range u.Feed(){
			xml, _ := xml.MarshalIndent(fd,"","")
			n, err := f.Write([]byte(string(xml)))
			if err != nil{
				return err
			}
			fmt.Printf("Wrote %d bytes\n", n)
		}
	}else if format == "yaml"{
	}
	return nil
}

package main

import (
	"fmt"
	//"go_classes/go_classes/Interfaceclass"
	"go_classes/go_classes/exporter"
	"go_classes/go_classes/facebook"
	"go_classes/go_classes/linkedin"
	"go_classes/go_classes/twitter"
)

func main() {
	// Prompt the user to select file format they want
	fmt.Println(`
	Input your designated file format index:
	1. ".txt"
	2. ".json"
	3. ".xml"
	4. ".yaml"
	`)
	var index int
	fmt.Scan(&index)



	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkdin := new(linkedin.Linkedin)

	// 1 = .txt
	if index == 1 {
		format := "txt"
		err := exporter.Export(fb, "fbdata", format)
		err = exporter.Export(twtr, "twtrdata", format)
		err = exporter.Export(lnkdin, "lnkdindata", format)

		if err != nil {
			panic(err)
		}
	// 2 = .json
	}else if index == 2 {
		format := "json"
		err := exporter.Export(fb, "fbdata", format)
		err = exporter.Export(twtr, "twtrdata", format)
		err = exporter.Export(lnkdin, "lnkdindata", format)

		if err != nil {
			panic(err)
		}
	// 3 = .xml
	}else if index == 3 {
		format := "xml"
		err := exporter.Export(fb, "fbdata", format)
		err = exporter.Export(twtr, "twtrdata", format)
		err = exporter.Export(lnkdin, "lnkdindata", format)

		if err != nil {
			panic(err)
		}
	// 4 = .yaml
	}// else if index == 4 {
		//format := yaml
		//err := exporter.Export(fb, "fbdata", format)
		//err = exporter.Export(twtr, "twtrdata", format)
		//err = exporter.Export(lnkdin, "lnkdindata", format)

		//if err != nil {
		//	panic(err)
		//}

	//else if format is invalid (not .txt, .xml, .yml, .json)
	//  error = "Invalid format selected"
}
