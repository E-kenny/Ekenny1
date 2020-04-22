package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"go_training/exporter"
	"os"
)

func main() {
	fb := new(exporter.Facebook)
	tw := new(exporter.Twitter)
	li := new(exporter.Linkedin)
	fb = &exporter.Facebook{
		URL:       "http:// www.Ekenny.com/facebook",
		Username:  "E_kenny",
		Followers: 1000,
	}
	tw = &exporter.Twitter{
		URL:       "http:// www.Ekenny.com/twitter",
		Username:  "E_kenny",
		Followers: 5000,
	}
	li = &exporter.Linkedin{
		URL:       "http:// www.Ekenny.com/linkin",
		Username:  "E_kenny",
		Followers: 15000,
	}
	err := export(fb, "fbdata.xml")
	err = export(tw, "twdata.xml")
	err = export(li, "lidata.xml")

	if err != nil {
		panic(err)
	}

}

func export(u exporter.SocialMedia, filename string) error {
	q, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		return errors.New("an error occured while opening file :" + err.Error())
	}

	n, err := xml.MarshalIndent(u, " ", "  ")

	q.Write(n)

	if err != nil {
		return errors.New("an error occured while writing to file :" + err.Error())
	}

	fmt.Println(string(n))

	return nil
}
