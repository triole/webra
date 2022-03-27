# WebRA ![example workflow](https://github.com/triole/webra/actions/workflows/build.yaml/badge.svg)

<!--- mdtoc: toc begin -->

1. [Synopsis](#synopsis)
2. [Help](#help)<!--- mdtoc: toc end -->

## Synopsis

WebRA is a simple Web Request Assertion engine that does http requests and evaluates their answer.

## Help

```go mdox-exec="r -h"

simple web request assertion tool

Arguments:
  [<config>]    config toml file name, positional arg required

Flags:
  -h, --help                       Show context-sensitive help.
  -p, --print-vars                 print available vars
  -d, --debug                      debug mode
      --user-agent="WebRA/0.1."    user agent
  -j, --json-log                   enable json log, instead of text one
  -l, --log-file="/dev/stdout"     log file
  -v, --verbose                    verbose, also print positive test messages
  -V, --version-flag               display version
```
