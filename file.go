package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func getWriteFile(config Config) (*os.File, error) {
	fpath := config.FileWriteLocation
	if fpath == "" {
		log.Fatal("file write location blank... defaulting to ./sessions.json")
		fpath = "./sessions.json"
	}

	if ext := path.Ext(fpath); ext != ".json" {
		log.Fatal("write file has to be a json file. Your file: ", ext)
	}

	f, err := os.Open(fpath)
	if err != nil {
		f, err = os.Create(fpath)

		if err != nil {
			return nil, err
		}
	}

	fmt.Println(f)

	return f, nil
}
