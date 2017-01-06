package github

// "github.com/google/go-github/github"

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

func (g *GithubCollector) authenticateUser(user string, password string) error {
	return nil
}

// Need to authenticate as a user if I have user information.
// Then query for that information unless
