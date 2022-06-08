package utility

import (
	"boardgame-tracker/src/pkg/documents"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetHotGames() string {
	//to not violate ToS
	time.Sleep(5 * time.Second)
	resp, err := http.Get("https://boardgamegeek.com/xmlapi/boardgame/167792")
	if err != nil {
		log.Printf("error : %e ", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var bg documents.BoardgamesXML

	xmlErr := xml.Unmarshal(body, &bg)

	if xmlErr != nil {
		log.Printf("error %e", xmlErr)
	}
	log.Printf("%b", bg.XMLName)
	log.Printf("%b", bg.BoardgameXML.Maxplayers)
	log.Printf("%b", bg.BoardgameXML.Description)

	return "done"
}
