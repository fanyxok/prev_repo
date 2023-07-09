#include <string>
#include <iostream>
using namespace std;

template <typename Type>
class Stack {
	private:
		int stack_size; //number of objects in the stack
		int array_capacity; //capacity of the array
		Type *array;
	public:
		Stack();
		~Stack();
		bool empty() const;
		Type top() const;
		void push( Type  elem );
		Type pop();
		};

template <typename Type>
Stack<Type>::Stack(void){
    stack_size = 0;
    array_capacity = 500;
    array = new Type[500];
}

template <typename Type>
Stack<Type>::~Stack(void){
    delete[]array;
}

template <typename Type>
bool Stack<Type>::empty(void)const{
    if ( stack_size == 0)
        return true;
    else
        return false;
}

template <typename Type>
Type Stack<Type>::top(void)const{
    if (stack_size == 0){
        return "";
    }
    else
        return array[stack_size-1];
}
template <typename Type>
void Stack<Type>::push(Type elem){
    if (stack_size <= array_capacity){
        array[stack_size] = elem;
        stack_size++;
    }
    else{
        cout <<"push error"<<endl;
    }
}
template <typename Type>
Type Stack<Type>::pop(void){
    if (empty() == true)
        return "pop error";
    else{
        stack_size -= 1;
        return array[stack_size+1];
    }
}

int main(void){
    Stack<string> *s = new Stack<string>();

    Stack<string> *l = new Stack<string>();
/*
    cout << s->empty()<<endl;
    s->push(h);
    cout << s->top()<<endl;
    s->pop();
    cout<<s->empty()<<endl;
    cout << s->stack_size << endl;
*/
    char c[256];
    string ip;
    getline(cin, ip);
    string t = "";
    string m;
    double v1;
    double v2;
    double v3;
    //from 1th char to last char of input
    for (int i = 0; i < ip.size(); i++){
        //If the token is an operator +, or -, push it on the opstack. 
        //However, first remove any operators already on the opstack 
        //and append them to the output list.
        if (ip[i] == '+' || ip[i] == '-'){
            if (t != ""){
                l->push(t);
                t = "";
            }
            while (s->top() == "*" || s->top() == "/" ||
                s->top() == "+" || s->top() == "-"){
                m = s->top();
                s->pop();
                l->push(m);
            }
            m = ip[i];
            s->push(m);
        }
        else if (ip[i] == '*' || ip[i] == '/'){
            //cout<< "deal *" <<endl;
            if (t != ""){
                l->push(t);
                t = "";
                //cout << "number in"<< endl;
            }
            while (s->empty() == false &&
                (s->top() == "*" || s->top() == "/")){
                m = s->top();
                s->pop();
                l->push(m);
            }
            m = ip[i];
            s->push(m);
        }
        else if (ip[i] == '('){
            m = ip[i];
            s->push(m);
        }
        else if (ip[i] == ')'){
            if (t != ""){
                l->push(t);
                t = "";
            }
            while (s->top() != "("){
                m = s->top();
                s->pop();
                l->push(m);
            }
            s->pop();
        }
        else if (ip[i] == '\t' || ip[i] == ' ' || ip[i] == '\n'){
            ;
        }
        else if (ip[i] == '0' || ip[i] == '1' || ip[i] == '2'||
                ip[i] == '3' || ip[i] == '4' || ip[i] == '5' ||
                ip[i] == '6' || ip[i] == '7' || ip[i] == '8' ||
                ip[i] == '9' || ip[i] == '.'){
            t = t + ip[i];
            //cout <<"t"<< t<<endl;
        }
    }
    if (t != ""){
        l->push(t);
        t = "";
        
    }
    while (s->empty() == false){
        m = s->top();
        s->pop();
        l->push(m);
    }
    while(l->empty() == false){
        m = l->top();
        l->pop();
        s->push(m);
        
    }
    while (s->empty() == false){
        if (s->top() == "+"){
            m = l->top();
            l->pop();
            t = l->top();
            l->pop();
            v1 = stod(m);
            v2 = stod(t);
            v3 = v1 + v2;
            sprintf(c,"%.15lf",v3);

            m = c;
            l->push(m);
            s->pop();
        }
        else if (s->top() == "-"){
            m = l->top();
            l->pop();
            t = l->top();
            l->pop();
            v1 = stod(m);
            v2 = stod(t);
            v3 = v2 - v1;
            sprintf(c,"%.15lf",v3);
            m = c;
            l->push(m);
            s->pop();

        }
        else if (s->top()== "*"){
            m = l->top();        
            l->pop();
            t = l->top();
            l->pop();
            v1 = stod(m);
            v2 = stod(t);
            v3 = v2 * v1;
            sprintf(c,"%.15lf",v3);
            m = c;
            l->push(m);
            s->pop();
        }
        else if (s->top() == "/"){
            m = l->top();
            l->pop();
            t = l->top();
            l->pop();
            v1 = stod(m);
            v2 = stod(t);
            v3 = v2 / v1;
            sprintf(c,"%.15lf",v3);
            m = c;
            l->push(m);
            s->pop();
        }
        else{
            m = s->top();
            s->pop();
            l->push(m);
        }    
    }
    m = l->top();
    l->pop();
    v3 = stod(m);
    printf("%.15lf\n",v3);
    //

}
