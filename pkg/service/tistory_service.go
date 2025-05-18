package service

import (
	"encoding/xml"
	"fmt"
	"github.com/RakunKo/Tistory-Readme-Go/internal/model"
	"net/http"
)

func GetTistoryRss(userName string) (model.Channel, error) {
	rssURL := fmt.Sprintf("https://%s.tistory.com/rss", userName)
	resp, err := http.Get(rssURL)
	if err != nil {
		return model.Channel{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return model.Channel{}, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	var rss model.Rss
	if err := xml.NewDecoder(resp.Body).Decode(&rss); err != nil {
		return model.Channel{}, err
	}
	return rss.Channel, nil
}
