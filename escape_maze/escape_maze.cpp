#include <string>
#include <iostream>
#include <cstdlib>
#include <cstdio>
#include <cstring>
using namespace std;
typedef struct node node;
typedef struct point point;

struct node
{
	point *p;
	node *next;
};

struct point
{
	int x;
	int y;
	bool v;
	int d;
	point *previous;
	node *adj; 
};
void insert(point *poi, point * n)
{
	if (poi->adj == NULL)
	{
		poi->adj = new node();
		poi->adj->p = n;
		poi->adj->next = NULL;
		return ;
	}
	node *curr= poi->adj;
	while(curr->next)
	{
		curr = curr->next;
	}
	curr->next = new node();
	curr->next->p = n;
	curr->next->next = NULL;
	return ;
}
void swap(point **x, point **y){
    point *t = *x;
    *x = *y;
    *y = t;
}
/*class queue
{
private:
	int cap;
	int size; 
public:
	point *array[];
	
	queue(int cap);
	void enqueue(point *poi);
	point *dequeue();
	int parent(int i) {return (i/2);}
    int left(int i) {return 2*i;}
    int right(int i) {return 2*i+1;}
    point *top(void){return array[1];}
    void minHeapify(int l);
    bool isempty() {return size;}
};

queue::queue(int cap)
{
    size = 0;
    queue::cap = cap;
    //array = ;
    array[0] = new point();
    array[0]->d = 99999999;
    array[0]->y = -1;
    array[0]->x = -1;
}
void queue::minHeapify(int i){
    int l = left(i);
    int r = right(i);
    int smallest = i;
    if (l <= size && array[l]->d < array[i]->d){
        smallest = l;
    }
    if (r <= size && array[r]->d < array[smallest]->d){
        smallest = r;
    }
    if (smallest == i)
    {
    	if (l <= size && array[l]->d == array[i]->d && array[l]->previous && !array[i]->previous){
        	smallest = l;
    	}
    	if (r <= size && array[r]->d == array[smallest]->d && array[r]->previous && !array[i]->previous){
        	smallest = r;
    	}
    }else if (smallest == l){
    	if (r <= size && array[r]->d == array[smallest]->d && array[r]->previous && !array[l]->previous){
        	smallest = r;
    	}
    }
    if (smallest != i){
        swap(&array[smallest], &array[i]);
        minHeapify(smallest);
    }    
}
void queue::enqueue(point *k ){
    if (size == cap){
        cout << "ENQUEUE ERROR size = "<<size<<';'<<"cap = "<<cap<<endl;
    }
    if (size == 0){
        array[1] = k;
        size++;
        return ;
    }
    size++;
    int i = size;
    array[i] = k;
    while ((array[parent(i)]->d > array[i]->d & i > 1) ){
        swap(&array[parent(i)], &array[i]);
        i = parent(i);
	}
}

point *queue::dequeue(void){
    if (size <= 0){
        return NULL;
    }
    if (size == 1){
        size--;
        return top();
    }
    array[1] = array[size];
    size--;
    point *curr = top();
    minHeapify(1);
    return curr;
}*/
class maze
{
private:
	point *Points[1000][1000];
	point *list[2000];
	string map[1000];
public:
	int start[2];
	int end[2];
	maze();
	int n;
	int m;
	int size;
	void escape();
	void update(point *poi);
};
maze::maze()
{
	string c;
	n = 0;
	size = 0;
	int i,j;
	while (getline(cin,c))
	{
		if (strlen(c.c_str())!=0)
		{	map[n]= c;
			n++;
		}
	}
	m = strlen(map[0].c_str());
	//cout << "m="<<m<<','<<"n="<<n<<endl;
	//for (i = 0; i < n;i++)
	//		cout<<map[i]<<endl;
	for (i = 0; i < n; i++)
	{
		for (j = 0; j < m; j++)
		{
			char now = map[i].c_str()[j];
			if (now == ' ' || now == '+' || now == '*')
			{	
				point *poi = new point();
				poi->x = i;
				poi->y = j;
				poi->v = false;
				poi->d = 99999999;
				poi->previous = NULL;
				poi->adj = NULL;
				Points[i][j] = poi;
				list[size] = poi;
				size++;
				if (now == '*')
				{
					start[0]=i;
					start[1]=j;
				}
				if (now == '+')
				{
					end[0] = i;
					end[1] = j;
				}
			}
		}
	}
	for (i = 0; i < n; i++)
	{
		for (j = 0; j < m; j++)
		{
			if (map[i][j] == ' ' || map[i][j] == '+' || map[i][j] == '*')
			{	
				//cout << "("<<i<<','<<j<<"):"<<Points[i][j]->x<<'.'<<Points[i][j]->y<<endl;
				if (i>=1)
				{
					if(map[i-1][j] == ' ' || map[i-1][j] == '+' || map[i-1][j] == '*' )//maybe have trouble
					{
						insert(Points[i][j],Points[i-1][j]);
					}
				}
				if (j>=1)
				{
					if(map[i][j-1] == ' ' || map[i][j-1] == '+' || map[i][j-1] == '*' )//maybe have trouble
					{
						insert(Points[i][j],Points[i][j-1]);
					}
				}
				if (i <= n-1)
				{
					if(map[i+1][j] == ' ' || map[i+1][j] == '+' || map[i+1][j] == '*' )//maybe have trouble
					{
						insert(Points[i][j],Points[i+1][j]);
					}
				}
				if (j <= m-1)
				{
					if(map[i][j+1] == ' ' || map[i][j+1] == '+' || map[i][j+1] == '*' )//maybe have trouble
					{
						insert(Points[i][j],Points[i][j+1]);
					}
				}
			}
		}
	}
	/*if (c == '*')
		{
			start = [idx, idy];
		}else if (c == '+')
		{
			end = [idx, idy];
		}
		if (c == ' ')
		{
			point n = new point;
			n->x = idx;
			n->y = idy;
		}
	*/
}
void maze::update(point* poi)
{
	int od,nd;
	od = poi->d;
	node *curr = poi->adj;
	//cout << '('<<poi->x<<'.'<<poi->y<<"):";
	int d = od + 1;
	while (curr)
	{
		if (d < curr->p->d)
		{
			//cout<< '('<<curr->p->x<<'.'<<curr->p->y<<")-("<<curr->p->d<<"->"<<d<<") ";
			curr->p->d = d;
			curr->p->previous = poi;
		}
		
		curr = curr->next;
	}
	//cout<<endl;
	return ;
}
void maze::escape()
{
	point *curr = Points[start[0]][start[1]];
	curr->d = 0;
	point *dump = new point();
	point *smallest = dump;
	smallest->d = 99999999+1;
	smallest->x = -1;
	smallest->y = -1;
	int count = size;
	while (count--)
	{
		for (int i = 0; i < size; i++)
		{
			if (list[i]->d < smallest->d && list[i]->v == false)
			{
				smallest = list[i];
			}
		}
		update(smallest);
		smallest->v = true;
		smallest = dump;
	}
	/*for (int i = 0; i < m; i++)
	{
		for (int j = 0; j < n; j++)
		{
			if (map[i][j] == ' ' || map[i][j] == '+' || map[i][j] == '*')
			{	
				cout << "("<<i<<','<<j<<"):"<<Points[i][j]->previous<<endl;
			}
		}
	}*/
	int step = 0;
	point *poi = Points[end[0]][end[1]];
	if (!poi->previous)	
	{
		cout << "NO SOLUTION"<<endl;
		return ;
	}
	while(!(poi->x == start[0] && poi->y == start[1]))
	{
		map[poi->x][poi->y] = '.';
		//cout << poi->previous<<endl;
		poi = poi->previous;
		step++;
	}
	map[start[0]][start[1]] = '.';
	for (int i = 0; i < n;i++)
			cout<<map[i]<<endl;
	cout << step<<endl;
}
/*void maze::escape()
{
	point *curr = Points[start[0]][start[1]];
	curr->d = 0;
	queue *q = new queue(100);
	queue *Q = new queue(100);
	for (int i = 0; i < size; i++)
	{
		q->enqueue(list[i]);
	}
	for (int i = 1; i < 17; i++)
	{
		cout << q->array[i]->x<<'.'<<q->array[i]->y<<"|"<<q->array[i]->d<<endl;
	}
	point *poi = NULL;
	point *ps = NULL;
	cout << "start at :"<<start[0]<<'.'<<start[1]<<endl;
	cout <<"update :";
	int j = size;
	while(j--)
	{
		poi = q->top();
		update(poi);
		poi->v = true;
		q->dequeue();
		point *rash[100];
		int count = 0;
		while(q->isempty())
		{
			rash[count] = q->dequeue();
			count++;
		}
		for (int i = 0; i < count;i++)
		{
			q->enqueue(rash[i]);
		}
		cout << " "<< poi->x<<'.'<<poi->y <<','<<poi->d<<endl;	
	}
	exit(1);
	for (int i = 0; i < m; i++)
	{
		for (int j = 0; j < n; j++)
		{
			if (map[i][j] == ' ' || map[i][j] == '+' || map[i][j] == '*')
			{	
				cout << "("<<i<<','<<j<<"):"<<Points[i][j]->previous<<endl;
			}
		}
	}
	ps= Points[end[0]][end[1]];
	int step = 0;

	while(!(poi->x == start[0] && poi->y == start[1]))
	{
		map[poi->x][poi->y] = '.';
		cout << poi->previous<<endl;
		poi = poi->previous;
	}
	map[start[0]][start[1]] = '.';
	for (int i = 0; i < n;i++)
			cout<<map[i]<<endl;
}
*/
int main()
{
	maze *m = new maze();
	m->escape();
	return 0;
}
