package common

import (
	"flag"
)

var (
	Port   = flag.Int("port", 3000, "the listening port")
	LogDir = flag.String("log-dir", "", "specify the log directory")
)
