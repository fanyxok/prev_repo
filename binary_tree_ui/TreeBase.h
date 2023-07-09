#pragma once
#ifndef TREE_BASE_
#define TREE_BASE_

#include "stdc++.h"
#include "Global.h"

using namespace std;

#define ls son[x][0]
#define rs son[x][1]
#define N 1010100
extern int dep[N];
extern int son[N][2];
extern int Rt;

extern int sz[N];
extern int val[N];
extern int cc;
extern int fa[N];

typedef struct AVL {
	AVL() { dep[0] = -1; }

	void update(int x) {
		sz[x] = sz[ls] + sz[rs] + 1;
		dep[x] = max(dep[ls], dep[rs]) + 1;
	}

	int isson(int x) { return son[fa[x]][1] == x; }

	void rotate(int x) {
		int y = fa[x], z = isson(x);
		fa[x] = fa[y];
		if (fa[y]) son[fa[y]][isson(y)] = x;
		else Rt = x;
		fa[y] = x;
		if (son[x][z ^ 1]) fa[son[x][z ^ 1]] = y;
		son[y][z] = son[x][z ^ 1];
		son[x][z ^ 1] = y;
		update(y);
		update(x);
	}

	void balance(int x) {
		if (dep[ls] - dep[rs] >= 2 || dep[ls] - dep[rs] <= -2) {
			int z = !(dep[ls] - dep[rs] >= 2);
			if (dep[son[son[x][z]][z]] > dep[son[son[x][z]][z ^ 1]]) {
				rotate(son[x][z]);
			}
			else {
				rotate(son[son[x][z]][z ^ 1]);
				rotate(son[x][z]);
			}
		}
	}

	void add(int x, int v) {
		if (v <= val[x]) {
			if (!ls) ls = cc, fa[cc] = x;
			else add(ls, v);
		}
		else {
			if (!rs) rs = cc, fa[cc] = x;
			else add(rs, v);
		}
		update(x);
		balance(x);
	}

	void insert(int v) {//printf("insert %d\n",v);
		val[++cc] = v;
		update(cc);
		if (!Rt) Rt = cc;
		else add(Rt, v);
	}

	int find(int x, int v) {
		if (val[x] == v) return x;
		if (val[x] > v) return find(ls, v);
		else return find(rs, v);
	}

	void remove(int v) {//printf("remove %d\n",v);
		int x = find(Rt, v), y, tmp;
		if (!ls || !rs) {
			y = ls + rs;
			if (y) fa[y] = fa[x];
			if (fa[x]) son[fa[x]][isson(x)] = y;
			else Rt = y;
			x = fa[x];
		}
		else {
			y = ls;
			while (son[y][1]) y = son[y][1];
			son[fa[y]][isson(y)] = son[y][0];
			if (son[y][0]) fa[son[y][0]] = fa[y];
			val[x] = val[y];
			tmp = fa[y];
			fa[y] = 0;
			x = tmp;
		}
		while (x) {//printf("x=%d\n",x);
			update(x);
			balance(x);
			x = fa[x];
		}
	}

	int findkth(int x, int k) {
		if (k <= sz[ls]) return findkth(ls, k);
		if (k == sz[ls] + 1) return val[x];
		return findkth(rs, k - sz[ls] - 1);
	}

	int findrank(int x, int v) {
		if (!x) return 0;
		if (val[x] < v) return findrank(rs, v) + sz[ls] + 1;
		else return findrank(ls, v);
	}

	int kth(int k) { return findkth(Rt, k + 1); }
	int rank(int v) { return findrank(Rt, v); }
	int size() { return sz[Rt]; }

	void test() {}
#undef N
} AVL;
extern AVL avl;

//struct Node_AVL {
//	int element, rank;
//	Node_AVL* left = NULL;
//	Node_AVL* right = NULL;
//	Node_AVL(int d, int r) {
//		element = d;
//		rank = r;
//	}
//	~Node_AVL();
//};
//
//struct Tree_AVL {
//	Node_AVL* root;
//	Tree_AVL() {
//		root = NULL;
//	}
//
//	void Create(int x, Node_AVL* now) {
//		if (x == Rt) {
//			Node_AVL* node = new Node_AVL(x, dep[x]);
//			root = node;
//			now = root;
//		}
//		// int ls = son[x][0];
//		// int rs = son[x][1];
//		if (ls != 0) {
//			Node_AVL* tmp = new Node_AVL(ls, dep[ls]);
//			now->left = tmp;
//			Create(ls, now->left);
//		}
//		if (rs != 0) {
//			Node_AVL* tmp = new Node_AVL(rs, dep[rs]);
//			now->right = tmp;
//			Create(rs, now->right);
//		}
//	}
//
//	void print_tree(Node_AVL* now) {
//		cout << now->element << " ";
//		if (now->left != NULL) {
//			print_tree(now->left);
//		}
//		if (now->right != NULL) {
//			print_tree(now->right);
//		}
//	}
//};
//
//void tree_print(int x) {
//	cout << x << " ";
//	if (ls != 0) {
//		tree_print(ls);
//	}
//	if (rs != 0) {
//		tree_print(rs);
//	}
//}

/* PLEASE DO NOT CHANGE BELOW*/


#endif // !TREE_BASE_
