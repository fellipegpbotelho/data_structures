## Array

> An array is a container object that holds a fixed number os values of a single type.

or

> An array is a collection of homogeneous (same type) data items stored in contiguous memory locations.

<br />

![Image of array memory representation](https://beginnersbook.com/wp-content/uploads/2018/10/array.jpg)

```go
integers := []int{10, 20, 30, 40}
floats := []float32{3.50, 2.25, 1.00}
strings := []string{"String", "String"}
arrayOfArrays := [][]string{{"String"}, {"String"}}
```

When the array is created, his length is determined by the number of elements inside, and this will be his length.

Array has a few operations you may want to perform:

- Transverse: print all the array elements one by one.
- Search: searches an element using the given index or by the value.
- Update: updates an element at the given index.

* You can't insert and delete from an array. Because they are fixed in size you can't do that.

- Inserting: you would actually have to create a whole entire new array and you have to do current size plus one if
you are inserting something into that array and you would have to copy all existing elements out of your previous
and put into the new array that is of a larger size.

- Deleting: If you are deleting something from array, you would have to create a whole entire new array and you have
to do current size minus one and you would have to copy all existing elements out of your previous except for the one
you want to delete into that new array.

### Why we need an array?

Array is particularly useful when we are dealing with lot of variables of the same type. For example say i need to
store the marks in math subject of 100 students. To solve this particular problem, either i have to create the 100
variables of int type or create an array of int type with the size 100.

Obviosly the second option is best, because keeping track of all the 100 different variables is a tedious task. On the
other hand, dealing with array is simple and easy, all 100 values can be stored in the same array at different indexes
(0 to 99).

### Advantages

- Reading an array element is simple and efficient. The read time of array is O(1) in both best and worst cases. This
is because any element can be instantly read using indexes without tranversing the whole array.

- Array is a foundation of the other data structures. For example other data structures such as LinkedList, Stack,
Queue etc.

- All the elements of an array can be accessed using a single name along with the index, which is readable,
user-fliendly and efficient rather than storing those elements in different variables.

### Disadvantages

- While using array, we must need to make the decision of the size of the array in the beginning, so if we are not
aware how many elements we are going to store in array, it would make the task difficult.

- The size os the array is fixed, so if at later point we need to store more elements in it then it can't be done. On
the other hand, if we store less number of elements than the declared size, the remaining allocated memory is wasted.

### Related Algorithms:
- arraylists
- heaps
- hash tables
- matrices
- bubble sort
- quick sort
