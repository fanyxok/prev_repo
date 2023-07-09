/* You may only include these two header files. */
#include <utility>      /* For std::pair. */
#include <functional>   /* For std::less. */


/* You may finally allowed to use // in C++. */

#define RAND_HALF ((RAND_MAX) / 2)


//
//
// The implementation for skiplist class follows here!
//
//
//                S K I P L I S T
//
//

/* ==================== Begin students' code ==================== */

/* One possible function definition has been filled for you,
 * you may change it as you wish. */
template<class Key, class Val, class Compare>
inline int
skiplist<Key, Val, Compare>::size() const
{
    return listsize;
}
template<class Key, class Val, class Compare>
inline int
skiplist<Key, Val, Compare>::level() const
{
    return list_level;
}
template<class Key, class Val, class Compare>
inline bool
skiplist<Key, Val, Compare>::empty() const
{
    return listsize==0;
}
template<class Key, class Val, class Compare>
inline int 
skiplist<Key, Val, Compare>::limit()
{
    return height_limit;
}

template<class Key, class Val, class Compare>
skiplist<Key, Val, Compare>::skiplist(int h)
{
    header_forward = new _listnode*[h+1]();
    list_level = 1;
    height_limit = h;
    listsize = 0;
    return ;
}

template<class Key, class Val, class Compare>
skiplist<Key, Val, Compare>::~skiplist()
{
    _listnode *x;
    x = header_forward[0];
    while (x != NULL){
        _listnode* curr = x;      
        x = x->forward[0];
        delete curr;
    }
    delete [] header_forward;
}

template<class Key, class Val, class Compare>
std::pair<typename skiplist<Key, Val, Compare>::iterator, bool>
skiplist<Key, Val, Compare>::insert(const Key &key, const Val &val)
{
    // Your code here ...
    iterator exist = find(key);
    if (exist == end())
    {
        _listnode *x;
        int node_level = 0;
        while (node_level < skiplist::height_limit - 1 && rand() < RAND_HALF)
        {   ++node_level;}      
        if (node_level+1 > (list_level))
        {   list_level = node_level+1;}    
        _listnode **wtf = header_forward;
        x = new _listnode(node_level, key, val);  
        for (int i = node_level; i >=0; i--)
        {
            while(wtf[i]){
                if(cmp(wtf[i]->_data.first,key)==1){
                    wtf = wtf[i]->forward;
                }else{
                    break;}
            }
            x->forward[i] = wtf[i];
            wtf[i] = x;
        }
        listsize++;
        return std::make_pair(iterator(x), true);  
    }else{
        exist->second = val;
        return std::make_pair(iterator(exist), false);
    }
    
}



template<class Key, class Val, class Compare>
bool
skiplist<Key, Val, Compare>::erase(const Key &key)
{
    if (find(key) != end())
    {
        _listnode **forw = header_forward;
        _listnode *move;
        for (int i = level()-1; i >=0; i--)
        {
            while(forw[i])
            {
                if (cmp(forw[i]->_data.first, key) == 1)
                {   forw = forw[i]->forward;}
                else if (cmp(key, forw[i]->_data.first) == 1)
                {   break;}
                else{
                    move = forw[i];
                    forw[i] = forw[i]->forward[i];
                    break;
                }
            }
        }
        listsize--;
        delete move;
        return true;
    }else{
        return false;
    }
}

template<class Key, class Val, class Compare>
typename skiplist<Key, Val, Compare>::iterator
skiplist<Key, Val, Compare>::find(const Key &key)
{
    _listnode **ford;
    ford = header_forward;
    for (int i = height_limit-2; i >= 0; i--)
    {
        while (ford[i]!=NULL)
        {
            if (cmp(ford[i]->_data.first, key) ==1)
            {
                ford = ford[i]->forward;
            }else if (cmp(key, ford[i]->_data.first) ==1)
            {
                break;
            }else{
                return iterator(ford[i]);
            }
        }
        //ford = header_forward;
    }
    return end();
}

template<class Key, class Val, class Compare>
typename skiplist<Key, Val, Compare>::const_iterator
skiplist<Key, Val, Compare>::find(const Key &key) const {
   _listnode **ford;
    ford = header_forward;
    for (int i = height_limit-2; i >= 0; i--)
    {
        while (ford[i]!=NULL)
        {
            if (cmp(ford[i]->_data.first, key) ==1)
            {
                ford = ford[i]->forward;
            }else if (cmp(key, ford[i]->_data.first) ==1)
            {
                break;
            }else{
                return const_iterator(ford[i]);
            }
        }
        //ford = header_forward;
    }
    return end();
}
template<class Key, class Val, class Compare>
inline typename skiplist<Key, Val, Compare>::iterator
skiplist<Key, Val, Compare>::begin()
{
    return iterator(header_forward[0]);
}

template<class Key, class Val, class Compare>
inline typename skiplist<Key, Val, Compare>::iterator
skiplist<Key, Val, Compare>::end()
{
    return iterator();
}

template<class Key, class Val, class Compare>
inline typename skiplist<Key, Val, Compare>::const_iterator
skiplist<Key, Val, Compare>::begin() const
{
    return const_iterator(header_forward[0]);
}

template<class Key, class Val, class Compare>
inline typename skiplist<Key, Val, Compare>::const_iterator
skiplist<Key, Val, Compare>::end() const
{
    return const_iterator();
}
/*
template<class Key, class Val, class Compare>
void
skiplist<Key, Val, Compare>::print()
{
    
    
    for (iterator i = begin(); i != end(); ++i)
    {
        std::cout<<"++"<<std::endl;
        std::cout<<i->first<<","<<i->second<<std::endl;
    }
}*/

/* ==================== End students' code ==================== */



//
//
// The implementation for _listnode subclass follows here!
//
//
//                L I S T   N O D E
//
//

/* ==================== Begin students' code ==================== */
/* One possible function definition has been filled for you,
 * you may change it as you wish. */
template<class Key, class Val, class Compare>
skiplist<Key, Val, Compare>::_listnode::_listnode( int _level, const Key &key, const Val &val):_data(std::make_pair(key, val))
{
    _listnode::forward = new _listnode*[_level+1];
}

template<class Key, class Val, class Compare>
skiplist<Key, Val, Compare>::_listnode::~_listnode()
{
    delete [] forward;
}


/* ==================== End students' code ==================== */



//
//
// The implementation for iterator subclass follows here!
//
//
//                I T E R A T O R
//
//

/* ==================== Begin students' code ==================== */

/* One possible function definition has been filled for you,
 * you may change it as you wish. */

template<class Key, class Val, class Compare>
inline skiplist<Key, Val, Compare>::iterator::iterator(_listnode *ptr)
{
    curr_ptr = ptr;
}

template<class Key, class Val, class Compare>
inline skiplist<Key, Val, Compare>::iterator::~iterator()
{
    ;
}

template<class Key, class Val, class Compare>
inline typename skiplist<Key, Val, Compare>::iterator &
skiplist<Key, Val, Compare>::iterator::operator++()
{
    curr_ptr = curr_ptr->forward[0];
    return *this;
}

template<class Key, class Val, class Compare>
inline typename skiplist<Key, Val, Compare>::iterator 
skiplist<Key, Val, Compare>::iterator::operator++(int)
{
    iterator temp = *this;
    ++(*this);
    return (iterator)temp;
}


template<class Key, class Val, class Compare>
inline bool 
skiplist<Key, Val, Compare>::iterator::operator==(const iterator &other) const
{
    return curr_ptr == other.curr_ptr;
}

template<class Key, class Val, class Compare>
inline bool 
skiplist<Key, Val, Compare>::iterator::operator!=(const iterator &other) const
{
    return curr_ptr != other.curr_ptr;
}


template<class Key, class Val, class Compare>
inline std::pair<const Key, Val> * 
skiplist<Key, Val, Compare>::iterator::operator->() const
{
    return &(curr_ptr->_data);
}
/* One possible function definition has been filled for you,
 * you may change it as you wish. */

template<class Key, class Val, class Compare>
inline std::pair<const Key, Val> &
skiplist<Key, Val, Compare>::iterator::operator*() const
{ 
    return (curr_ptr->_data);
}

/* ==================== End students' code ==================== */


//
//
// The implementation for iterator subclass follows here!
//
//
//                C O N S T      I T E R A T O R
//
//
/* ==================== Begin students' code ==================== */
/* Fill in the const_iterator part, this should be similiar to iterator */

template<class Key, class Val, class Compare>
inline skiplist<Key, Val, Compare>::const_iterator::const_iterator(_listnode *ptr)
{
    curr_ptr = ptr;
}

template<class Key, class Val, class Compare>
inline skiplist<Key, Val, Compare>::const_iterator::~const_iterator()
{
    ;
}

template<class Key, class Val, class Compare>
inline typename skiplist<Key, Val, Compare>::const_iterator &
skiplist<Key, Val, Compare>::const_iterator::operator++()
{
    curr_ptr = curr_ptr->forward[0];
    return *this;
}

template<class Key, class Val, class Compare>
inline typename skiplist<Key, Val, Compare>::const_iterator 
skiplist<Key, Val, Compare>::const_iterator::operator++(int)
{
    const_iterator temp = *this;
    ++(*this);
    return (const_iterator)temp;
}

template<class Key, class Val, class Compare>
inline bool 
skiplist<Key, Val, Compare>::const_iterator::operator==(const const_iterator &other) const
{

    return curr_ptr == other.curr_ptr;
}

template<class Key, class Val, class Compare>
inline bool 
skiplist<Key, Val, Compare>::const_iterator::operator!=(const const_iterator &other) const
{

    return curr_ptr != other.curr_ptr;
}


template<class Key, class Val, class Compare>
inline std::pair<const Key, Val> * 
skiplist<Key, Val, Compare>::const_iterator::operator->() const
{
    return &(curr_ptr->_data);
}
/* One possible function definition has been filled for you,
 * you may change it as you wish. */

template<class Key, class Val, class Compare>
inline std::pair<const Key, Val> &
skiplist<Key, Val, Compare>::const_iterator::operator*() const
{ 
    return (curr_ptr->_data);
}




/* ==================== End students' code ==================== */

