package youtube

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Video struct {
	ApiVersion string `json:"apiVersion"`
	Data       struct {
		AccessControl struct {
			AutoPlay     string `json:"autoPlay"`
			Comment      string `json:"comment"`
			CommentVote  string `json:"commentVote"`
			Embed        string `json:"embed"`
			List         string `json:"list"`
			Rate         string `json:"rate"`
			Syndicate    string `json:"syndicate"`
			VideoRespond string `json:"videoRespond"`
		} `json:"accessControl"`
		AspectRatio   string `json:"aspectRatio"`
		Category      string `json:"category"`
		CommentCount  int    `json:"commentCount"`
		Description   string `json:"description"`
		Duration      int    `json:"duration"`
		FavoriteCount int    `json:"favoriteCount"`
		ID            string `json:"id"`
		LikeCount     string `json:"likeCount"`
		Player        struct {
			Default string `json:"default"`
			Mobile  string `json:"mobile"`
		} `json:"player"`
		Rating      float64 `json:"rating"`
		RatingCount int     `json:"ratingCount"`
		Thumbnail   struct {
			HqDefault string `json:"hqDefault"`
			SqDefault string `json:"sqDefault"`
		} `json:"thumbnail"`
		Title     string `json:"title"`
		Updated   string `json:"updated"`
		Uploaded  string `json:"uploaded"`
		Uploader  string `json:"uploader"`
		ViewCount int    `json:"viewCount"`
	} `json:"data"`
	PlayerURL string `json:"-"`
}

func (v Video) ID() string {
	return v.Data.ID
}

func (v Video) Player() string {
	return v.PlayerURL
}

func (v Video) URI() string {
	return v.Data.ID
}

func (v *Video) decode(res *http.Response) error {
	return json.NewDecoder(res.Body).Decode(v)
}

func New(src string) (Video, error) {
	var res *http.Response
	var req *http.Request
	var vid Video
	var err error

	id := extractID(src)

	location := fmt.Sprintf("http://gdata.youtube.com/feeds/api/videos/%s?v=2&alt=jsonc", id)

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

	vid.PlayerURL = "//www.youtube.com/embed/" + vid.ID()

	return vid, nil
}

// to possibilities: https://www.youtube.com/watch?v=YLKzvHrU8OU or https://youtu.be/YLKzvHrU8OU
func extractID(srcString string) string {
	src, err := url.Parse(srcString)

	if err != nil {
		return ""
	}

	// try query string
	id := src.Query().Get("v")

	if id == "" {
		// try uri part
		parts := strings.Split(src.Path, "/")
		id = parts[len(parts)-1] // id should be the last part of the path
	}

	return id
}
