package graphite_log

import (
	"bytes"
	"encoding/json"
	"html/template"
	"path"
	"slices"
	"strings"

	"github.com/marpaia/graphite-golang"
	"go.uber.org/zap"
)

type LogLine struct {
	Level       string              `json:"level"`
	Date        float64             `json:"ts"`
	Logger      string              `json:"logger"`
	Msg         string              `json:"msg"`
	Request     Request             `json:"request"`
	BytesRead   int64               `json:"bytes_read"`
	UserID      string              `json:"user_id"`
	Duration    float64             `json:"duration"`
	Size        int64               `json:"size"`
	Status      int                 `json:"status"`
	RespHeaders map[string][]string `json:"resp_headers"`

	DirName  string `json:"DirName"`
	FileName string `json:"FileName"`
}
type Request struct {
	RemoteIP   string              `json:"remote_ip"`
	RemotePort string              `json:"remote_port"`
	ClientIP   string              `json:"client_ip"`
	Proto      string              `json:"proto"`
	Method     string              `json:"method"`
	Host       string              `json:"host"`
	URI        string              `json:"uri"`
	Headers    map[string][]string `json:"headers"`
}

type GraphiteWriter struct {
	GraphiteLog *GraphiteLog
	Graphite    *graphite.Graphite
}

func (g *GraphiteWriter) Write(p []byte) (n int, err error) {
	// g.GraphiteLog.logger.Info(string(p))
	j := LogLine{}
	err = json.Unmarshal(p, &j)
	if err != nil {
		g.GraphiteLog.logger.Error(err.Error())
	}
	if j.Status == 200 {
		if len(g.GraphiteLog.Methods) > 0 {
			if !slices.Contains(g.GraphiteLog.Methods, j.Request.Method) {
				return len(p), nil
			}
		}
		sanitized := strings.Replace(j.Request.URI, ".", "_", -1)[1:]
		j.DirName = strings.Replace(path.Dir(sanitized), "/", ".", -1)[1:]
		j.FileName = strings.Replace(path.Base(sanitized), ".", "_", -1)

		pathTemplate, err := template.New("path").Parse(g.GraphiteLog.Path)
		if err != nil {
			g.GraphiteLog.logger.Error(err.Error())
		}
		valueTemplate, err := template.New("path").Parse(g.GraphiteLog.Value)
		if err != nil {
			g.GraphiteLog.logger.Error(err.Error())
		}

		var r bytes.Buffer
		err = pathTemplate.Execute(&r, j)
		if err != nil {
			g.GraphiteLog.logger.Error(err.Error())
		}
		path := r.String()

		r.Reset()
		err = valueTemplate.Execute(&r, j)
		if err != nil {
			g.GraphiteLog.logger.Error(err.Error())
		}
		value := r.String()

		g.GraphiteLog.logger.Info("Writing value to carbon", zap.String("path", path), zap.String("value", value))

		err = g.Graphite.SimpleSend(path, value)
		if err != nil {
			g.GraphiteLog.logger.Error(err.Error())
		}
	}
	return len(p), nil
}

func (g *GraphiteWriter) Close() error {
	return nil
}
