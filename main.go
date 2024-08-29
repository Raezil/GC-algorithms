package main

import (
	"fmt"
)

type Object struct {
	id         int
	marked     bool
	references []*Object
}

var heap []*Object
var roots []*Object

func NewObject(id int) *Object {
	return &Object{
		id:         id,
		marked:     false,
		references: []*Object{},
	}
}

func (o *Object) AddReference(obj *Object) {
	o.references = append(o.references, obj)
}
func marking(obj *Object) {
	obj.marked = true
	for _, ref := range obj.references {
		mark(ref)
	}
}
func mark(obj *Object) {
	if obj == nil || obj.marked {
		return
	}
	marking(obj)
}

func sweep() {
	for i := 0; i < len(heap); i++ {
		obj := heap[i]
		if !obj.marked {
			fmt.Printf("Object %d is unreachable and will be collected.\n", obj.id)
			heap = append(heap[:i], heap[i+1:]...)
			i--
		} else {
			obj.marked = false
		}
	}
}

func markSweep() {
	for _, root := range roots {
		mark(root)
	}
	sweep()
}

func main() {
	// Create objects
	obj1 := NewObject(1)
	obj2 := NewObject(2)
	obj3 := NewObject(3)
	obj4 := NewObject(4)

	// Set up references
	obj1.AddReference(obj2)
	obj2.AddReference(obj3)
	// obj4 is unreachable, so it should be collected

	// Set up root objects
	roots = []*Object{obj1}

	// Add all objects to the heap
	heap = []*Object{obj1, obj2, obj3, obj4}

	fmt.Println("Before garbage collection:")
	for _, obj := range heap {
		fmt.Printf("Object %d\n", obj.id)
	}

	// Run garbage collector
	markSweep()

	fmt.Println("\nAfter garbage collection:")
	for _, obj := range heap {
		fmt.Printf("Object %d\n", obj.id)
	}
}
