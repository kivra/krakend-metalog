# Metalog Middleware

The `metalog` middleware can be used for adding metadata to request logs. The
package includes the `metalog` middleware as well as a custom log formatter.

## Installation

To install `metalog` from GitHub:

    go get -u github.com/kivra/krakend-metalog@<commit hash>

Then add `metalog` to the KrakenD [`handler_factory`](https://github.com/devopsfaith/krakend-ce/blob/master/handler_factory.go)
chain:

```go
handlerFactory = metalog.HandlerFactory(handlerFactory)
```

Finally, add the `metalog` log formatter to the `luragin.EngineOptions` when
creating KrakenD's [`router_engine`](https://github.com/devopsfaith/krakend-ce/blob/master/router_engine.go):

```go
opt.Formatter = metalog.LogFormatter
```

## Usage

The `metalog` middleware can be added to an endpoint's `extra_config` and allows
to define metadata that is logged together with the request log for each handled
request. For example, the following configuration adds `owner`, `group` and
`tags` data to the request log.

```json
"endpoints": [
  {
    "endpoint": "/test",
    "extra_config": {
      "kivra/metalog": {
        "owner": "My Team",
        "group": "Test",
        "tags": ["some", "tags"]
      }
    },
    "backend": [ "..." ]
  }
]
```
