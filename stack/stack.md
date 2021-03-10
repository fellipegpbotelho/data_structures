## Stack

> A stack is a linear data structure which follows a particular order in which the
> operations are performed. The order may be LIFO(Last In First Out) or FILO(First In
> Last Out).

### Basic features

- Is an ordered list of similar data type.
- `push` function is used to insert new elements into the stack and `pop` function is
used to remove an element from the stack. Both insertion and removal are allowed at only
onde end of Stack called `Top`.
- Stack is said to be in `Overflow` state when it is completely full and is said to be in
`Underflow` state if it is completely empty.

### Applications

The simplest application of a stack is to reverse a word. You push a given word to stack
letter by letter ant then pop letters from the stack.

There are other uses also like:
    - Parsing
    - Expression Conversion

### Implementation

Stack can be easily implemented using an Array or a Linked List. Array are quick, but are
limited in size and linked list requires overhead to allocate, link, unlink and deallocate,
but is not limited in size.

#### Push operation

- Check if the stack is full or not.
- If the stack is full, the print error of overflow and exit the program.
- If the stack is not full, the increment the top and add the element.

#### Pop operation

- Check if the stack is empty or not.
- If the stack is empty, the print error of underflow and exit the program.
- If the stack is not empty, the print the element at the top and decrement the top.
