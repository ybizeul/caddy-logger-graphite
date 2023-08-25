package graphite_log

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"golang.org/x/exp/slices"
)

func TestUnmarshalCaddyfile1(t *testing.T) {
	var token string
	var g GraphiteLog
	var ctx caddy.Context

	// Test 1
	token = `graphite {
		server 127.0.0.1
		port 2003
		path "downloads{{ .Dirname }}.{{ .Filename }}.count"
		value 1
		methods GET HEAD
		}`

	ctx = caddy.Context{}
	g = GraphiteLog{}
	g.Provision(ctx)

	err := g.UnmarshalCaddyfile(caddyfile.NewTestDispenser(token))
	if err != nil {
		t.Error(err)
	}

	g.Validate()
	if err != nil {
		t.Error(err)
	}

	if !(g.Server == "127.0.0.1" &&
		g.Port == 2003 &&
		g.Path == "downloads{{ .Dirname }}.{{ .Filename }}.count" &&
		g.Value == "1" &&
		slices.Contains(g.Methods, "GET") &&
		slices.Contains(g.Methods, "HEAD")) {
		t.Error(fmt.Errorf("Unexpected error in arguments"))
	}
}
func TestUnmarshalCaddyfile2(t *testing.T) {
	ctx := caddy.Context{}
	g := GraphiteLog{}

	g.Provision(ctx)

	// Test 2
	token := `graphite {
		server 127.0.0.1
		path "downloads{{ .Dirname }}.{{ .Filename }}.count"
		}`

	err := g.UnmarshalCaddyfile(caddyfile.NewTestDispenser(token))

	if err != nil {
		t.Error(err)
	}

	g.Validate()
	if err != nil {
		t.Error(err)
	}

	if !(g.Server == "127.0.0.1" &&
		g.Port == 2003 &&
		g.Path == "downloads{{ .Dirname }}.{{ .Filename }}.count" &&
		g.Value == "1") {
		t.Error(fmt.Errorf("Unexpected error in arguments"))
	}
}
func TestUnmarshalCaddyfile3(t *testing.T) {
	ctx := caddy.Context{}
	g := GraphiteLog{}

	g.Provision(ctx)

	// Test 3
	token := `graphite {
		server 127.0.0.1
		port test
		path "downloads{{ .Dirname }}.{{ .Filename }}.count"
		value 1
		methods [ GET HEAD ]`

	g.Provision(ctx)

	err := g.UnmarshalCaddyfile(caddyfile.NewTestDispenser(token))

	if err == nil {
		t.Error(fmt.Errorf("Invalid port number should have been caught"))
	}
}
