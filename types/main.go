package main

import "fmt"

// Definition: This creates an alias for the string type. A can be used interchangeably with string throughout your code.
// Characteristics:
// No new type is created; A is just another name for string.
// You cannot define methods on A since it is merely an alias for an existing type.
// Assignments between A and string are direct, with no need for type conversion.
type A = string

// Definition: This creates a new type B that has the same underlying representation as string, but it is distinct from string.
// Characteristics:
// B is a completely new type, separate from string, even though it has the same underlying structure.
// You can define methods on B, allowing you to extend its functionality.
// Assignments between B and string require explicit type conversion.
type B string

func (b *B) foo() {
	fmt.Println("foo")
}

func main() {
	var a A = "hello" // a is of type string
	var b B = "world" // b is of type B

	fmt.Printf("a is %T\n", a) // Output: a is string
	fmt.Printf("b is %T\n", b) // Output: b is main.B

	// Assigning between types
	// var c B = a // This would cause a compile-time error
	var c B = B(a) // This works: explicit conversion from A to B
	fmt.Println(c) // Output: hello
}
