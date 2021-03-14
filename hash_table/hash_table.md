## Hash Table

> A hash table is a data structure that stores values which have keys
> associated with each of them.

In a hash table, data is stored in an array format, where each value
has its own unique index value. Access of data becomes very fast if
we know the index ot the desired data.

Thus, it becomes a data structure in which insertion and search
operations are very fast irrespective of the size of the data.
Hash Table uses an array as a storage medium and uses hash technique
to generate an index where an element is to be inserted or is to be
located from.

The average time complexity of hash tables is O(1)

### Hashing

Hashing is a technique to convert a range of key values into a range
of indexes of an array. We're going to use modulo operator to get a
range of key values.

![Image of hash table structure](https://www.tutorialspoint.com/data_structures_algorithms/images/hash_function.jpg)

A collision occurs when two keys are hashed to the same index in a
hash table. Collisions are a problem because every slot in a hash
table is supposed to store a single element.

Chaining is a technique used for avoiding collisions in a hash table
by making the hash table an array of linked lists, each index has its
own linked list.

Probind solves collision by finding empty slots in the hash table.

### Operations

- Search: searches an element in a hash table
- Insert: inserts an element in a hash table
- Delete: deletes an element from a hash table

### Real-world applications

In the real-world, hash tables are used to store data for:

- Databases
- Associative arrays
- Dict (python)
- Sets
- Memory cache

### Advantages

- Hash tables have high performance when looking up data, insert, and deleting
existing values.
- The time complexity for hash table is constant regardless of the number of
items in the table.
- They perform very well even when working with large datasets.

### Disadvantages

- You cannot use a null value as a key.
- Collisions cannot be avoided when generating keys using hash functions.
Collisions occur when a key thats is already in use is generated.
- If the hashing function has many collisions, this can lead to performance
decrease.
