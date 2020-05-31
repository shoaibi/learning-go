package main

func main() {
	// Do not create goroutines in libraries. Let consumers decide
	// When creating goroutine, know when it ends. Could have memory leaks if it is too old
	// Check for race condition on compile time: go run -race program.go
}
