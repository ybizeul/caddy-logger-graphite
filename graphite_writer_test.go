package graphite_log

import (
	"testing"

	mock "github.com/ybizeul/caddy-logger-graphite/mock"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
)

var completeLogLine = `{
  "level": "info",
  "ts": 1724065262.1466215,
  "logger": "http.log.access",
  "msg": "handled request",
  "request": {
    "remote_ip": "10.98.0.175",
    "remote_port": "50662",
    "client_ip": "10.98.0.175",
    "proto": "HTTP/1.1",
    "method": "GET",
    "host": "dl.example.org",
    "uri": "/file/file.txt",
    "headers": {
      "Accept-Language": [
        "ru,en-US;q=0.9,en;q=0.8"
      ],
      "Cookie": [],
      "Referer": [
        "https://dl.example.org/"
      ],
      "User-Agent": [
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"
      ],
      "Accept": [
        "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"
      ],
      "Priority": [
        "u=0, i"
      ],
      "Sec-Fetch-Site": [
        "same-site"
      ],
      "Sec-Fetch-User": [
        "?1"
      ],
      "X-Real-Ip": [
        "193.35.100.225"
      ],
      "Sec-Ch-Ua-Mobile": [
        "?0"
      ],
      "Sec-Ch-Ua-Platform": [
        "\"Windows\""
      ],
      "Sec-Fetch-Dest": [
        "document"
      ],
      "X-Forwarded-Port": [
        "443"
      ],
      "Sec-Ch-Ua": [
        "\"Not)A;Brand\";v=\"99\", \"Google Chrome\";v=\"127\", \"Chromium\";v=\"127\""
      ],
      "Sec-Fetch-Mode": [
        "navigate"
      ],
      "Upgrade-Insecure-Requests": [
        "1"
      ],
      "X-Forwarded-For": [
        "193.35.100.225"
      ],
      "X-Forwarded-Host": [
        "dl.example.org"
      ],
      "X-Forwarded-Proto": [
        "https"
      ],
      "X-Forwarded-Server": [
        "traefik-79d97d7c9c-lwf4p"
      ],
      "Accept-Encoding": [
        "gzip, deflate, br, zstd"
      ]
    }
  },
  "bytes_read": 0,
  "user_id": "",
  "duration": 0,
  "size": 793930340,
  "status": 200,
  "resp_headers": {
    "Strict-Transport-Security": [
      "max-age=31536000; includeSubDomains"
    ],
    "X-Amz-Id-2": [
      "dd9025bab4ad464b049177c95eb6ebf374d3b3fd1af9251148b658df7ac2e3e8"
    ],
    "Server": [
      "Caddy",
      "MinIO"
    ],
    "X-Amz-Request-Id": [
      "17ED1C31922C8500"
    ],
    "Date": [
      "Mon, 19 Aug 2024 11:00:31 GMT"
    ],
    "X-Ratelimit-Limit": [
      "1911"
    ],
    "Accept-Ranges": [
      "bytes"
    ],
    "Content-Type": [
      "application/gzip"
    ],
    "Vary": [
      "Origin",
      "Accept-Encoding"
    ],
    "X-Content-Type-Options": [
      "nosniff"
    ],
    "Last-Modified": [
      "Wed, 14 Aug 2024 20:12:17 GMT"
    ],
    "Etag": [
      "\"d85e0839c79a2d32430f07383dca2cbe-48\""
    ],
    "X-Xss-Protection": [
      "1; mode=block"
    ],
    "X-Ratelimit-Remaining": [
      "1908"
    ],
    "Content-Length": [
      "793930340"
    ]
  }
}`

var completeLogLinePost = `{
	"level": "info",
	"ts": 1724065262.1466215,
	"logger": "http.log.access",
	"msg": "handled request",
	"request": {
	  "remote_ip": "10.98.0.175",
	  "remote_port": "50662",
	  "client_ip": "10.98.0.175",
	  "proto": "HTTP/1.1",
	  "method": "POST",
	  "host": "dl.example.org",
	  "uri": "/file/file.txt",
	  "headers": {
		"Accept-Language": [
		  "ru,en-US;q=0.9,en;q=0.8"
		],
		"Cookie": [],
		"Referer": [
		  "https://dl.example.org/"
		],
		"User-Agent": [
		  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"
		],
		"Accept": [
		  "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"
		],
		"Priority": [
		  "u=0, i"
		],
		"Sec-Fetch-Site": [
		  "same-site"
		],
		"Sec-Fetch-User": [
		  "?1"
		],
		"X-Real-Ip": [
		  "193.35.100.225"
		],
		"Sec-Ch-Ua-Mobile": [
		  "?0"
		],
		"Sec-Ch-Ua-Platform": [
		  "\"Windows\""
		],
		"Sec-Fetch-Dest": [
		  "document"
		],
		"X-Forwarded-Port": [
		  "443"
		],
		"Sec-Ch-Ua": [
		  "\"Not)A;Brand\";v=\"99\", \"Google Chrome\";v=\"127\", \"Chromium\";v=\"127\""
		],
		"Sec-Fetch-Mode": [
		  "navigate"
		],
		"Upgrade-Insecure-Requests": [
		  "1"
		],
		"X-Forwarded-For": [
		  "193.35.100.225"
		],
		"X-Forwarded-Host": [
		  "dl.example.org"
		],
		"X-Forwarded-Proto": [
		  "https"
		],
		"X-Forwarded-Server": [
		  "traefik-79d97d7c9c-lwf4p"
		],
		"Accept-Encoding": [
		  "gzip, deflate, br, zstd"
		]
	  }
	},
	"bytes_read": 0,
	"user_id": "",
	"duration": 0,
	"size": 793930340,
	"status": 200,
	"resp_headers": {
	  "Strict-Transport-Security": [
		"max-age=31536000; includeSubDomains"
	  ],
	  "X-Amz-Id-2": [
		"dd9025bab4ad464b049177c95eb6ebf374d3b3fd1af9251148b658df7ac2e3e8"
	  ],
	  "Server": [
		"Caddy",
		"MinIO"
	  ],
	  "X-Amz-Request-Id": [
		"17ED1C31922C8500"
	  ],
	  "Date": [
		"Mon, 19 Aug 2024 11:00:31 GMT"
	  ],
	  "X-Ratelimit-Limit": [
		"1911"
	  ],
	  "Accept-Ranges": [
		"bytes"
	  ],
	  "Content-Type": [
		"application/gzip"
	  ],
	  "Vary": [
		"Origin",
		"Accept-Encoding"
	  ],
	  "X-Content-Type-Options": [
		"nosniff"
	  ],
	  "Last-Modified": [
		"Wed, 14 Aug 2024 20:12:17 GMT"
	  ],
	  "Etag": [
		"\"d85e0839c79a2d32430f07383dca2cbe-48\""
	  ],
	  "X-Xss-Protection": [
		"1; mode=block"
	  ],
	  "X-Ratelimit-Remaining": [
		"1908"
	  ],
	  "Content-Length": [
		"793930340"
	  ]
	}
  }`

var incompleteLogLine = `{
	"level": "info",
	"ts": 1724065262.1466215,
	"logger": "http.log.access",
	"msg": "handled request",
	"request": {
	  "remote_ip": "10.98.0.175",
	  "remote_port": "50662",
	  "client_ip": "10.98.0.175",
	  "proto": "HTTP/1.1",
	  "method": "GET",
	  "host": "dl.example.org",
	  "uri": "/file/file.txt",
	  "headers": {
		"Accept-Language": [
		  "ru,en-US;q=0.9,en;q=0.8"
		],
		"Cookie": [],
		"Referer": [
		  "https://dl.example.org/"
		],
		"User-Agent": [
		  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"
		],
		"Accept": [
		  "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"
		],
		"Priority": [
		  "u=0, i"
		],
		"Sec-Fetch-Site": [
		  "same-site"
		],
		"Sec-Fetch-User": [
		  "?1"
		],
		"X-Real-Ip": [
		  "193.35.100.225"
		],
		"Sec-Ch-Ua-Mobile": [
		  "?0"
		],
		"Sec-Ch-Ua-Platform": [
		  "\"Windows\""
		],
		"Sec-Fetch-Dest": [
		  "document"
		],
		"X-Forwarded-Port": [
		  "443"
		],
		"Sec-Ch-Ua": [
		  "\"Not)A;Brand\";v=\"99\", \"Google Chrome\";v=\"127\", \"Chromium\";v=\"127\""
		],
		"Sec-Fetch-Mode": [
		  "navigate"
		],
		"Upgrade-Insecure-Requests": [
		  "1"
		],
		"X-Forwarded-For": [
		  "193.35.100.225"
		],
		"X-Forwarded-Host": [
		  "dl.example.org"
		],
		"X-Forwarded-Proto": [
		  "https"
		],
		"X-Forwarded-Server": [
		  "traefik-79d97d7c9c-lwf4p"
		],
		"Accept-Encoding": [
		  "gzip, deflate, br, zstd"
		]
	  }
	},
	"bytes_read": 0,
	"user_id": "",
	"duration": 0,
	"size": 79393040,
	"status": 200,
	"resp_headers": {
	  "Strict-Transport-Security": [
		"max-age=31536000; includeSubDomains"
	  ],
	  "X-Amz-Id-2": [
		"dd9025bab4ad464b049177c95eb6ebf374d3b3fd1af9251148b658df7ac2e3e8"
	  ],
	  "Server": [
		"Caddy",
		"MinIO"
	  ],
	  "X-Amz-Request-Id": [
		"17ED1C31922C8500"
	  ],
	  "Date": [
		"Mon, 19 Aug 2024 11:00:31 GMT"
	  ],
	  "X-Ratelimit-Limit": [
		"1911"
	  ],
	  "Accept-Ranges": [
		"bytes"
	  ],
	  "Content-Type": [
		"application/gzip"
	  ],
	  "Vary": [
		"Origin",
		"Accept-Encoding"
	  ],
	  "X-Content-Type-Options": [
		"nosniff"
	  ],
	  "Last-Modified": [
		"Wed, 14 Aug 2024 20:12:17 GMT"
	  ],
	  "Etag": [
		"\"d85e0839c79a2d32430f07383dca2cbe-48\""
	  ],
	  "X-Xss-Protection": [
		"1; mode=block"
	  ],
	  "X-Ratelimit-Remaining": [
		"1908"
	  ],
	  "Content-Length": [
		"793930340"
	  ]
	}
  }`

func TestWriteLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	grmock := mock.NewMockGraphiteInterface(ctrl)
	g := GraphiteWriter{
		GraphiteLog: &GraphiteLog{
			Path:    "downloads.{{ .DirName }}.{{ .FileName }}.count",
			Value:   "1",
			Server:  "graphite.mon",
			Methods: []string{"GET"},
			logger:  zap.New(nil),
		},
		Graphite: grmock,
	}

	grmock.EXPECT().SimpleSend("downloads.file.file_txt.count", "1").Return(nil)

	i, err := g.Write([]byte(completeLogLine))
	if err != nil {
		t.Error(err)
	}
	if i != len(completeLogLine) {
		t.Error("Error in write")
	}

	i, err = g.Write([]byte(completeLogLinePost))
	if err != nil {
		t.Error(err)
	}
	if i != len(completeLogLinePost) {
		t.Error("Error in write")
	}

	i, err = g.Write([]byte(incompleteLogLine))
	if err != nil {
		t.Error(err)
	}
	if i != len(incompleteLogLine) {
		t.Error("Error in write")
	}
}
