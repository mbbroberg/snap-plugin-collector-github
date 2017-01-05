# UNDER DEVELPOMENT - Snap collector plugin - GitHub
A Go-based plugin for the [Snap telemetry framework](http://snap-telemetry.io) that collects some statistics from each repository for a specific user.

## Features

* collects "forks_count", "stargazers_count", "watchers_count", "open_issues_count", "rate_limit_hit"
* takes `username` and `password` as part of its policy, but does not require it.  

Notes: 

https://developer.github.com/v3/#rate-limiting

> For requests using Basic Authentication or OAuth, you can make up to 5,000 requests per hour. For unauthenticated requests, the rate limit allows you to make up to 60 requests per hour. Unauthenticated requests are associated with your IP address, and not the user making requests.


# Default README below - burn after reading.


## Plugin Naming, Files, and Directory
For your collector plugin, create a new repository and name your plugin project using the following format:

>snap-plugin-[plugin-type]-[plugin-name]

For example: 
>snap-plugin-collector-rand


Example files and directory structure:  
```
snap-plugin-[plugin-type]-[plugin-name]
 |--[plugin-name]
  |--[plugin-name].go  
  |--[plugin-name]_test.go  
 |--main.go
```

For example:
```
snap-plugin-collector-rand
 |--rand
  |--rand.go  
  |--rand_test.go  
 |--main.go
```

* The [plugin-name] folder (for example `rand`) will include all files to implement the appropriate interface methods
* Your [plugin-name] folder  will also include your test files.



## Interface Methods

In order to write a plugin for Snap, it is necessary to define a few methods to satisfy the appropriate interfaces. These interfaces must be defined for a collector plugin:


```go
// Plugin is an interface shared by all plugins and must implemented by each plugin to communicate with Snap.
type Plugin interface {
	GetConfigPolicy() (ConfigPolicy, error)
}

// Collector is a plugin which is the source of new data in the Snap pipeline.
type Collector interface {
	Plugin

	GetMetricTypes(Config) ([]Metric, error)
	CollectMetrics([]Metric) ([]Metric, error)
}
```
The interface is slightly different depending on what type (collector, processor, or publisher) of plugin is being written. Please see other plugin types for more details.



## Starting a plugin

After implementing a type that satisfies one of {collector, processor, publisher} interfaces, all that is left to do is to call the appropriate plugin.StartX() with your plugin specific meta options. For example with no meta options specified:

```go
	plugin.StartCollector(rand.RandCollector{}, pluginName, pluginVersion)
```

### Meta options

The available options are defined in [plugin/meta.go](https://github.com/intelsdi-x/snap-plugin-lib-go/tree/master/v/1/plugin/meta.go). You can use some or none of the options. The options with definitions/explanations are below:

```go
// ConcurrencyCount is the max number of concurrent calls the plugin
// should take.  For example:
// If there are 5 tasks using the plugin and its concurrency count is 2,
// snapd will keep 3 plugin instances running.
// ConcurrencyCount overwrites the default (5) for a Meta's ConcurrencyCount.
func ConcurrencyCount(cc int) MetaOpt {
}

// Exclusive == true results in a single instance of the plugin running
// regardless of the number of tasks using the plugin.
// Exclusive overwrites the default (false) for a Meta's Exclusive key.
func Exclusive(e bool) MetaOpt {
}

// Unsecure results in unencrypted communication with this plugin.
// Unsecure overwrites the default (false) for a Meta's Unsecure key.
func Unsecure(e bool) MetaOpt {
}

// RoutingStrategy will override the routing strategy this plugin requires.
// The default routing strategy is Least Recently Used.
// RoutingStrategy overwrites the default (LRU) for a Meta's RoutingStrategy.
func RoutingStrategy(r router) MetaOpt {
}

// CacheTTL will override the default cache TTL for the this plugin. snapd
// caches metrics on the daemon side for a default of 500ms.
// CacheTTL overwrites the default (500ms) for a Meta's CacheTTL.
func CacheTTL(t time.Duration) MetaOpt {
}
```

An example using some arbitrary values::

```go
        plugin.StartCollector(
                mypackage.Mytype{},
                pluginName,
                pluginVersion,
                plugin.ConcurrencyCount(2),
                plugin.Exclusive(true),
                plugin.Unsecure(true),
                plugin.RoutingStrategy(StickyRouter),
                plugin.CacheTTL(time.Second))				
```

## Testing
For testing reference the [Snap Testing Guidelines](https://github.com/intelsdi-x/snap/blob/master/CONTRIBUTING.md#testing-guidelines). To test your plugin with Snap you will need to have [Snap](https://github.com/intelsdi-x/snap) installed, check out these docs for [Snap setup details](https://github.com/intelsdi-x/snap/blob/master/docs/BUILD_AND_TEST.md#getting-started).

Each test file should specify the appropriate build tag such as "small", "medium", "large" (e.g. // +build small).


For example if you want to run only small tests:
```
// +build small
// you must include at least one line between the build tag and the package name.
package rand
```

For example if you want to run small and medium tests:
```
// +build small medium

package rand
```

## Ready to Share
You've made a plugin! Now it's time to share it. Create a release by following these [steps](https://help.github.com/articles/creating-releases/). We recommend that your release version match your plugin version, see example [here](https://github.com/intelsdi-x/snap-plugin-lib-go/blob/master/examples/collector/main.go#L29).

Don't forget to announce your plugin release on [slack](https://intelsdi-x.herokuapp.com/) and get your plugin added to the [Plugin Catalog](https://github.com/intelsdi-x/snap/blob/master/docs/PLUGIN_CATALOG.md)!



