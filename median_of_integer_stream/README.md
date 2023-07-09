# Median of Integer Stream
## Description
 using a min heap and a max heap to get the median from an integer stream.

## Definition of Median
Median is defined as following: 

* If n is odd then Median (M) = value of ((n + 1)/2)th item term.
* If n is even then Median (M) = value of [(n/2)th item term + (n/2 + 1)th item term]/2

## Program Behaviour
When the input is one integer, your program should print the median after added the integer, and print the min heap and max heap now.

Your program should exit with no error, when the input is EOF.

## How to Get the Median
<p>For the first two elements add the smaller one to the maxHeap, and the bigger one to the minHeap. Then process the next added numbers one by one as follows:

* Step 1. Add the next item to one of the heaps
if the item is smaller than the root of maxHeap, add it to maxHeap, else add it to minHeap</p>

* Step 2. Balance the heaps (after this step heaps will be either balanced or one of them contains one more item)

if the number of elements in one of the heaps is greater than the other by more than 1, remove the root element from the one containing more
elements and add them to the other one.

Then you can calculate median like this:

If the heaps contain equal amount of elements:

` median = (root of maxHeap + root of minHeap)/2`

Else

 `median = root of the heap with more elements`
 
## How to Build Heap
In order to make your output same as the reference output, you should follow these rules when building heap.

### Binary Heap
Your should build two Binary Heap. One is Min Heap, another is Max Heap.

These heap should be Binary Heap at all times.

### Insert number
To add an element to a heap we must perform an up-heap operation (also known as bubble-up, percolate-up, sift-up, trickle-up, heapify-up, or cascade-up),

by following this steps:

`1\. Add the element to the last empty space of the heap(The right most empty space if lowest layer is not full,
or the left most item if the lowest layer is full).`

`2\. Compare the added element with its parent; if they are in the correct order, stop.`

`3\. If not, swap the element with its parent and return to the previous step.`

### Remove root
The procedure for deleting the root from the heap (effectively extracting the maximum element in a max-heap or the minimum element in a min-heap)

and restoring the properties is called down-heap (also known as bubble-down, percolate-down, sift-down, trickle down, heapify-down, cascade-down,

and extract-min/max).

`1\. Replace the root of the heap with the last element in the last level.`

`2\. Compare the new root with its children; if they are in the correct order, stop.`

`3\. If not, swap the element with one of its children and return to the previous step.`

(Swap with its smaller child in a min-heap and its larger child in a max-heap If children are the same, sway with the left one.)

## Input
In the first line we will give two integers. Then each line of the input is an integer. Integers are limited to int32.
## Output
First line is the median. Then you should print the min heap and the max heap.

To avoid the confusion with floating numbers, If the median is a floating number, its precision should be cut to 0.1. For example:

If the median is 1.5, the output should be exactly 1.5

If the median is 2, the output should be exactly 2

When printing the heap, you should print every layer of the heap. If there is a space you should print "S". If the heap likes the following:

        1
       / \
      /   \
     /     \
    2       3
   / \     / \
  4   5   6
  
you should print it like this:

`1
2 3
4 5 6 S`

There should be no more space at the end of each line.
