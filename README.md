# arss
An rss reader that doesn't try to do too much

## Motivation

I wanted a very minimal RSS reader that **wasn't** feature packed, but could be extended. Something super simple.

## Installation

### From source

- Clone the repo
```
git clone https://github.com/1ntEgr8/arss
```

- Build the client and server
```
npm --prefix client run build && go build arss
```

- Run the program
```
./arss
```

## Usage

Run `arss`. You should see the rss reader automatically open in your default browser.

`arss` comes with a default web client. However, you can also write your own. See [Writing your own client](#writing-your-own-client) for more information.

To tell `arss` to use your client instead of the default, use the `client-path` flag
```bash
arss --client-path /path/to/client
```

## Developing

The client is written in [Svelte](https://svelte.dev/). The server is in [Go](https://golang.org/)

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

The `arss` server exposes a REST api that you can use to populate your client. Take a look at [main.go](./main.go) for a list of routes.

## Bugs/New features

Feel free to [file an issue](https://github.com/1ntEgr8/arss/issues/new) if you spot a bug or want to make a feature request

## Contributing

PRs are welcome :) See [Developing](#Developing) for details on how to set up a dev environment.
