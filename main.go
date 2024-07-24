package main

import (
	"fmt"

	"github.com/golangframework/system/Collections"
)

func main() {
	dictionary := make(Collections.Dictionary[int, string])
	dictionary.Add(1, "ala")
	dictionary.Add(2, "basia")
	// keys := dictionary.GetKeys()
	// for _, key := range keys {
	// 	fmt.Println((key))
	// }
	// values := dictionary.GetValues()
	// for _, key := range values {
	// 	fmt.Println((key))
	// }

	itemS := dictionary[1]
	fmt.Println(itemS)

	itemS, ok := dictionary[1]
	if ok {
		fmt.Println(itemS)
	}

	queue := make(Collections.Queue[int], 0)
	queue.Enqueue(11)
	queue.Enqueue(12)
	queue.Enqueue(13)
	if len(queue) > 0 {
		item := queue.Dequeue()
		fmt.Println(item)
	}
	item := queue.Dequeue()
	fmt.Println(item)
	item = queue.Dequeue()
	fmt.Println(item)

	stack := make(Collections.Stack[int], 0)
	stack.Push(11)
	stack.Push(12)
	stack.Push(13)
	item, _ = stack.Pop()
	fmt.Println(item)
	item, _ = stack.Pop()
	fmt.Println(item)
	item, _ = stack.Pop()
	fmt.Println(item)

}
