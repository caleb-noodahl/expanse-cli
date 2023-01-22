package utils

import (
	"io/ioutil"
	"log"
)

func ReadJSONFile(path string) []byte {
	out, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return out
}
