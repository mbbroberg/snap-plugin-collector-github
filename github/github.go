package github

import (
	"os"
	"time"

	"golang.org/x/oauth2"

	gh "github.com/google/go-github/github"
	log "github.com/intelsdi-x/snap-plugin-utilities/logger"
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
)

var (
	personalAccessToken string
	org                 string
)

// RepositoryContentGetOptions represents an optional ref parameter
type RepositoryContentGetOptions struct {
	Ref string `url:"ref,omitempty"`
}

type GithubCollector struct{}

// tokenSource is an encapsulation of the AccessToken string
type tokenSource struct {
	AccessToken string
}

// NewGithubCollector is used to abstract how we dereference GithubCollector
func NewGithubCollector() *GithubCollector {
	return &GithubCollector{}
}

// Token authenticates via oauth
func (t *tokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func (g *GithubCollector) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	return c, nil
}

func (g *GithubCollector) CollectMetrics([]plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	org = "intelsdi-x"
	personalAccessToken = os.Getenv("GITHUB_ACCESS_TOKEN")
	if len(personalAccessToken) == 0 {
		log.LogFatal("Before you can use this you must set the GITHUB_ACCESS_TOKEN environment variable.")
	}
	// authentication has to happen here
	tokenSource := &tokenSource{
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
	// readmeLibrary := make(map[string]string)
	for {
		repos, resp, err := client.Repositories.ListByOrg(org, opt)
		if err != nil {
			log.LogFatal("Error on GitHub Request. Error & Response: ", err, resp)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}
	for _, rp := range allRepos {
		log.LogInfo("Found a repo: ", *rp.Name)
	}

	log.LogInfo("")
	// return something?
	mts := []plugin.PluginMetricType{}
	mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"intel", "github", "foo"}})
	mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"intel", "github", "bar"}})
	return mts, nil
}

func (g *GithubCollector) GetMetricTypes(plugin.PluginMetricType) ([]plugin.PluginMetricType,
	error) {
	mts := []plugin.PluginMetricType{}
	mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"intel", "github", "foo"}})
	mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"intel", "github", "bar"}})
	return mts, nil
}

//Meta returns meta data for testing
func Meta() *plugin.PluginMeta {
	pluginType := plugin.CollectorPluginType
	return plugin.NewPluginMeta(
		NAME,
		VERSION,
		pluginType,
		[]string{plugin.SnapGOBContentType},
		[]string{plugin.SnapGOBContentType},
		plugin.CacheTTL(100*time.Millisecond),
		plugin.RoutingStrategy(plugin.StickyRouting),
	)
}
