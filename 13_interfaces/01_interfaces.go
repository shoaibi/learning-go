package main

import "fmt"

// Always accept interfaces, specially for code that would be exposed to public

// single method interface it should be "function"-er
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

// implicitly implemented interface by adding the same method
func (cw *ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {
	c := &ConsoleWriter{}
	fmt.Println(c.Write([]byte("Hello World!")))

	// benefit? create an instance of type Writer that actually have ConsoleWriter
	// polymorphic behavior
	// right side can be anything, i shouldn't care because i just call the interface method
	var w Writer = &ConsoleWriter{}
	w.Write([]byte("Using main writer type instances"))
}
