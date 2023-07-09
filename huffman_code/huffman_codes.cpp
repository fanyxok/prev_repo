#include <stdio.h>
#include <stdlib.h>
#include <cstring>
#include <string>
#include <iostream>
using namespace std;
#define MAX_TREE_HT 256
/* A Huffman tree node*/
struct MinHeapNode {
    int data;
    unsigned freq;
    int priority;
    struct MinHeapNode *left, *right;
};
 
/* A Min Heap:  Collection of min heap (or Hufmman tree) nodes*/
struct MinHeap {
    unsigned size;
    unsigned capacity;
    /* Attay of minheap node pointers */
    struct MinHeapNode** array;
};
struct MinHeapNode* newNode(char data, unsigned freq)
{
    struct MinHeapNode* temp = (struct MinHeapNode*)malloc(sizeof(struct MinHeapNode));
    temp->left = temp->right = NULL;
    temp->data = data;
    temp->freq = freq;
    temp->priority = (int) data;
 
    return temp;
}
struct MinHeap* createMinHeap(unsigned capacity)
{
    struct MinHeap* minHeap = (struct MinHeap*)malloc(sizeof(struct MinHeap));
    /* current size is 0*/
    minHeap->size = 0;
    minHeap->capacity = capacity;
    minHeap->array = (struct MinHeapNode**)malloc(minHeap->
capacity * sizeof(struct MinHeapNode*));
    return minHeap;
}

void swapMinHeapNode(struct MinHeapNode** a,
                     struct MinHeapNode** b)
{
    struct MinHeapNode* t = *a;
    *a = *b;
    *b = t;
}
/* The standard minHeapify function.*/
void minHeapify(struct MinHeap* minHeap, int idx)
{
    int smallest = idx;
    int left = 2 * idx + 1;
    int right = 2 * idx + 2;
    int len  = minHeap->size;
    if (left < len && (minHeap->array[left]->
freq < minHeap->array[smallest]->freq))
        smallest = left;
    if (right < len && minHeap->array[right]->
freq < minHeap->array[smallest]->freq)
        smallest = right;
    if (smallest == idx){
        if (left < len && (minHeap->array[left]->freq == minHeap->array[smallest]->freq) &&
                (minHeap->array[left]->priority < minHeap->array[smallest]->priority))
            smallest = left;
        if (right < len && (minHeap->array[right]->freq == minHeap->array[smallest]->freq) &&
                (minHeap->array[right]->priority < minHeap->array[smallest]->priority))
            smallest = right;
    }
    else if (smallest == left){
        if (right < len && (minHeap->array[right]->freq == minHeap->array[smallest]->freq) &&
                (minHeap->array[right]->priority < minHeap->array[smallest]->priority)){
            smallest = right;
        }
    }
    if (smallest != idx) {
        swapMinHeapNode(&minHeap->array[smallest],
                        &minHeap->array[idx]);
        minHeapify(minHeap, smallest);
    }
}
 
/* A utility function to check if size of heap is 1 or not*/
int isSizeOne(struct MinHeap* minHeap)
{
    return (minHeap->size == 1);
}
// A standard function to extract minimum value node from heap*/
struct MinHeapNode* extractMin(struct MinHeap* minHeap)
 
{
 
    struct MinHeapNode* temp = minHeap->array[0];
    minHeap->array[0] = minHeap->array[minHeap->size - 1];
    --minHeap->size;
    minHeapify(minHeap, 0);
    return temp;
}
 
// A utility function to insert
// a new node to Min Heap
void insertMinHeap(struct MinHeap* minHeap,
                   struct MinHeapNode* minHeapNode)
{
    ++minHeap->size;
    int i = minHeap->size - 1;
    while (i && minHeapNode->freq < minHeap->array[(i - 1) / 2]->freq) {
        minHeap->array[i] = minHeap->array[(i - 1) / 2];
        i = (i - 1) / 2;
    }
    while (i && minHeapNode->freq == minHeap->array[(i-1)/2]->freq){
        if (minHeapNode->priority < minHeap->array[(i - 1)/2]->priority){
            minHeap->array[i] = minHeap->array[(i-1)/2];
            i = (i - 1)/2;
            continue;
        }
        break;
    }
    minHeap->array[i] = minHeapNode;
}
 
// A standard funvtion to build min heap
void buildMinHeap(struct MinHeap* minHeap)
{
    int n = minHeap->size - 1;
    int i;
    for (i = (n - 1) / 2; i >= 0; --i)
        minHeapify(minHeap, i);
}
// A utility function to print an array of size n
void printArr(int arr[], int n)
{
    int i;
    for (i = 0; i < n; ++i)
        printf("%d", arr[i]);
    printf("\n");
}
 
// Utility function to check if this node is leaf
int isLeaf(struct MinHeapNode* root)
{
    return !(root->left) && !(root->right);
}
// Creates a min heap of capacity
// equal to size and inserts all character of
// data[] in min heap. Initially size of
// min heap is equal to capacity
struct MinHeap* createAndBuildMinHeap(int data[], int freq[], int size)
{
    struct MinHeap* minHeap = createMinHeap(size);
    int rsize= 0;
    for (int i = 0; i < size; ++i){
        if (freq[i] != 0)
        {   
            rsize ++;
            minHeap->array[i] = newNode(data[i], freq[i]);
        }
    }
    minHeap->size = rsize;
    buildMinHeap(minHeap);
    return minHeap;
}
// The main function that builds Huffman tree
struct MinHeapNode* buildHuffmanTree(int data[], int freq[], int size)
 
{
    struct MinHeapNode *left, *right, *top;
 
    // Step 1: Create a min heap of capacity
    // equal to size.  Initially, there are
    // modes equal to size.
    struct MinHeap* minHeap = createAndBuildMinHeap(data, freq, size);
 
    // Iterate while size of heap doesn't become 1
    while (!isSizeOne(minHeap)) {
 
        // Step 2: Extract the two minimum
        // freq items from min heap
        right = extractMin(minHeap);
        left = extractMin(minHeap);
        // Step 3:  Create a new internal
        // node with frequency equal to the
        // sum of the two nodes frequencies.
        // Make the two extracted node as
        // left and right children of this new node.
        // Add this node to the min heap
        // '$' is a special value for internal nodes, not used
        top = newNode('$', left->freq + right->freq);
        top->left = left;
        top->right = right;
        if (left->priority < right->priority)   {top->priority = left->priority;}
        else    {top->priority = right->priority;}
        insertMinHeap(minHeap, top);
    }
 
    // Step 4: The remaining node is the
    // root node and the tree is complete.
    return extractMin(minHeap);
}
 
// Prints huffman codes from the root of Huffman Tree.
// It uses arr[] to store codes
void printCodes(struct MinHeapNode* root, int arr[], int top, string encode[256])
{
    // Assign 0 to left edge and recur
    if (root->left) {
        arr[top] = 0;
        printCodes(root->left, arr, top + 1, encode);
    }
    // Assign 1 to right edge and recur
    if (root->right) {
        arr[top] = 1;
        printCodes(root->right, arr, top + 1, encode);
    }
    // If this is a leaf node, then
    // it contains one of the input
    // characters, print the character
    // and its code from arr[]
    if (isLeaf(root)) {
        string code;
        for (int i=0; i < top; i++){
            code = code + to_string(arr[i]);
        }
        int ridx = root->data;
        if (ridx < 0)   ridx = 256 + root->data;
        //cout << "is leaf:"<<ridx<<endl;
        encode[ridx] = code;
        //printf("%c:", root->data);
        //cout<<"|"<<code<<"|"<<endl;
        //printArr(arr, top);
    }
}
// The main function that builds a
// Huffman Tree and print codes by traversing
// the built Huffman Tree
void HuffmanCodes(int data[], int freq[], int size, string encode[256])
{
    // Construct Huffman Tree
    struct MinHeapNode* root = buildHuffmanTree(data, freq, size);
    // Print Huffman codes using
    // the Huffman tree built above
    int arr[MAX_TREE_HT], top = 0;
    printCodes(root, arr, top, encode);
}

uint32_t big2lit(uint32_t num){
    uint32_t b0, b1, b2, b3;
    b0 = (num & 0x000000ff)<<24u;
    b1 = (num & 0x0000ff00)<<8u;
    b2 = (num & 0x00ff0000)>>8u;
    b3 = (num & 0xff000000)>>24u;
    return num;
}

void Encode(void){
    int freq[256]={0};
    int rf[256];
    int ra[256];
    int rsize = 0;
    string encode[256];
    string s;
    string S;
    char in;
    int input[1000000];
    int size = 0;
    while (cin.get(in)){
        int inn = in;
        if (inn < 0) inn = 256+in;
        if (inn < 0 || inn >= 256) {continue;}
        input[size]= inn;
        size++;
        //cout << (int)input[size-1]<<endl;
     }
    if (size == 0){
        char huffman[8];
        strcpy(huffman,"HUFFMAN");
        for (int i = 0; i <(int)strlen(huffman); i++){
            cout.put(huffman[i]);
        }
        cout.put('\0');
        for (int i = 0; i< 256; i ++){
            uint8_t zero = 0;
            cout.put(zero);
            cout.put(zero);
            cout.put(zero);
            cout.put(zero);
        }
        return ;  
    }
    //string copy = input;
    int len = size;
    for (int i=0; i < len; i++ ){
        int idx = input[i];
        //if (idx < 0)    idx = 256 + idx;
        freq[idx]++;
    }
    for (int i = 0; i < 256; i++){
        if (freq[i] != 0){
            rf[rsize] = freq[i];
            ra[rsize] = i;
            rsize++;
        }
    }
    //for (int i = 0; i < rsize; i++){
    //   printf("%c:%d:%d\n",ra[i],(unsigned)ra[i],rf[i]);
    //}
    HuffmanCodes(ra, rf, rsize, encode);
    //for (int i = 0; i < 256; i++){
    //    cout<<(char)i<<":"<<encode[i]<<"|"<<freq[i]<<endl;
    //}
    //cout.put('D');
    //cout.put('E');
    //cout.put('C');
    //cout.put('O');
    //cout.put('D');
    //cout.put('E');
    string out;
    char huffman[8];
    strcpy(huffman,"HUFFMAN");
    for (size_t i = 0; i <strlen(huffman); i++){
        cout.put(huffman[i]);
    }
    cout.put('\0');
    for (int i = 0; i < 256; i++){
        if (freq[i] == 0){
            uint8_t zero = 0;
            cout.put(zero);
            cout.put(zero);
            cout.put(zero);
            cout.put(zero);

        }
        else{
            uint32_t num = big2lit(freq[i]);
            cout.put(num & 0x000000ff);
            cout.put((num & 0x0000ff00)>>8u);
            cout.put((num & 0x00ff0000)>>16u);
            cout.put((num & 0xff000000)>>24u);
        }
    }
    string encodes;
    for (int i = 0; i < len; i++){
            //int idx = copy.c_str()[i];
            int idx = input[i];
            string bit = encode[idx];
            //cout << idx<<"|"<<bit<<endl;
            encodes = encodes + bit;
    }
    int codelen = encodes.size();
    if (codelen%8 != 0){
        int need = 8 - codelen%8;
        for (int k = 0; k < need; k++){
            encodes = encodes + "0";
        }
        codelen = codelen + need;
    }
    int bytes = codelen / 8;
    for (int k = 0; k < bytes; k ++){
        string numcode = encodes.substr(k*8, 8);
         char t = 0;
        for (int l = 0; l < 8; l++){
            char ith = numcode.c_str()[l];
            if (ith == '1'){
                t = t + (1 << l);
            }
        }
        cout.put(t);
    }
}

void Decode(){
    char magic[8];
    cin.get(magic,9);
    if (strcmp(magic, "HUFFMAN\0") != 0)  {return ;}
    int freq[256];
    int total = 0;
    char b0, b1, b2, b3;
    for (int i = 0; i < 256; i++){
        cin.get(b0);
        cin.get(b1);
        cin.get(b2);
        cin.get(b3);
        int fre = b0 & 0xFF;
        fre |= (b1 << 8) & 0xFF00;
        fre |= (b2 << 16) & 0xFF0000;
        fre |= (b3 << 24) & 0xFF000000;
        freq[i] = fre;
    }
    int rf[256] = {0};
    int ra[256];
    int rsize = 0;
    for (int i = 0; i < 256; i++){
        if (freq[i] != 0){
            rf[rsize] = freq[i];
            ra[rsize] = (unsigned char) i;
            rsize++;
            total += freq[i];
        }
    }
    if (rsize == 0)     {return ;}
    string encode[256];
    HuffmanCodes(ra, rf, rsize, encode);
    //for (int i = 0; i < 256; i++){
    //    cout<<(char)i<<":"<<encode[i]<<"|"<<freq[i]<<endl;
    //}
    char get;
    string codes;
    while (cin.get(get)){
        string byte = "";
        for (uint8_t i = 0; i < 8; i++){
            if (((get >> i) & 1) == 1){
                byte += "1";
            }else{
                byte += "0";
            }
        }
        codes += byte;
    }
    int len = codes.size();
    string s = "";
    for (int i = 0; i < len; i++){
        s += codes[i];
        for (int k = 0; k < 256; k++){
            if (encode[k] == s){
                cout.put(char(k));
                total -= 1;
                if (total == 0) return ;
                s = "";
                break;
            }
        }
    }
}

int main()
{   
    char opcode[6];
    cin.get(opcode, 7);
    if (strcmp(opcode, "ENCODE")==0){
        ;
    }
    else if (strcmp(opcode, "DECODE")==0){
        Decode();
    }
    return 0;
}
