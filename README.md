# Web IM

[中文文档](./README_ZH.md)

This sample is about using long polling and WebSocket to build a web-based chat room based on beego.

- [Documentation](http://beego.me/docs/examples/chat.md)

## Installation

```bash
cd $GOPATH/src/samples/WebIm
go get github.com/gorilla/websocket
go get github.com/beego/i18n
bee run
```

## Usage

enter chat room from

```http
http://127.0.0.1:8080 
```

## Deployment

See the guide at <https://render.com/docs/deploy-beego>.
