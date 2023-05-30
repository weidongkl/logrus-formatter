# logrus-formatter

The formatter of logrus, like python's logging.
```
[INFO] 2023-05-30 05:59:28,99  [xx.go:40]:  messags
```

## Installation

`go get gitee.com/weidongkl/logrus-formatter`

## Usage

```go
package main

import (
	formatter "gitee.com/weidongkl/logrus-formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetReportCaller(true)
	log.SetFormatter(&formatter.Formatter{})
	log.Infoln("test messages")
}

```
