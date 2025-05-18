package model

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title          string `xml:"title"`
	ManagingEditor string `xml:"managingEditor"`
	Image          Image  `xml:"image"`
}

type Image struct {
	URL string `xml:"url"`
}
