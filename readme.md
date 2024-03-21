# Web Request Assert ![example workflow](https://github.com/triole/webra/actions/workflows/build.yaml/badge.svg)

<!-- toc -->

- [Synopsis](#synopsis)
- [Help](#help)

<!-- /toc -->

## Synopsis

Web Request Assert is a simple web request assertion engine that does http requests and evaluates their answer. Configuration is provided in a toml file of which examples can be found inside the `conf` folder. Requests are fired asynchronously to provide a maximum of speed. Basic authentication is supported as well.

## Help

```go mdox-exec="r -h"

simple web request assertion tool

Arguments:
  [<config>]    config toml file name, positional arg required

Flags:
  -h, --help                       Show context-sensitive help.
  -u, --user-agent="WebRA/0.1."    user agent
  -n, --threads=512                max threads, default no of avail. cpu threads
                                   times 32
  -t, --timeout=5                  request timeout in seconds
  -x, --export=STRING              export full test data into json file
  -l, --log-file="stdout"          log file
  -e, --log-level="info"           log level
      --log-no-colors              disable output colours, print plain text
      --log-json                   enable json log, instead of text one
  -d, --debug                      debug mode
  -V, --version-flag               display version
```
