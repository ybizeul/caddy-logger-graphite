package graphite_log

import (
	"fmt"
	"io"
	"strconv"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/marpaia/graphite-golang"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(GraphiteLog{})
}

/*
	GraphiteLog is a Caddy logger used to send serve activity to a Graphite database.

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

	.DirName  string
	.FileName string
*/
type GraphiteLog struct {
	// IP addess or ost name of the graphite server
	Server string `json:"server"`

	// Port number to be used (usually 2003)
	Port int `json:"port"`

	// Metrics Path, can be templated
	Path string `json:"path"`

	// Value to be sent, can be templated
	Value string `json:"value"`

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
	for d.Next() {
		if d.NextArg() {
			// Gaphite Server
			l.Server = d.Val()
		} else {
			return d.ArgErr()
		}
		if d.NextArg() {
			// Graphite Port
			p, err := strconv.Atoi(d.Val())
			if err != nil {
				l.logger.Error(err.Error())
				return err
			}
			l.Port = p
		} else {
			return d.ArgErr()
		}
		if d.NextArg() {
			// Graphite Path
			l.Path = d.Val()
		} else {
			return d.ArgErr()
		}
		if d.NextArg() {
			// Graphite Value
			l.Value = d.Val()
		} else {
			return d.ArgErr()
		}
		if d.NextArg() {
			// too many args
			return d.ArgErr()
		}
	}
	return nil
}

func (l *GraphiteLog) Validate() error {
	if l.Server == "" {
		return fmt.Errorf("No Server Set")
	}

	if l.Port == 0 {
		return fmt.Errorf("No Port Set")
	}

	if l.Path == "" {
		return fmt.Errorf("No Path Set")
	}

	if l.Value == "" {
		return fmt.Errorf("No Value Set")
	}

	return nil
}

func (g *GraphiteLog) String() string {
	return "graphite"
}

func (g *GraphiteLog) WriterKey() string {
	return "graphite_log"
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
