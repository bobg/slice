# Slice - a demonstration of the workings of Go slices

This is slice,
a library illustrating how Go slices work.

The type `Slice[T]` in this library works just like `[]T` does in Go,
with copy, append, index, and subslicing operations.
Obviously if you’re writing code you should use Go’s intrinsic slice type,
and not this library.
But if you want to understand what’s happening under the covers,
take a look at how the methods here are implemented.

The specification for the Go language is quite clear and accessible.
You can find it at [go.dev/ref/spec](https://go.dev/ref/spec).
These subsections pertain specifically to slices:

- [Slice types](https://go.dev/ref/spec#Slice_types)
- [Slice expressions](https://go.dev/ref/spec#Slice_expressions)
- [Appending to and copying slices](https://go.dev/ref/spec#Appending_and_copying_slices)
- [Making slices, maps, and channels](https://go.dev/ref/spec#Making_slices_maps_and_channels)
