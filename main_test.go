package graphite_log

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

func TestUnmarshalCaddyfile(t *testing.T) {
	var token string
	var g GraphiteLog
	var ctx caddy.Context

	token = `graphite 127.0.0.1 2003 "downloads{{ .Dirname }}.{{ .Filename }}.count" 1`

	g = GraphiteLog{}
	ctx = caddy.Context{}
	var err error

	g.Provision(ctx)

	err = g.UnmarshalCaddyfile(caddyfile.NewTestDispenser(token))

	if err != nil {
		t.Error(err)
	}

	if !(g.Server == "127.0.0.1" &&
		g.Port == 2003 &&
		g.Path == "downloads{{ .Dirname }}.{{ .Filename }}.count") {
		t.Error(fmt.Errorf("Unexpected error in arguments"))
	}

	token = `graphite 127.0.0.1 test "downloads{{ .Dirname }}.{{ .Filename }}.count" 1`

	g = GraphiteLog{}
	ctx = caddy.Context{}

	g.Provision(ctx)

	err = g.UnmarshalCaddyfile(caddyfile.NewTestDispenser(token))

	if err == nil {
		t.Error(fmt.Errorf("Invalid port number should have been caught"))
	}
}
