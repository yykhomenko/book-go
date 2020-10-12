package github

const CommitsURLFormat = "https://api.github.com/repos/%s/%s/commits"

type Commit struct {
	Sha    string
	Author *User
}
