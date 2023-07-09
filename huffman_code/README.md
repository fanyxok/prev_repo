# Description

## Use Huffman codes for compression/decompression.

The input to compression is a sequence of 8-bit characters.

When computing the Huffman tree, do not compute the code for any character that is absent from the input (also, do not insert these characters into the min-heap).

To ensure consistent behavior between your program and the reference program during the delete operation on the min-heap, you need to determine the priority of the subtrees that have the same weight. 

Let S and T be two subtrees. S has a higher priority than T if and only if:

* S's weight is smaller than T's weight, or 

* S and T have the same weight, and the smallest character (in ASCII value) in any of S's leaves is smaller than that in any of T's leaves.

Under this definition, the delete operation should remove the subtree with the highest priority from the min-heap.

Also when merging two subtrees, set the subtree with the lower priority as the left subtree (whose code is 0), and the subtree with the higher priority as the right subtree (whose code is 1).
### Command line:

Read input from cin and write output to cout.Ex: ./huffman < input > output

## Input

### Uncompressed data

The uncompressed data contains a sequence of 8-bit characters. The input contains at most 2^32-1 characters.

### Compressed data

The compressed data contains three sections:

* Magic cookie. This section contains 8 characters: the string "HUFFMAN" followed by the ASCII 0 character ('\0').
* Frequencies. This section contains the frequencies of all the characters from ASCII 0 to ASCII 255, even if a characercharacter is absent from the uncompressed data. The frequency of a character is its count in the uncompressed data. Order the frequencies by the ASCII values of their corresponding characters. Write each frequency as a 4-byte unsigned integer in the little-endian form.
* Compressed data. This section contains the codes of all the characters in the same order as they appear in the uncompressed data. Since this section contains a sequence of bits but the smallest unit of data is a byte in files, you need to convert bits into bytes by the following rules:

Starting from the beginning of the bit sequence, convert each sequence of 8 consecutive bits into 1 byte. If the number of bits is not a multiple of 8, pad the end of the bit sequence with 0s.
When converting 8 bits into 1 byte, let the first bit be the least significant bit (LSB) in the byte, the second bit be the second LSB, and so on.We will test the decompression function of your program with only valid compressed data, so your program need not handle errors in the compressed data.
The first six hex of the input should either be "ENCODE" or "DECODE"

Then read the remaining input as the data.
When this argument is "ENCODE", compress the data and writes the compressed data to the output.
When this argument is "DECODE", decompress the data and writes the decompressed data to the output.

## Output

It is the same format as input data, and didn't have the "ENCODE" or "DECODE" prefix.
