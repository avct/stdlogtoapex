# stdlogtoapex
A small library to bridge Go's log package to [github.com/apex/log](https://github.com/apex/log).

# Redirecting standard log messages to apex log.

`stdlogtoapex` provides `Writer` type that implements the `io.Writer` interface and can be passed to the standard library `log.SetOutput` function.  When this is done, all subsequent calls to the `log` packages output functions will send their output to via this `Writer`, and it will, in turn, log that output via Apex log's default handler. 

## Example

```go
package main

import (
    "log"
    ]
    "github.com/avct/stdlogtoapex"
    alog "github.com/apex/log"
    "github.com/apex/log/handlers/cli"
)

func main() {
    handler := cli.Default
	alog.SetHandler(handler)
	writer := stdlogtoapex.NewWriter()
    
    // From this point all output from the standard library log package will be ouptup via the Writer
	log.SetOutput(writer)
	log.Print("Hello!")
}
```

# Caveat

This package will not capture logging sent via the "log/syslog" package.
