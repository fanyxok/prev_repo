Task 1: Skip Lists in C
Instructions
Implement a Skip List in C. See the header file skiplist.h for further details.

Your submission should only include the source code for the Skip List, skiplist.c. 

Compilation of the source code requires no errors or warnings and conformity to the C89 standard. The Makefile syntax is given below.

CC=gcc
CFLAGS=-Wpedantic -Wall -Wextra -Werror -std=c89

skiplist.o: skiplist.c skiplist.h
  ${CC} ${CFLAGS} skiplist.c -o skiplist.o
The above compilation requires a main function. However, your submission should NOT include any main function.

We will test your program using the functions declared in the header file provided as well as some other functions defined by us. The output of the individual test cases will NOT be given in auto-grader and any attempt at retrieving the test cases or the program outputs will be seen as plagiarism.

We require that you will have one line with comments for every four non-blank lines. Comments have to be meaningful and in English.

No memory leaks are allowed. Memory leaks will be automatically detected and manually screened and will result in a decrease in your score after the deadline.

We will also impose a limit of 450 on your LoC (lines of code). Do not exceed this limit as we will deduct points for those who exceed this limit.

Consider Learning Before You Start
Pointers and References
C Memory Management
Function Pointers
Skip List
Skip List
A skip list (Pugh, 1990) is a randomized tree-like data structure based on linked lists. It consists of a level 0 list that is an ordinary sorted linked list, together with higher-level lists that contain a random sampling of the elements at lower levels. When inserted into the level i list, an element flips a coin that tells it with probability p to insert itself in the level i+1 list as well.

Searches in a skip list are done by starting in the highest-level list and searching forward for the last element whose key is smaller than the target; the search then continues in the same way on the next level down. The idea is that the higher-level lists act as express lanes to get us to our target value faster. To bound the expected running time of a search, it helps to look at this process backwards; the reversed search path starts at level 0 and continues going backwards until it reaches the first element that is also in a higher level; it then jumps to the next level up and repeats the process. On average, we hit 1 + 1/p nodes at each level before jumping back up; for constant p (e.g. 1/2), this gives us O(logn) steps for the search.

source: Yale CPSC223
There are various ways to implement a skip list. You can learn about the basics of skip list here. However, note that the implementation required in this homework differs from that on wikipedia

Implementation details
You are required to implement skiplists with following rules:

The height of each node is determined via geometric distribution on node creation. Randomly determine the level of a node by:
#define RAND_HALF (RAND_MAX)/2
node_level = 0;
while (node_level < s->height_limit - 1 && rand() < RAND_HALF) ++node_level;

The header of skip list is always pointing to a dummy node, whose key = NULL and value = NULL . The height of the dummy header always equals the height_limit for the tree list.
The initial height of an empty tree list is 1
The allocInt and deleteInt functions are used for allocating memory for int and copying its value.
Submission
Refer to the submission section at the end.

Task 2: Libraries
For all tasks assume that skiplist.h will be provided in the working directory and your skiplist.c will be copied to the current working directory.

Task 2.1: Static Library
Instructions
Write a shell script named staticlib.sh that when invoked, produces a static library named liblist.a from the source files skiplist.c and skiplist.h.

Then create a static-linked executable staticlist from liblist.a and test.c. Do not execute the program in the shell script.

Hint: use the command ar to create the static library. Then use the command ld to create the executable file.

Task 2.2 Dynamic-link Library
Instructions
Write a shell script named dynamiclib.sh that when invoked, produces a dynamic-linked library named liblist.so from skiplist.c.

Then generate an executable named dynamiclist using liblist.so as well as test.c. Do not execute the program in the shell script.

Submission
You should submit a compressed file named as skiplist.tar.

The directory tree of your submission should look like the following :

├── skiplist.c
├── dynamiclib.sh
└── staticlib.sh
Note: your submission should NOT contain main() function

Test Environment
The test environment on autolab is Ubuntu 16.04 with gcc 5.x

Last Modified: Mar. 9th, 2018
