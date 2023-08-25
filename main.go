package graphite_log

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/marpaia/graphite-golang"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(GraphiteLog{})
}

/*
GraphiteLog is a Caddy logger used to send server activity to a Graphite
database.

Templating is available as follow :

	.Level
	.Date
	.Logger
	.Msg
	.Request
		.RemoteIP
		.RemotePort
		.ClientIP
		.Proto
		.Method
		.Host
		.URI
		.Headers
	.BytesRead
	.UserID
	.Duration
	.Size
	.Status
	.RespHeaders map[string][]string

	.DirName
	.FileName
*/
type GraphiteLog struct {
	// IP address or host name of the graphite server
	Server string `json:"server"`

	// Port number to be used (usually 2003)
	Port int `json:"port"`

	// Metrics Path, can be templated
	Path string `json:"path"`

	// Value to be sent, can be templated
	Value string `json:"value"`

	// Methods to be logged
	Methods []string `json:"methods"`

	logger *zap.Logger
}

func (GraphiteLog) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "caddy.logging.writers.graphite",
		New: func() caddy.Module { return new(GraphiteLog) },
	}
}

func (l *GraphiteLog) Provision(ctx caddy.Context) error {
	l.logger = ctx.Logger() // g.logger is a *zap.Logger
	return nil
}

func (l *GraphiteLog) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	if !d.NextArg() {
		return d.ArgErr()
	}
	for block := d.Nesting(); d.NextBlock(block); {
		switch d.Val() {
		case "server":
			if !d.NextArg() {
				return d.ArgErr()
			}
			l.Server = d.Val()

		case "port":
			if !d.NextArg() {
				return d.ArgErr()
			}
			p, err := strconv.Atoi(d.Val())
			if err != nil {
				l.logger.Error(err.Error())
				return d.ArgErr()
			}
			l.Port = p
		case "path":
			if !d.NextArg() {
				return d.ArgErr()
			}
			l.Path = d.Val()

		case "value":
			if !d.NextArg() {
				return d.ArgErr()
			}
			l.Value = d.Val()

		case "methods":
			for d.NextArg() {
				l.Methods = append(l.Methods, d.Val())
			}
		}
	}
	return nil
}

func (l *GraphiteLog) Validate() error {
	if l.Server == "" {
		return fmt.Errorf("No Server Set")
	}

	if l.Port == 0 {
		l.Port = 2003
	}

	if l.Path == "" {
		return fmt.Errorf("No Path Set")
	}

	if l.Value == "" {
		l.Value = "1"
	}

	return nil
}

func (g *GraphiteLog) String() string {
	return "graphite"
}

func (g *GraphiteLog) WriterKey() string {
	return fmt.Sprintf("graphite_log_%s_%d_%s_%s_%s", g.Server, g.Port, g.Path, g.Value, strings.Join(g.Methods, ","))
}

func (l *GraphiteLog) OpenWriter() (io.WriteCloser, error) {
	// Open connection to Graphite server
	graphite, err := graphite.NewGraphite(l.Server, l.Port)
	if err != nil {
		l.logger.Error(err.Error())
	}

	return &GraphiteWriter{
		GraphiteLog: l,
		Graphite:    graphite,
	}, nil
}

// Interface guards
var (
	_ caddy.Provisioner     = (*GraphiteLog)(nil)
	_ caddy.WriterOpener    = (*GraphiteLog)(nil)
	_ caddyfile.Unmarshaler = (*GraphiteLog)(nil)
)
