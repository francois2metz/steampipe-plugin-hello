package main

import (
	"github.com/judell/steampipe-plugin-hello/hello"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: hello.Plugin})
}
