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
		Facebookfeed: []string{
			"This is start.NG internship", "I know you know it", "but I have to tell you"},
	}
	tw = &exporter.Twitter{
		URL:         "http:// www.Ekenny.com/twitter",
		Username:    "E_kenny",
		Followers:   5000,
		Twitterfeed: []string{"I know you know I need the money", "Thanks to the Cyberelf himself", "You are the best and I am happy being here"},
	}
	li = &exporter.Linkedin{
		URL:          "http:// www.Ekenny.com/linkin",
		Username:     "E_kenny",
		Followers:    15000,
		Linkedinfeed: []string{"I know you know I need the money", "Time for business", "I love my Job"},
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
