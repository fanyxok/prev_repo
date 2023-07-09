#include <string>
#include <iostream>
#include <cstdlib>
#include <cstdio>
using namespace std;
typedef struct vertex vertex;

struct node
{
	vertex* head;
	node* next;
};
struct vertex
{
	char letter;
	node* next;
};
int index(char c)
{
	return (c-97);
}

class toposort
{
	public:
		int orders[26];
		vertex* LinkList;
		vertex** queue;
		int ihead;
		int itail;
		bool valid;
		toposort();
		vertex* pop();
		void push(vertex *u);
		bool isEmpty();
		void sort();
		void print();
		bool decrement(vertex* v);
		bool insert(struct vertex* v, struct vertex* u);
};

toposort::toposort()
{
	string input;
	string raw_input;
	string buffer;
	string n;
	int nums;
	char u, v;
	int len, rlen;
	LinkList = new vertex[26];
	ihead = 0;
	itail = -1;
	valid = true;
	for (int i = 0; i < 26; i++)
	{
		LinkList[i].letter = (char)(i+97);
		LinkList[i].next = NULL;
		orders[i] = 0;
	}
	getline(cin, n);
	nums = atoi(n.c_str());
	getline(cin, buffer);
	for (int i = 0; i < buffer.length(); i++)
	{
		if (buffer[i] == ' ')
			continue;
		raw_input+=tolower(buffer[i]);
	}
	rlen = raw_input.length();
	for (int i = 1; i < nums; i++)
	{
		getline(cin, buffer);
		for (int j = 0; j < buffer.length(); j++)
		{
			if (buffer[j] == ' ')
				continue;
			input+=tolower(buffer[j]);
		}
		len = input.length();
		int id = 0; 
		while (raw_input[id] == input[id] && id < len && id < rlen)
		{
			id++;
		}
		if (id < rlen && id >= len)
		{
			valid = false;
			return ;
		}else if (id >= len || id >= rlen){
			;
		}else{
			valid = insert(&LinkList[index(raw_input[id])], &LinkList[index(input[id])]);
			orders[index(input[id])]++;
		}	
		raw_input = input;
		rlen = len;
		input = "";
	}
	/*for (int i = 0; i < 26; i++)
	{
		if (LinkList[i].next)
		cout<<LinkList[i].letter<<":"<<LinkList[i].next->head->letter<<endl;
	}*/
	/*
	for (int i = 0; i < 26; i++)
	{
		cout<<(char)(i+97)<<":"<<orders[i]<<endl;
	}
	*/
	queue = new vertex*[27];
}

bool toposort::decrement(vertex* v)
{
	node* temp = v->next;
	while (temp)
	{
		if (orders[index(temp->head->letter)] <= 0)
		{
			return false;
		}
		(orders[index(temp->head->letter)])--;
		temp = temp->next;
	}
	return true;
}

bool toposort::insert(vertex* v, vertex* u)
{
	if (v == u)
		return false;
	if (!v->next)
	{
		v->next = new node;
		v->next->head = u;
		v->next->next = NULL;
		return true;
	}
	node* temp = v->next;
	while (temp->next)
	{	
		if (temp->head == u)
		{
			return false;
		}
		temp = temp->next;
	}	
	temp->next = new node;
	temp->next->head = u;
	temp->next->next = NULL;
	return true;
}

vertex* toposort::pop()
{
	vertex* i = queue[ihead];
	++ihead; 
	return i;
}

void toposort::push(vertex* u)
{
	++itail;
	queue[itail] = u;
}

bool toposort::isEmpty()
{
	return (itail+1)==ihead;
}

void toposort::sort()
{	
	if (!valid)
	{
		return ;
	}
	for (int i = 0; i < 26; i++)
	{
		if (orders[i] == 0)
		{
			push(&LinkList[i]);
			orders[i] = -1;
			break;
		}
	}
	while (!isEmpty())
	{
		vertex* v = pop();
		valid = decrement(v);
		if (!valid)
			return ;
		for (int i = 0; i < 26; i++)
		{
			if (orders[i] == 0)
			{	
				push(&LinkList[i]);
				orders[i] = -1;
				break;
			}
		}
	}
	/*for (int i = 0; i < 26; i++)
	{
		if (queue[i])
			cout<<queue[i]->letter;
	}
	cout << "\nihead = "<<ihead<<endl;
	*/	

}

void toposort::print()
{	
	if (!valid)
	{
		cout<<"另请高明吧"<<endl;
		return ;
	}
	for(int i = 0; i < 26; i++)
	{
		if (!queue[i])
		{
			cout<<"另请高明吧"<<endl;
			return ;
		}
	}
	for(int i = 0; i < 26; i++)
	{
		cout<<queue[i]->letter;
	}
	cout<<endl;
}
int main(void)
{	
	toposort s;
	s.sort();
	s.print();
	return 0;
}
