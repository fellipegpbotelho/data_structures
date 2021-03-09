## Linked Lists

> A linked list is a linear data structure where each element is a separate object.

or

> A linked list data structure includes a series of connected nodes. Each node store
> the data and the address of the next node.

So, the elements of a linked list are not stored in a contiguous location. It actually consists of a sequence of items
in linear order that are linked together using pointers. So, you have to access the data sequentially rather than
having random access like in array.

Each node of a linked list is made up of the value as well as the pointer. There are typically known as `key` for value
and `next` for the pointer.

You have to start somewhere, so we give the address of the first node a special name called `HEAD`.

Also, the last node in the linked list can be identified because its next portion points to `NULL` and we call him
`TAIL`.

### Types of Linked List

- Singly Linked List

    You are only to do transverse to forward.

    ![Image of singly linked list](https://cdn.programiz.com/sites/tutorial2program/files/linked-list-concept.png)

- Doubly Linked List

    You are able to do transverse forwards or backwards, and each node also has a previous pointer. It is a pointer to
    its previous predecessor node rather than its successor node.

    ![Image of doubly linked list](https://cdn.programiz.com/sites/tutorial2program/files/doubly-linked-list-concept.png)

- Circular Linked List

    Is a variation of a linked list in which the last element is linked to the first element. This form a circular loop.
    Is when the `TAIL` points to the `HEAD` and the `HEAD` points to the `TAIL`.

    ![Image of circular linked list](https://cdn.programiz.com/sites/tutorial2program/files/circular-linked-list.png)

### Advantages

- Dynamic Data Structure: Linked lists is a dynamic data structure so it can grow and shrink ant runtime by allocating
and deallocating memory. So there no need to give initial size of linked list.

- Insertion and Deletion: Insertion and Deletion of nodes are really easier. Unlike array here we don't have to shift
elements after insertion or deletion of an element. In linked list we just have to update the address present in next
pointer of a node.

- No Memory Wastage: As size of linked list can increase or decrease at runtime so there is no memory wastage. In case of array there is a lot of memory wastage, like if we declare an array os size 10 and store only 6 elements in it then space of 4 elements are wasted. There is no such problem in linked list as memory is allocated only when required.

- Implementation: Data structures such as stack and queues can be easily implemented using linked list.

### Disadvantages

- Memory Usage: More memory is requid to store elements in linked list as compared to array. Because in linked list each node contains a pointer and it requires extra memory for itself.

- Traversal: Elements or nodes traversal is difficult in linked list. We cannot randomly access any element as we do in array by index. For example if we want to access a node at position `n` then we have to traverse all the nodes before ir. So, time required to access a node is large.

- Reverse Traversing: In linked list reverse traversing is really difficult. In case of doubly linked list its easier but extra memory is required for back pointer hence wastage of memory.

### Operations

#### Transverse

Displaying the contents of a linked list is very simple. We keep moving the temp node to the next one and display its
contents.

When temp is `NULL`, we know that we have reached the end of the linked list so we get out of the while loop.

#### Insert

You can add elements to either the beginning, middle or the end of the linked list.

- Beginning: allocate memory for new node; store data; change next of new node to point do head; change head to point to
recently created node.

- End: allocate memory for new node; store data; tranverse to last node; change next of last node to recently created
node.

- Middle: allocate memory and store data for the new node; transverse to node just before the required position of new
node; change next pointers to include new node in between;

#### Delete

You can delete either from the beginning, end or from a particular position.

- Beginning: point head to the second node.

- End: transverse to second last element; change its next pointer to null.

- Middle: transverse to element before the element to be deleted; change next pointers to exclude the node from the
chain.
