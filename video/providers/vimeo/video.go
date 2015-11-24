package vimeo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Video struct {
	AuthorName      string `json:"author_name"`
	AuthorURL       string `json:"author_url"`
	Description     string `json:"description"`
	Duration        int    `json:"duration"`
	Height          int    `json:"height"`
	Html            string `json:"html"`
	IsPlus          string `json:"is_plus"`
	ProviderName    string `json:"provider_name"`
	ProviderURL     string `json:"provider_url"`
	ThumbnailHeight int    `json:"thumbnail_height"`
	ThumbnailURL    string `json:"thumbnail_url"`
	ThumbnailWidth  int    `json:"thumbnail_width"`
	Title           string `json:"title"`
	Type            string `json:"type"`
	Uri             string `json:"uri"`
	Version         string `json:"version"`
	VideoID         int    `json:"video_id"`
	Width           int    `json:"width"`
	PlayerURL       string `json:"-"`
}

func (v Video) ID() string {
	return strconv.Itoa(v.VideoID)
}

func (v Video) Player() string {
	return v.PlayerURL
}

func (v Video) URI() string {
	return v.Uri
}

func (v *Video) decode(res *http.Response) error {
	return json.NewDecoder(res.Body).Decode(v)
}

func New(src string) (Video, error) {
	var res *http.Response
	var req *http.Request
	var vid Video
	var err error

	location := fmt.Sprintf("https://vimeo.com/api/oembed.json?url=%s", url.QueryEscape(src))

	client := &http.Client{}

	req, err = http.NewRequest("GET", location, nil)

	if err != nil {
		return vid, err
	}

	res, err = client.Do(req)

	if err != nil {
		return vid, err
	}

	defer res.Body.Close()

	if res.StatusCode > 300 {
		return vid, errors.New(res.Status)
	}

	if err := vid.decode(res); err != nil {
		return vid, err
	}

	vid.PlayerURL = "//player.vimeo.com/video/" + vid.ID()

	return vid, nil
}
