//This is Data Struct hw3
//Implement Min-heap and Max-heap using array to find the median value of a list of number
#include <iostream>
#include <stdlib.h>
using namespace std;

void swap(int *x, int *y);
class minHeap{
    //a heap array;
    int *harr;
    int cap;
    public:
    int size;
    minHeap(int cap);
    //given the index, retrn the index
    int parent(int i) {return (i/2);}
    int left(int i) {return 2*i;}
    int right(int i) {return 2*i+1;}
    int top(void){return harr[1];}
    //pop the root
    void pop(void);
    //insert a number
    void push(int k);
    //maintain heap properties
    void minHeapify(int l);
    void show(void);
};

minHeap::minHeap(int cap){
    minHeap::cap = cap;
    size = 0;
    harr = new int[cap];
    harr[0] = 0;
}

void minHeap::minHeapify(int i){
    int l = left(i);
    int r = right(i);
    int smallest = i;
    if (l <= size && harr[l] < harr[i]){
        smallest = l;
    }
    if (r <= size && harr[r] < harr[smallest]){
        smallest = r;
    }
    if (smallest != i){
        swap(&harr[smallest], &harr[i]);
        minHeapify(smallest);
    }    


}

void minHeap::show(void){
    if (size == 0){
        return ;
    }
    int j = 2;
    for (int i = 1; i <= size; i++){
      	printf("%d", harr[i]);
        if (i == j-1 && j-1 != size){
            j = j * 2;
            printf("\n");
            continue;
        }
        if (i == j-1 && j-1 == size){
            printf("\n");
            continue;
        }
        printf(" ");

    }
    for (int i = size + 1; i < j; i++){
        printf("%c", 'S') ;
        if (i == j-1){
            printf("\n");
            break;
        }
        printf(" ");
    }
}

void minHeap::push(int k ){
    if (size == cap){
        exit(1);
    }
    if (size == 0){
        harr[1] = k;
        size++;
        return ;
    }
    size++;
    int i = size;
    harr[i] = k;
    while (harr[parent(i)] > harr[i] && i > 1){
        swap(&harr[parent(i)], &harr[i]);
        i = parent(i);
    }
    
}

void minHeap::pop(void){
    if (size <= 0){
        return ;
    }
    if (size == 1){
        size--;
        return ;
    }
    harr[1] = harr[size];
    size--;
    minHeapify(1);
    return ;
}
void swap(int *x, int *y){
    int t = *x;
    *x = *y;
    *y = t;
}

class maxHeap{
 //a heap array;
    int *harr;
    int cap;
    public:
    int size;
    maxHeap(int cap);
    //given the index, retrn the index
    int parent(int i) {return (i/2);}
    int left(int i) {return 2*i;}
    int right(int i) {return 2*i+1;}
    int top(void){return harr[1];}
    //pop the root
    void pop(void);
    //insert a number
    void push(int k);
    //maintain heap properties
    void maxHeapify(int l);
    void show(void);
};

maxHeap::maxHeap(int cap){
    maxHeap::cap = cap;
    size = 0;
    harr = new int[cap];
    harr[0] = 0;
}

void maxHeap::maxHeapify(int i){
    int l = left(i);
    int r = right(i);
    int largest = i;
    if (l <= size && harr[l] > harr[i]){
        largest = l;
    }
    if (r <= size && harr[r] > harr[largest]){
        largest = r;
    }
    if (largest != i){
        swap(&harr[largest], &harr[i]);
        maxHeapify(largest);
    }    


}

void maxHeap::show(void){
    if (size == 0){
        return ;
    }
    int j = 2;
    for (int i = 1; i <= size; i++){
        printf("%d",harr[i]);
        if (i == j-1 && j-1 != size){
            j = j * 2;
           printf("\n");
            continue;
        }
        if (i == j-1 && j-1 == size){
            printf("\n");
            continue;
        }
        printf(" ");
    }
    for (int i = size + 1; i < j; i++){
        printf("%s", "S") ;
        if (i == j-1){
            printf("\n");
            break;
        }
        printf(" ");
    }
}

void maxHeap::push(int k ){
    if (size == cap){
        exit(1);
    }
    if (size == 0){
        harr[1] = k;
        size++;
        return ;
    }
    size++;
    int i = size;
    harr[i] = k;
    while (harr[parent(i)] < harr[i] && i > 1){
        swap(&harr[parent(i)], &harr[i]);
        i = parent(i);
    }
    
}

void maxHeap::pop(void){
    if (size <= 0){
        return ;
    }
    if (size == 1){
        size--;
        return ;
    }
    harr[1] = harr[size];
    size--;
    maxHeapify(1);
    return ;
}

void balance(minHeap *l, maxHeap *s){
    int t;
    if (l->size > s->size+1){
        t = l->top();
        s->push(t);
        l->pop();
        return ;
    }
    if (s->size > l->size+1){
        t = s->top();
        l->push(t);
        s->pop();
        return ;
    }
}

int main(void){
    maxHeap mx(100000);
    minHeap mi(100000);
    int a, b;
  	int v1;
  	double v2;
    cin>>a>>b;
    if (a < b){
        mx.push(a);
        mi.push(b);
    }
    else {
        mx.push(b);
        mi.push(a);
    }
    if ((mi.top()+mx.top())%2==0){
      	v1 = (mi.top()+mx.top())/2;
        printf("%d\n",v1);
    }
    else{
        int m = mi.top();
        int n = mx.top();
        v2 = (m+n)/2;
        printf("%.1f\n",v2);
    }
    mi.show();
    mx.show();
    int k;
    while (cin >> k){
        if (k < mx.top()){
            mx.push(k);
        }else{
            mi.push(k);
        }    
        balance(&mi, &mx);
        if (mi.size > mx.size){
          	v1 = mi.top();
           printf("%d\n", v1);
        }
       else if (mx.size > mi.size){
         	v1 = mx.top();
          printf("%d\n",v1);
        }
        else{
            if ((mi.top()+mx.top())%2==0){
              	v1 = (mi.top()+mx.top())/2;
                 printf("%d\n", v1);
            }
            else{
                double m = mi.top();
                double n = mx.top();
              	v2 = (m+n)/2;
                printf("%.1f\n",v2);
            }       
        }
     mi.show();
     mx.show();
    }
}
