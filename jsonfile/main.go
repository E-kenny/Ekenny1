package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_training/exporter"
	"os"
)

func main() {
	fb := new(exporter.Facebook)
	tw := new(exporter.Twitter)
	li := new(exporter.Linkedin)

	err := export(fb, "fbdata.json")
	err = export(tw, "twdata.json")
	err = export(li, "lidata.json")

	if err != nil {
		panic(err)
	}

}

func export(u exporter.SocialMedia, filename string) error {
	q, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)

	if err != nil {
		return errors.New("an error occured while opening file :" + err.Error())
	}
	n, err := json.Marshal(u.Feed())

	q.Write(n)

	if err != nil {
		return errors.New("an error occured while writing to file :" + err.Error())
	}
	fmt.Println(string(n))

	return nil
}
