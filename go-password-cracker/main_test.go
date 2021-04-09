package main

import (
	"testing"

	"github.com/emiguelt/go-password-cracker/pass"
)

func Test_superman(t *testing.T) {
	assertTest("18c28604dd31094a8d69dae60f1bcd347f1afc5a", "superman", false, t)
}

func TestSalted_superman(t *testing.T) {
	assertTest("53d8b3dc9d39f0184144674e310185e41a87ffd5", "superman", true, t)
}

func Test_NotFound(t *testing.T) {
	assertTest("asdfasdf", "PASSWORD NOT IN DATABASE", false, t)
}

func assertTest(hashToCrack string, expected string, salted bool, t *testing.T) {
	password := pass.CrackSha1Hash(hashToCrack, salted)

	if password == expected {
		t.Log("Superman password cracked")
	} else {
		t.Errorf("Found %s, expected %s", password, expected)
	}
}
