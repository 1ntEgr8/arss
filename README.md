# arss
An rss reader

## Installation

TODO

## Usage

TODO

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
