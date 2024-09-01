# GC summary

```
Mark sweep algorithm.

Mark phase:
- it traverses through all roots and all roots' references marking all objects.
Sweep phase:
- once marked it traverses through all heaps' objects and deallocate unreachable objects (those who has not been marked).
- heap is graph object.
```

