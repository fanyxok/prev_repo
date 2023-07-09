import copy
def ism(A):
    
        
    if len(A) is 0:
        return A
    else:
        return True

def numtype(s):
    if 'j' in s:
        return complex
    elif '.' in s:
        return float
    else:
        return int
    
def IsStandard(s):
    
    if ism(s) is True:
        if len(s) <= 1:
            return True
        elif type(s[1]) == dict:
            return False
        else:
            return True
    else:
        return None
        
def Str2Mat(A):
    H=copy.deepcopy(A)
    if H[0] is '[' :
        I=H.count(';')
        s=H.replace(';',',')
        s=eval(s)
        n=int(len(s)/(I+1))
        D=list()
        for i in range(I+1):
            D.append(s[(i*n):(n*(i+1))])
        return D
    else:
        b=eval(A[0])
        c=eval(A[2])
        a=A[3:]
        a=eval(a)
        D={}
        for i in a:
            x=i[:2]
            
            y=i[2]
            S=numtype(str(y))
            s=S(y)
            D[x]=s
        L=[[b,c],D]
        return L 

def Mat2StrStandard(A):
    s=copy.deepcopy(A)
    if IsStandard(s) is True:
        a=len(s)  
        l=list()
        for i in s:
            i=str(i)
            c=i[1:-1]
            l.append(c)    
        b=';'.join(l)
        if 'j' in b:
            b=b.replace('(','')
            b=b.replace(')','')
        d='['+b+']'
        drk=''.join(d.split()) 
        return drk
    elif ism(s) is s:
        return '[]'
    else:
        
        c=Sparse2Standard(s)
        p=Mat2StrStandard(c)
        return p
def Mat2StrSparse(A):
    s=copy.deepcopy(A)
    if len(s) == 0:
        s='[[0,0],{}]'
        return s
    elif IsStandard(s) is False:
        b=s[0][0]
        c=s[0][1]
        s1='{0}-{1}'.format(b,c)
        d=s[1]
        L=list()
        for m,n in d.items():
            m=str(m)
            n=str(n)
            m=m.replace(')',', ')
            if 'j' in n:
                n=n.lstrip('(')
                n=n.rstrip(')')
            a=m+n+')'
            L.insert(-1,a)
            
        L.sort(key=str)
        
        L=','.join(L) 
        L=str(L)
        
        
        
        P=s1+'{'+L+'}'
        Prk=''.join(P.split())
        return Prk
    else:
        c=Standard2Sparse(s)
        
        p=Mat2StrSparse(c)
        return p

def Standard2Sparse(A):
    s=copy.deepcopy(A)
    if IsStandard(s) is True:
        b=len(s)
        c=len(s[0])
        i=q=0
        D=dict()
        for n in s:
            q+=1
            i=0
            for m in n:
                i=i+1
                if m is not 0:
                    x=(q,i)
                    D[x]=m
                
                    
        L=[[b,c],D]
        return L 
    elif ism(s) is s:
        s=[[0,0],{}]
        return s
    else:
        return None

def Sparse2Standard(A):
    s=copy.deepcopy(A)
    if IsStandard(s) is True:
        return s
    elif IsStandard(s) is None:
        s=Standard2Sparse(s)
        return s
    b=s[0][0]
    c=s[0][1]    
    L=list()        
    for i in range(b):
        l=[]
        for j in range(c):
            if (i+1,j+1) in s[1].keys():
                l.append(s[1][(i+1,j+1)])
            else:
                l.append(0)
        L.append(l)
    return L
    
def MatAdd(A,B):
    s=copy.deepcopy(A)
    p=copy.deepcopy(B)
    if IsStandard(s) is True:
        if IsStandard(p) is True:
            D=list()
            l=list()
            for i in range(len(s)):
                for n in range(len(s[i])):
                    l.append(s[i][n]+p[i][n])
            for i in range(len(s)):
                x=i*len(s[i])
                y=(i+1)*len(s[i])
                D.append(l[x:y])
            return D
        else:
            p=Sparse2Standard(p)
            S=MatAdd(s,p)
            return S
    else:
        s=Sparse2Standard(s)
        S=MatAdd(s,p)
        return S

def MatSub(A,B):
    s=copy.deepcopy(A)
    p=copy.deepcopy(B)
    if IsStandard(s) is True:
        if IsStandard(p) is True:
            D=list()
            l=list()
            for i in range(len(s)):
                for n in range(len(s[i])):
                    l.append(s[i][n]-p[i][n])
            for i in range(len(s)):
                x=i*len(s[i])
                y=(i+1)*len(s[i])
                D.append(l[x:y])
            return D
        else:
            p=Sparse2Standard(p)
            S=MatSub(s,p)
            return S
    elif len(s) ==0:
        return s
    else:
        s=Sparse2Standard(s)
        S=MatSub(s,p)
        return S

def MatScalarMul(A,p):
    s=copy.deepcopy(A)
    if IsStandard(s) is True:
        for m in s:
            i=0
            for n in m:
                c=n*p
                m[i]=c
                i+=1
        return s
    elif ism(s) is None:
        s=Standard2Sparse(s)
        return s
    else:
        s=Sparse2Standard(s)
        P=MatScalarMul(s,p)
        return P

def MatTransposition(A):
    s=copy.deepcopy(A)
    if IsStandard(s) is True:
        b=len(s[0])
        c=len(s)
        l=list()
        for x in range(b):
            for y in s:
                l.append(y[x])
        D=list()
        for i in range(c+1):
            D.append(l[(i*c):(c*(i+1))])
        return D
    elif IsStandard(s) is False:
        p=Sparse2Standard(s)
        b=MatTransposition(p)
        c=Sparse2Standard(b)
        return c
    else:
        s=Standard2Sparse(s)
        return s

def MatEq(s,q):
    if s==[]:
        if q==[]:
            return True
        elif q==[[0,0],{}]:
            return True
   
        
    elif IsStandard(s) is True:
        if IsStandard(q) is True:
            if s==q:
                return True
            else:
                return False
        else:
            q=Sparse2Standard(q)
            if s==q:
                return True
            else:
                return False

    else:
        if IsStandard(q) is False:
            if s==q:
                return True
            else:
                return False
        else:
            q=Standard2Sparse(q)
            if s==q:
                return True
            else:
                return False


            
                
