package utility

import (
	"boardgame-tracker/src/pkg/documents"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetHotGames(gameId string) documents.BoardgamesXML {
	//to not violate ToS
	time.Sleep(5 * time.Second)
	resp, err := http.Get(fmt.Sprintf("https://boardgamegeek.com/xmlapi/boardgame/%s", gameId))
	if err != nil {
		log.Printf("error : %e ", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var bg documents.BoardgamesXML

	xmlErr := xml.Unmarshal(body, &bg)

	if xmlErr != nil {
		log.Printf("error %e", xmlErr)
	}
	log.Printf("Fetched game :  %s | ID : %s", bg.BoardgameXML.Name, bg.BoardgameXML.Objectid)
	log.Printf("Fetched game :  %s | ID : %s", bg.BoardgameXML.Name[0].Text, bg.BoardgameXML.Name[0].Primary)
	log.Printf("Image :  %s", bg.BoardgameXML.Image)
	CreateSimpleBoardGameStruct(bg)
	return bg
}

func CreateSimpleBoardGameStruct(bgXML documents.BoardgamesXML) Boardgame {
	var simpleBG Boardgame
	boardgameXML := bgXML.BoardgameXML
	simpleBG.Image = boardgameXML.Image
	minPlayers, minPlayError := strconv.Atoi(boardgameXML.Minplayers)

	if minPlayError != nil {
		log.Printf("Error converting minplayers game with ID :%d , error : %e", boardgameXML.Minplayers, minPlayError)
	}
	maxPlayers, maxPlayError := strconv.Atoi(boardgameXML.Maxplayers)

	if maxPlayError != nil {
		log.Printf("Error converting minplayers game with ID :%d , error : %e", boardgameXML.Maxplayers, maxPlayError)
	}
	simpleBG.MinPlayers = minPlayers
	simpleBG.MaxPlayers = maxPlayers
	simpleBG.MinPlaytime = boardgameXML.Minplaytime
	simpleBG.Description = boardgameXML.Description.Text
	for _, elem := range boardgameXML.Name {
		if elem.Primary == "true" {
			simpleBG.Name = elem.Text

		}
	}
	log.Printf("\n %o", simpleBG)
	return simpleBG
}

type Boardgame struct {
	Image       string
	MinPlaytime string
	MinPlayers  int
	MaxPlayers  int
	Description string
	Name        string
}
