package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/emiguelt/nhlapi/nhlapi"
)

func main() {
	//help to check used time
	now := time.Now()

	rosterFile, err := os.OpenFile("~/limbo/rosters.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		log.Fatalf("error opening file rosters.txt: %v", err)
	}

	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlapi.GetAllTeams()
	if err != nil {
		log.Fatalf("error getting all the teams: %v", err)
	}

	for i, team := range teams {
		log.Println("------------")
		log.Printf("%v -  Name: %s", i, team.Name)
	}

	log.Printf("took %v", time.Now().Sub(now).String())
}
