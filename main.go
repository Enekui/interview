package main

import (
	// importing fmt package
	"fmt"
)

func main() {

	// defining new slice of strings
	// using short variables declaration
	names := []string{
		"Adianny",
		"Ramirez",
		"Armenteros",
	}

	// creating new empty stack
	myStack := Stack{}

	// printing the empty stack
	fmt.Printf("Here the stack is empty: %v\n", myStack)

	// iterating over the slice of names
	// this way we avoid to repeat code
	for _, n := range names {
		myStack.push(n)
	}

	// printing the stack
	fmt.Printf("Now we have added the names: %v\n", myStack.Name)

	// storing the returned value
	// in "r" variable
	// then printing everyting oput
	r := myStack.pop()
	fmt.Printf("Poped element: %s\n", r)
	fmt.Printf("We we have removed the last element from the stack: %v\n\n", myStack.Name)

	myQueue := Queue{}
	fmt.Printf("Here the queue is empty: %v\n", myQueue)

	// iterating over the slice of names
	// this way we avoid to repeat code
	for _, n := range names {
		myQueue.enq(n)
	}

	fmt.Printf("Now we have added the names: %v\n", myQueue.Name)

	// storing the returned value
	// in "r" variable
	// then printing everyting oput
	r = myQueue.deq()
	fmt.Printf("Dequeued element: %s\n", r)
	fmt.Printf("We we have removed the first element from the queue: %v\n", myQueue.Name)

}

// defining a new type Stack
// this tipe will contain an slice of strings
type Stack struct {
	Name []string
}

// using a pointer receiver of type Stack
// this way we create a push "method"
func (s *Stack) push(name string) {
	s.Name = append(s.Name, name)
}

// using a pointer receiver of type Stack
// this way we create a pop "method"
// here we return the removed element
func (s *Stack) pop() string {
	l := len(s.Name) - 1
	r := s.Name[l]
	s.Name = s.Name[:l]
	return r
}

// defining Queue type
// it will hold a slice of strings
type Queue struct {
	Name []string
}

// using a pointer receiver of type Queue
// This will link the func Enq to type Queue
// Enq = Enqueue
func (q *Queue) enq(name string) {
	q.Name = append(q.Name, name)
}

// using a pointer receiver of type Queue
// This will link the func Deq to type Queue
// we are again returning the removed element
// Deq = Dequeue
func (q *Queue) deq() string {
	r := q.Name[0]
	q.Name = q.Name[1:]
	return r
}
