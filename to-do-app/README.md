## Example for realtime app with Golang and Rethinkdb

- Start your db first, by running
```
  rethinkdb
```

- Clone this project, then
```
  go get
  go run *.go
```

- Open http://localhost:3000, for testing websocket

- App Routes:

```
  - /items GET
  - /items POST body<{ text: string }>
  - /items/${id} GET
  - /items/${id} PUT body<{ text: string }>
  - /items/${id} DELETE
```
