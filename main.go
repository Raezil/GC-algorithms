package main

import (
	"fmt"
)

type Object struct {
	id         int
	marked     bool
	references []*Object
}

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

type GC struct {
	heap  []*Object
	roots []*Object
}

func NewGC() *GC {
	return &GC{
		heap:  []*Object{},
		roots: []*Object{},
	}
}

func (gc *GC) marking(obj *Object) {
	obj.marked = true
	for _, ref := range obj.references {
		gc.Mark(ref)
	}
}
func (gc *GC) Mark(obj *Object) {
	if obj == nil || obj.marked {
		return
	}
	gc.marking(obj)
}

func (gc *GC) Sweep() {
	for i := 0; i < len(gc.heap); i++ {
		obj := gc.heap[i]
		if !obj.marked {
			fmt.Printf("Object %d is unreachable and will be collected.\n", obj.id)
			gc.heap = append(gc.heap[:i], gc.heap[i+1:]...)
			i--
		} else {
			obj.marked = false
		}
	}
}

func (gc *GC) MarkSweep() {
	for _, root := range gc.roots {
		gc.Mark(root)
	}
	gc.Sweep()
}

func main() {
	gc := NewGC()
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
	gc.roots = []*Object{obj1}

	// Add all objects to the heap
	gc.heap = []*Object{obj1, obj2, obj3, obj4}

	fmt.Println("Before garbage collection:")
	for _, obj := range gc.heap {
		fmt.Printf("Object %d\n", obj.id)
	}

	// Run garbage collector
	gc.MarkSweep()

	fmt.Println("\nAfter garbage collection:")
	for _, obj := range gc.heap {
		fmt.Printf("Object %d\n", obj.id)
	}
}
