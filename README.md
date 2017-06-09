![Themis Project](https://raw.githubusercontent.com/michaelkleinhenz/themis/master/logo.png)

This is a standalone backend for fabric8-planner that uses Mongo as a storage backend.

_Themis is the Titan goddess of divine law and order._

## Testing

Install ginko, then:

  $ ~/gopath/bin/ginkgo -r -cover

## Building

  $ go build -ldflags "-X 'main.ThemisBuildDate=$(date -u '+%Y-%m-%d %H:%M:%S')' -X main.ThemisVersion=$(git log --pretty=format:'%h' -n 1)"

  