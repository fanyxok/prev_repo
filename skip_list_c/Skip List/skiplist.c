    #include <stdio.h>
    #include <stdlib.h>
    #include "skiplist.h"

    /* make  The "cmp" function for comparing integer values. */
    int compareInt(const void *a, const void *b){
        int A;
        int B;
        /* detecte that input a valid*/
        if (a != NULL && b != NULL){
            /* let A B be the int value of the pointer*/
            A = *(int *)a;
            B = *(int *)b;
        }else{;}
        /* compare and return*/
        if (A > B) return 1;
        else if (A < B)  return -1;
        else return 0;   
    }
    /* The "del" function for freeing memory allocated for integers.*/
    void deleteInt(void *a){
        /* free it and let it be nil*/
        free(a);
        a = NULL;
    }
    /* The "alloc" function for allocating memory for integers .*/
    void *allocInt(const void *a){
        int *m;
        m = NULL;
        m = (int *)malloc(sizeof(int *));
        *m = *(int *)a;
        return m;
    }
    /*
    void display(list s){
        int i;
        node n;
        printf("\n******Skip List******\n");
        for ( i=0; i < s->level ;i++)
        {
            n = s->header->forward[i];
            printf( "Level %d \n", i);
            while ( n!= NULL)
            {   
                if (n->key != NULL){
                printf("%d\t", *(int *)(n->key));
                }
                n = n->forward[i];
            }
            printf("\n");
        }

    }*/

    /*A function to create an empty tree. A dummy node is created as the header.*/
    list createSkipList(int height_limit,
                             int (*cmp)(const void *a, const void *b),
                             void *(*alloc)(const void *),
                             void (*del)(void *)){
        /*create skiplist s */
        list s;
        int i;
        /* allocate memory for s and header*/
        s = (list)malloc(sizeof(list)+sizeof(node)+sizeof(void *)*5);
        s->header = (node)malloc(sizeof(node)+sizeof(node *)*(height_limit+1));
        s->header->val = NULL;
        s->header->key = NULL;
        /*allocate memory for header->forward */
        s->header->forward = (node *)malloc(sizeof(node *) * (height_limit));
        /* let forward[i] all be NULL*/
        for (i = 0; i < height_limit; i++){
            s->header->forward[i] = NULL;
        }
        /* init the elements of s*/
        s->level = 1;
        s->height_limit = height_limit;
        /* pass three function to s*/
        s->cmp = cmp;
        s->alloc = alloc;
        s->del = del;
        return s;
                             }
    /* check if a list s is empty*/
    int isListEmpty(list s){
        if (s->header->forward[0] == NULL &&
            s->level == 1) return 1;
        else return 0;
    }
    /* A function to insert a new node with key & value determined by "key" into a list.*/
    int insertNode(list s, void *key, void *val){
        /* left is a array with node contains the left node of new node*/
        node *left;
        int i;
        const int lvl = s->height_limit;
        /*x is the new node
        p is a template node varible*/
        node p, x ;
        int node_level;
        if (key == NULL) {return -1;}
        #define RAND_HALF (RAND_MAX)/2
            node_level = 0;
            while (node_level < s->height_limit - 1 && rand() < RAND_HALF) {
                ++node_level;
            }   
        /*if list s is empty, insert x and return*/
        if (isListEmpty(s) == 1){
            x = (node)malloc(sizeof(node)+sizeof(node *)*(lvl +1));
            x->val = s->alloc(val);
            x->key = s->alloc(key);
            /* forward need to have a memory*/
            x->forward = (node *)malloc(sizeof(node *)*lvl);
            /* init the pointer to NULL*/
            for (i = 0; i < lvl; i++){
            x->forward[i] = NULL;
            }
            /* linked s with x*/
            s->header->forward[0] = x;
            return 0;
        }
        /* left is a pointer array*/  
        left = (node *)malloc((lvl) * sizeof(node *));
        /* init is pointer element to NULL*/
        for (i = 0; i < lvl; i++){
            left[i] = NULL;
        }
        /*p is a temp to access the left side node of new node*/
        p = s->header;
        /*
        from level s->level down to level 1 to find the left node p of x in each level
        */
        for (i = s->level-1  ; i >= 0; i--){             
            /*
            find the p of the new key from high level to low level until level 0           
            */
            p = s->header;
            while (p->forward[i] != NULL &&
                s->cmp(p->forward[i]->key, key) == -1){
                p = p->forward[i];
            }
            /* p is a left node of x in i th level*/
            left[i] = p;   
                
        }
        /*p be the right side key to new key in level 0 or NULL   
        */
        p = p->forward[0];
        /* old key = new key */
        if (p != NULL &&
            s->cmp(p->key, key) == 0)
            {
            s->del(p->val);
            p->val = s->alloc(val);
            free(left);
            left=NULL;
            return 1;}
        /*a differ new key */
        else{
                /* define a value as required*/
                   
            if (node_level > (s->level-1)){
                for (i = s->level; i < node_level+1; i++ ){
                    left[i] = s->header;
                }
                /* update s->level*/
                s->level = node_level+1;
            }
            /* create new node x*/
            x = (node)malloc(sizeof(node)+sizeof(node *)*lvl);
            x->val = s->alloc(val);
            x->key = s->alloc(key);
            x->forward = (node *)malloc(lvl * sizeof(node *));  
            /* init pointer to NULL*/
            for (i = 0; i < lvl; i++){
            x->forward[i] = NULL;
            }
            
            /* connect x and the right and the left node of x*/    
            for (i = 0; i <= node_level; i++){
                x->forward[i]= left[i]->forward[i];
                left[i]->forward[i] = x;
            }
            /* if x has a high level ensure it  */
        }
        /* free left and ensure it is NULL*/
        free(left);
        left=NULL;
        return 0;
    }
    /* A function to search for the value corresponding to the key.*/
    void *searchNode(list s, void *key){
        /* variable declear*/
        int i;
        node x;
        void *r;
        x = s->header;
        if (key == NULL){
            return NULL;
        }
        /* from high level to low level to find the left node of key*/
        for (i = s->level-1; i >= 0; i--){
            while (x->forward[i]!= NULL &&
                s->cmp(x->forward[i]->key, key) == -1)
                x = x->forward[i];
        }
        /*let x = the possible node with this key*/
        x = x->forward[0];
        if (x == NULL) {return NULL;}
        /* compare and return */
        if (s->cmp(x->key, key)==0){
            r = x->val;
            return r;
        }
        else {return NULL;}
    }
    /* A function to remove a node with value given by "key" from a list.*/
    int deleteNode(list s, void *key){
        /*declare variable
        left is a array contain the left node of deletenode*/
        node *left;
        node x;
        int i;
        /*when the key is not exaist return a failrue*/
        if (searchNode(s, key) == NULL) return -1;
        /*give memory to left*/
        left = (node *)malloc(sizeof(node *) * (s->height_limit));
        for (i = 0; i< s->height_limit; i++){
            left[i] = NULL;
        }
        x = s->header;
        /*from high level to low level to find the left node of this key node in each level*/
        for (i = s->level-1; i >=0; i--){
            x = s->header;
            while (x->forward[i] != NULL && s->cmp(x->forward[i]->key, key) == -1){
                x = x->forward[i];
            }
            left[i] = x;
            
        }
        /* let x equal the key node*/
        x = x->forward[0];
        /* connect the left and the right node of the key node*/
        for (i=0; i < s->level; i++){
            if (left[i]->forward[i] != x) break;
            left[i]->forward[i] = x->forward[i];
        }
        /*let x go*/
        s->del(x->val);
        s->del(x->key);
        free(x->forward);
        x->forward = NULL;
        
        free(x);
        x=NULL;
        /*if s has empty level, delete the level*/
        /*while (s->level > 1 &&
            s->header->forward[s->level-1] ==NULL)
            s->level--;
        */
        /* free left and ensure it NULL*/
        free(left);
        left=NULL;
        return 0;
    }
    /* A function to free all the dynamically allocated memory used by the tree.*/
    void freeSkipList(list s){
        /* some variable*/
        node x, t;
        x = s->header;
        t = NULL;
        /*if s is empty free it directly*/
        if (isListEmpty(s) == 1){
            free(s->header->forward);
            free(s->header);
            free(s);
            return ;
        }
        /* find the rightest node and delete it */
        while (t != x ){
            while (x->forward[0] != NULL){
                t = x;
                x = x->forward[0];
            }       
            deleteNode(s,x->key);     
            /* let x = header and start from the header again*/
            x = s->header;
        }
        /* x is ues out*/
        x=NULL;
        /* node a all freed at now*/
        /* free the struct of list s*/
        free(s->header->forward);  
        /* free header before free s*/
        free(s->header);
        free(s);
    }




