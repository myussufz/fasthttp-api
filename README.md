# fasthttp-api

fasthttp-api use fasthttp to sending request. In short, fasthttp server is up to 10 times faster than net/http. Below are benchmark results. [Benchmark](https://raw.githubusercontent.com/valyala/fasthttp/master/README.md)

This repo still under development. We accept any pull request. ^\_^

## Installation

```bash
  // dependency
  $ go get github.com/myussufz/fasthttp-api
  $ go get github.com/valyala/fasthttp
```

## Quick Start

### Simple Request Method

#### JSON Usage

```go
    var request struct {
        Name string `json:"name"`
    }

    request.Name = "test"

    var response struct {
        Name string `json:"name"`
    }

    if err = api.Fetch("http://google.com", &api.Option{
            Method: http.MethodPost,
            ContentType: api.ContentTypeJSON,
            Body: request,
        }).ToJSON(&response); err != nil {
        log.Println("error: ", err)
    }
```

#### XML Usage

```go
  var request struct {
        Name string `xml:"name"`
    }

    request.Name = "test"

    var response struct {
        Name string `xml:"name"`
    }

    if err = api.Fetch("http://google.com", &api.Option{
            Method: http.MethodPost,
            ContentType: api.ContentTypeXML,
            Body: request,
        }).ToXML(&response); err != nil {
        log.Println("error: ", err)
    }
```

#### Return String

```go
  var request struct {
        Name string `json:"name"`
    }

    request.Name = "test"

    data, err = api.Fetch("http://google.com", &api.Option{
        Method: http.MethodPost,
        ContentType: api.ContentTypeJSON,
        Body: request,
    }).ToString();
    if err != nil {
        log.Println("error: ", err)
    }

    log.Println("data: ", data)
```
