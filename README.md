# arss
An rss reader

## Installation

TODO

## Usage

Run `arss`. You should see the rss reader automatically open in your default browser.

`arss` comes with a default web client. However, you can also write your own. See [Writing your own client]() for more information.

To tell `arss` to use your client instead of the default, use the `client-path` flag
```bash
arss --client-path /path/to/client
```

## Developing

The client is written in [Svelte](https://svelte.dev/) and the server is in [Go](https://golang.org/)

- Run the client
```bash
cd client && npm run dev
```
- Run the server
```bash
go run arss
```

You should see the client open up in the browser. 

The client supports hot-reload, however, the server doesn't. You will have to rebuild the server each time you make a change. By default, the server auto-opens the page in your browser. If you do not want this, pass in the `--headless` flag.

By default, the server uses port 8080. You can change this with the `--port` flag.

## Writing your own client

The `arss` server exposes a REST api that you can use to populate your client. Take a look at `main.go` for a list of routes.
