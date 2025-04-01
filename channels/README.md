## Channel

Channel is a pipe between goroutines.
There are 2 types of channel

- unbuffered channel:

```go
ch := make(chan int)
ch <- 1
```

- buffered channel

```go
ch := make(chan int, 10) # capacity is 10
ch <- 1
```
