package main

import (
	"fmt"
	"go_data_structures_algorithms/stacks_and_queues"
)

func main() {
	// stll := stacks_and_queues.NewStackLinkedList()
	// stll.Push(1)
	// stll.Push(2)
	// stll.Push(3)
	// stll.Push(4)
	// stll.Print()
	// stll.Push(5)
	// stll.Print()
	// fmt.Println(stll.Peek())
	// fmt.Println(stll.Pop())
	// fmt.Println(stll.Pop())
	// fmt.Println(stll.Pop())
	// stll.Print()

	// starr := stacks_and_queues.NewStackArray()
	// starr.Push(1)
	// starr.Push(2)
	// starr.Push(3)
	// starr.Push(4)
	// starr.Print()
	// starr.Push(5)
	// starr.Print()
	// fmt.Println(starr.Peek())
	// fmt.Println(starr.Pop())
	// fmt.Println(starr.Pop())
	// fmt.Println(starr.Pop())
	// starr.Print()

	// qll := stacks_and_queues.NewQueueLinkedList()
	// qll.Enqueue(1)
	// qll.Enqueue(2)
	// qll.Enqueue(3)
	// qll.Enqueue(4)
	// qll.Print()
	// qll.Enqueue(5)
	// qll.Print()
	// fmt.Println(qll.Peek())
	// fmt.Println(qll.Dequeue())
	// qll.Print()

	qs := stacks_and_queues.NewQueueStack()
	qs.Enqueue(1)
	qs.Enqueue(2)
	qs.Enqueue(3)
	qs.Enqueue(4)
	qs.Print()
	qs.Enqueue(5)
	qs.Print()
	fmt.Println(qs.Peek())
	fmt.Println(qs.Dequeue())
	qs.Print()
}