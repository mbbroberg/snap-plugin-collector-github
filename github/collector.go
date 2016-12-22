import (
    "golang.org/x/oauth2"
)

var (
	personalAccessToken string
	org                 string
)

// RepositoryContentGetOptions represents an optional ref parameter
type RepositoryContentGetOptions struct {
	Ref string `url:"ref,omitempty"`
}

// tokenSource is an encapsulation of the AccessToken string
type tokenSource struct {
	AccessToken string
}

// NewGithubCollector is used to abstract how we dereference GithubCollector
func NewGithubCollector() *GithubCollector {
	return &GithubCollector{}
}