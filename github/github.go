package github

import (
	"os"
	"time"

	"golang.org/x/oauth2"

	log "github.com/Sirupsen/logrus"
	gh "github.com/google/go-github/github"
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
)

const (
	// VENDOR namespace part
	VENDOR = "intel"
	// PLUGIN name namespace part
	NAME = "github"
	// VERSION of GitHub plugin
	VERSION = 1
	// TYPE is the plugin type
	TYPE = plugin.CollectorPluginType
)

var (
	personalAccessToken string
	org                 string
)

// RepositoryContentGetOptions represents an optional ref parameter
type RepositoryContentGetOptions struct {
	Ref string `url:"ref,omitempty"`
}

type githubCollector struct{}

// TokenSource is an encapsulation of the AccessToken string
type TokenSource struct {
	AccessToken string
}

func NewGithubCollector() *githubCollector {
	return &githubCollector{}
}

// Token authenticates via oauth
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func (g *githubCollector) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
}

func (g *githubCollector) CollectMetrics([]PluginMetricType) ([]PluginMetricType,
	error) {
	org = "intelsdi-x"
	personalAccessToken = os.Getenv("GITHUB_ACCESS_TOKEN")
	if len(personalAccessToken) == 0 {
		log.Fatal("Before you can use this you must set the GITHUB_ACCESS_TOKEN environment variable.")
	}
	// authentication has to happen here
	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := gh.NewClient(oauthClient) // authenticated to GitHub here

	opt := &gh.RepositoryListByOrgOptions{
		ListOptions: gh.ListOptions{PerPage: 10},
	}
	// get all pages of results
	var allRepos []gh.Repository

	// initialize the map of all Readmes
	readmeLibrary := make(map[string]string)
	for {
		repos, resp, err := client.Repositories.ListByOrg(org, opt)
		if err != nil {
			log.Fatal(err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}
}

func (g *githubCollector) GetMetricTypes(PluginConfigType) ([]PluginMetricType,
	error) {
	mts := []plugin.PluginMetricType{}
	mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"intel", "github", "foo"}})
	mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"intel", "github", "bar"}})
	return mts, nil
}

//Meta returns meta data for testing
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(
		NAME,
		VERSION,
		TYPE,
		[]string{plugin.SnapGOBContentType},
		[]string{plugin.SnapGOBContentType},
		plugin.CacheTTL(100*time.Millisecond),
		plugin.RoutingStrategy(plugin.StickyRouting),
	)
}
