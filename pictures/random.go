package pictures

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"
)

type Picture struct {
	ID             string      `json:"id"`
	CreatedAt      string      `json:"created_at"`
	UpdatedAt      string      `json:"updated_at"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Color          string      `json:"color"`
	Description    interface{} `json:"description"`
	AltDescription string      `json:"alt_description"`
	Urls           struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	Links struct {
		Self             string `json:"self"`
		HTML             string `json:"html"`
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
	} `json:"links"`
	Categories             []interface{} `json:"categories"`
	Sponsored              bool          `json:"sponsored"`
	SponsoredBy            interface{}   `json:"sponsored_by"`
	SponsoredImpressionsID interface{}   `json:"sponsored_impressions_id"`
	Likes                  int           `json:"likes"`
	LikedByUser            bool          `json:"liked_by_user"`
	CurrentUserCollections []interface{} `json:"current_user_collections"`
	User                   struct {
		ID              string `json:"id"`
		UpdatedAt       string `json:"updated_at"`
		Username        string `json:"username"`
		Name            string `json:"name"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		TwitterUsername string `json:"twitter_username"`
		PortfolioURL    string `json:"portfolio_url"`
		Bio             string `json:"bio"`
		Location        string `json:"location"`
		Links           struct {
			Self      string `json:"self"`
			HTML      string `json:"html"`
			Photos    string `json:"photos"`
			Likes     string `json:"likes"`
			Portfolio string `json:"portfolio"`
			Following string `json:"following"`
			Followers string `json:"followers"`
		} `json:"links"`
		ProfileImage struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"profile_image"`
		InstagramUsername string `json:"instagram_username"`
		TotalCollections  int    `json:"total_collections"`
		TotalLikes        int    `json:"total_likes"`
		TotalPhotos       int    `json:"total_photos"`
		AcceptedTos       bool   `json:"accepted_tos"`
	} `json:"user"`
	Exif struct {
		Make         string `json:"make"`
		Model        string `json:"model"`
		ExposureTime string `json:"exposure_time"`
		Aperture     string `json:"aperture"`
		FocalLength  string `json:"focal_length"`
		Iso          int    `json:"iso"`
	} `json:"exif"`
	Views     int `json:"views"`
	Downloads int `json:"downloads"`
}

func GetPicture() Picture {
	//Get picture
	url := fmt.Sprintf("https://api.unsplash.com/photos/random?client_id=%s", os.Getenv("UNSPLASH_ACCESS_KEY"))
	fmt.Println(url)
	resp, _ := http.Get(url)
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	var pic Picture
	buf, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buf, &pic)
	fmt.Printf("%#v\n", pic)
	return pic
}

func DownloadPicture(pic Picture) (m image.Image, err error) {
	resp, _ := http.Get(pic.Urls.Small)
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	m, _, err = image.Decode(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return m, err
}
