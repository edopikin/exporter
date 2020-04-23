package main

import (
	"fmt"
	"go_classes/go_classes/Interfaceclass"
	"go_classes/go_classes/facebook"
	"go_classes/go_classes/linkedin"
	"go_classes/go_classes/twitter"
	"os"
	"go_classes/go_classes/exporter"
)

func main() {
	// Prompt the user to select file format they want

	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkdin := new(linkedin.Linkedin)


	// if format is .txt
	err :=  export(fb, "fbdata.txt")
	err = export(twtr, "twtrdata.txt")
	err = export(lnkdin, "lnkdindata.txt")

	// else if format is .json

	// else if format is .xml

	// else if format is .yaml


	//else if format is invalid (not .txt, .xml, .yml, .json)
	//  error = "Invalid format selected"

	if err != nil {
		panic(err)
	}

	//Scrollfeeds(fb, twtr, lnkdin)
}

// "ScrollFeeds prints all social media feeds"
//func Scrollfeeds(platforms ...interfaceclass.SocialMedia) {
//	for _, sm := range platforms {
//		for _, fd := range sm.Feed() {
//			fmt.Println(fd)
//		}
//		fmt.Println("===========")
//	}
//}

func export(u interfaceclass.SocialMedia, filename string) error {
	f, err:= os.OpenFile(filename, os.O_CREATE | os.O_WRONLY, 0755)
	if err != nil{
		return err
	}
	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil{
			return err
		}
		fmt.Printf("Wrote %d bytes\n", n)
	}
	return nil
}