package documents

import "encoding/xml"

type BoardgamesXML struct {
	XMLName      xml.Name `xml:"boardgames"`
	Text         string   `xml:",chardata"`
	Termsofuse   string   `xml:"termsofuse,attr"`
	BoardgameXML struct {
		Text          string `xml:",chardata"`
		Objectid      string `xml:"objectid,attr"`
		Yearpublished string `xml:"yearpublished"`
		Minplayers    string `xml:"minplayers"`
		Maxplayers    string `xml:"maxplayers"`
		Playingtime   string `xml:"playingtime"`
		Minplaytime   string `xml:"minplaytime"`
		Maxplaytime   string `xml:"maxplaytime"`
		Age           string `xml:"age"`
		Name          []struct {
			Text      string `xml:",chardata"`
			Sortindex string `xml:"sortindex,attr"`
			Primary   string `xml:"primary,attr"`
		} `xml:"name"`
		Description struct {
			Text string   `xml:",chardata"`
			Br   []string `xml:"br"`
		} `xml:"description"`
		Thumbnail          string `xml:"thumbnail"`
		Image              string `xml:"image"`
		Boardgamepublisher []struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgamepublisher"`
		Boardgameaccessory []struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgameaccessory"`
		Boardgamepodcastepisode []struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgamepodcastepisode"`
		Boardgamehonor []struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgamehonor"`
		Boardgameversion []struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgameversion"`
		Boardgamemechanic []struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgamemechanic"`
		Boardgamefamily []struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgamefamily"`
		Boardgamecategory []struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgamecategory"`
		Boardgameexpansion []struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgameexpansion"`
		Boardgamedesigner struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgamedesigner"`
		Boardgameartist struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgameartist"`
		Boardgamegraphicdesigner struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgamegraphicdesigner"`
		Commerceweblink struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"commerceweblink"`
		Boardgamesubdomain struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgamesubdomain"`
		Videogamebg struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"videogamebg"`
		Boardgameimplementation []struct {
			Text     string `xml:",chardata"`
			Objectid string `xml:"objectid,attr"`
		} `xml:"boardgameimplementation"`
		Poll []struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"name,attr"`
			Title      string `xml:"title,attr"`
			Totalvotes string `xml:"totalvotes,attr"`
			Results    []struct {
				Text       string `xml:",chardata"`
				Numplayers string `xml:"numplayers,attr"`
				Result     []struct {
					Text     string `xml:",chardata"`
					Value    string `xml:"value,attr"`
					Numvotes string `xml:"numvotes,attr"`
					Level    string `xml:"level,attr"`
				} `xml:"result"`
			} `xml:"results"`
		} `xml:"poll"`
	} `xml:"boardgame"`
}
