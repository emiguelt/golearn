package pass

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func CrackSha1Hash(hashToCrack string, useSalt bool) string {
	salts := []string{}
	if useSalt {
		salts = readSaltFile()
	}

	for _, pass := range readPassFile() {
		if !useSalt {
			if isPassword(hashToCrack, pass) {
				return pass
			}
		} else {
			for _, salt := range salts {
				if isPassword(hashToCrack, salt+pass) {
					return pass
				} else if isPassword(hashToCrack, pass+salt) {
					return pass
				}
			}
		}
	}

	return "PASSWORD NOT IN DATABASE"
}

func isPassword(hashToCrack string, valToHash string) bool {
	return hashToCrack == sha1Hash(valToHash)
}

func sha1Hash(str string) string {
	bs := sha1.Sum([]byte(str))
	return fmt.Sprintf("%x", bs)
}

func readPassFile() []string {
	return readFile("top-10000-passwords.txt")
}

func readSaltFile() []string {
	return readFile("known-salts.txt")
}

func readFile(file string) []string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file: %s", file)
	}

	contents := string(b)
	return strings.Split(contents, "\n")
}
