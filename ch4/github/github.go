package github

const APIURL = "https://api.github.com"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
