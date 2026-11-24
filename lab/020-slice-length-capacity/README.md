# #20: Not understanding slice length and capacity

- initializing a slice `s := make([]int, len, cap)
- appending to a slice `s = append(s, 342)` or `s = append(s, 2, 3, 4)`  # to append multiple items

When it reaches the capacity it will create a new array double the size of the existing one and copy all the elements to that.

_slice length_ is the number of available elements in the slice, whereas the _slice capacity_ is the number of elements in the backing array.

(lenght == capacity) leads  to creating a new backing array with a new capacity, copying all the elements from the previous array, and updating the slice pointer to the new array.


