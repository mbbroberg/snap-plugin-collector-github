package github

import "github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

const (
	// Name is the name of this plugin
	Name = "github"

	// Vendor is the company supporting this plugin
	Vendor = "evilcorp"

	// Version of GitHub plugin
	Version = 1
)

// make sure that we actually satisify required interface
var _ plugin.Collector = (*GithubCollector)(nil)

type GithubCollector struct {
}

var nsTypes = struct {
	githubMetricTypes []string
}{
	// names of available metrics
	githubMetricTypes: []string{"forks_count", "stargazers_count", "watchers_count", "open_issues_count", "rate_limit_hit"},
}

// CollectMetrics gets the metrics << make this better :)
func (GithubCollector) CollectMetrics(mst []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	metrics = append(metrics, plugin.Metric{
		Namespace: plugin.NewNamespace(Vendor, "foo"),
		Version:   1,
	})
	return metrics, nil
}

// GetMetricTypes gathers available measurements from this plugin
func (GithubCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	for _, metric := range nsTypes.githubMetricTypes {
		metric := plugin.Metric{Namespace: plugin.NewNamespace(Vendor, Name).AddStaticElement(metric)}
		metrics = append(metrics, metric)
	}
	return metrics, nil
}

// GetConfigPolicy must be implemented
func (GithubCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	policy.AddNewStringRule([]string{Vendor, Name}, "user", false)
	policy.AddNewStringRule([]string{Vendor, Name}, "password", false)
	return *policy, nil
}
