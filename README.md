# Go pub/sub

A minimal in-memory publish/subscribe broker in Go, built from goroutines, channels and a mutex.

- `subscribe(topic)` registers a new channel for a topic.
- `publish(topic, msg)` fans a message out to every subscriber of that topic.
- Publishers send their messages at random intervals, and a `sync.WaitGroup` keeps `main` alive until they finish.

Everything lives in a single file, `beesPublisher.go`.

## Run

```bash
go run beesPublisher.go
```

Each subscriber prints what it receives, so the interleaving changes from run to run. That variation is what the exercise is meant to show.
