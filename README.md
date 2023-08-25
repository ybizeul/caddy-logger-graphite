# Caddy Log Graphite Exporter

This Caddy module allows you to write your logs to a Gaphite TSDB

You can customize the metric path based on path and file name of the request URI

## Install

1. Install `xcaddy`command

```
$ go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
```

You should have `xcaddy` in your go `bin/` folder, which should be `~/go/bin/`.
It's generally a good idea to have that in your `$PATH`

2. Build Caddy with the module

```
$ xcaddy build --with github.com/ybizeul/caddy-logger-graphite
```

## Usage

Your log settings must be set to `json`, a sample configuration is :

### Caddyfile
```
{
	auto_https off
}

http://localhost:8080 {
	file_server browse
	log graphite {
		format json
		output graphite {
			server 127.0.0.1
			port 2003
			path "downloads{{ .DirName }}.{{ .FileName }}.count"
			value "1"
			methods GET
		}
	}
}
```

### caddy.json
```
{
    "logging": {
        "logs": {
            "default": {
                "exclude": [
                    "http.log.access"
                ]
            },
            "log0": {
                "encoder": {
                    "format": "json"
                },
                "writer": {
                    "output": "graphite",
                    "server": "127.0.0.1",
                    "port": 2003,
                    "path": "downloads{{ .DirName }}.{{ .FileName }}.count",
                    "value": "1",
                    "methods": ["GET"]
                },
                "include": [
                    "http.log.access"
                ]
            }
        }
    },
    "apps": {
        "http": {
            "servers": {
                "srv0": {
                    "automatic_https": {
                        "disable": true
                    },
                    "listen": [
                        ":8080"
                    ],
                    "routes": [
                        {
                            "handle": [
                                {
                                    "handler": "file_server",
                                    "browse": {}
                                }
                            ]
                        }
                    ],
                    "logs": {}
                }
            }
        }
    }
}
```