# Slice - a demonstration of the workings of Go slices

This is slice,
a library illustrating how Go slices work.

The type `Slice[T]` in this library works just like `[]T` does in Go,
with copy, append, index, and subslicing operations.
Obviously if you’re writing code you should use Go’s intrinsic slice type,
and not this library.
But if you want to understand what’s happening under the covers,
take a look at how the methods here are implemented.
