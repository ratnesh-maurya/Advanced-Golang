Channels must be  inside the gocoroutines package. They are used to communicate between goroutines. Channels are a typed conduit through which you can send and receive values with the channel operator, <-. The data flows in the direction of the arrow. A channel is created with the make function, and the type of data that can be sent and received is specified in the channel declaration. 


this can be prevented by buffering the channel. A buffered channel has a capacity, which is the number of values that can be stored in the channel. If the capacity is greater than zero, the channel is buffered; if the capacity is zero or absent, the channel is unbuffered. Here is an example of a buffered channel for strings with a capacity of 2:

```go
ch := make(chan string, 2)
```

when using channel in loops make sure to close the channel when you are done with it. This is done by calling the close function on the channel. Closing a channel indicates that no more values will be sent on it. This can be useful to communicate completion to the channel’s receivers. Here is an example of closing a channel when all the values have been sent:

```go
func main() {
    ch := make(chan string, 2)
    ch <- "hello"
    ch <- "world"
    close(ch)
}
``` 