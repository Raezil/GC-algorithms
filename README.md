# mark-sweep

```
Mark sweep algorithm.

Mark phase:
- it traverses through all roots and all roots' references marking all objects.
Sweep phase:
- once marked it traverses through all heaps' objects and deallocate unreachable objects (those who has not been marked).
- heap is graph object.
```

```
  Three color marking starts initially gray root object.
  GC explores gray objects, turning them black
  All white objects referenced by gray objects are turned gray, meaning they will be explored in the future

  This continues until there are no more gray objects.
  All objects initially are white on the heap.

  After all reachable objects have been marked black, the remaining white objects are considered
  unreachable and can be safely deallocated.
```

