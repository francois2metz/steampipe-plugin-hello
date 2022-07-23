package hello

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"

)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-hello",
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"hello": tableHello(ctx),
		},
	}
	return p
}
