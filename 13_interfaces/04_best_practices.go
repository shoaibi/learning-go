package main

func main() {
	// Prefer small interfaces, use composition to build large ones, io.Reader, io.Writer, empty interface
	// Don't export interfaces for types that will be consumed, export concrete type
	//        - Consumers can create there own interface against the methods being use
	// Export interface for own types being used in Packages
	// Design functions to use interfaces where possible, unless access to underlying data is required
	// Pass only the relevant interface. If a function needs only Writer, pass only Writer instead of ReaderWriter
}
