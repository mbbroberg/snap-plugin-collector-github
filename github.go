package main

import "github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

const (
	// NAMESPACE part. Change this when you package up your own.
	NAMESPACE = "intel"
	// NAME is the name of this plugin
	NAME = "github"
	// VERSION of GitHub plugin
	VERSION = 1
)

// make sure that we actually satisify required interface
var _ plugin.CollectorPlugin = (*GithubCollector)(nil)

type GithubCollector struct {
}

// CollectMetrics gets the metrics << make this better :)
func (g *GithubCollector) CollectMetrics(mst []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	metrics = append(metrics, plugin.Metric{
		Namespace: plugin.NewNamespace(NAMESPACE, "foo"),
		Version:   1,
	})
	return metrics, nil
}

// GetMetricTypes gathers available measurements from this plugin
func (GithubCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}
	metrics = append(metrics, plugin.Metric{
		Namespace: plugin.NewNamespace(NAMESPACE, "foo"),
		Version:   1,
	})
	return metrics, nil
}

// GetConfigPolicy must be implemented
func (GithubCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()
	return *policy, nil
}
