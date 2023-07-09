#pragma once
//#include "TreeBase.h"
#include "Node.h"
int son[1010100][2];
int dep[1010100];
int val[1010100];
int Rt;


AVL avl;

int A, B, C, lfsr;
double P[4][4];
int lfsr_generator() {
	auto ret = lfsr;
	return (lfsr ^= lfsr << 13, lfsr ^= lfsr >> 17, lfsr ^= lfsr << 5, ret);
}
int state = 0;
tuple<int, int> command(int &state) {
	auto imm = lfsr_generator();
	auto p = double(lfsr_generator() & 0x7fffffff) / INT32_MAX;
	for (int i = 0; i < 4; i++)
		if ((p -= P[state % 4][i]) < 0) {
			state += 4 - state % 4 + i;
			break;
		}
	return tuple<int, int>(state % 4, (imm * A + state * B + C) & 0x7fffffff);
}
/* PLEASE DO NOT CHANGE ABOVE*/
void DynTreeMain() {

	// clean for multi time use
	for (int i = 0; i < 1010100; i++) {
		sz[i] = 0;
		val[i] = 0;
		son[i][0] = 0;  son[i][1] = 0;
		Rt = 0;
		cc = 0;
		dep[i] = 0;
		fa[i] = 0;
	}
	state = 0;
	dep[0] = -1;
	tans = 0;
	int m = P1[0]; lfsr = P1[1]; A = P1[2]; B = P1[3]; C = P1[4];

	for (int i = 0; i < 4; i++)
		for (int j = 0; j < 4; j++) P[i][j] = P2[i * 4 + j];
	for (int i = 1; i <= m; i++) {
		int op, imm;
		tie(op, imm) = command(state);
		if (op == 0) avl.insert(imm);
		if (op == 1) avl.remove(avl.kth(imm % avl.size()));
		if (op == 2) tans ^= avl.rank(imm);
		if (op == 3) tans ^= avl.kth(imm % avl.size());
		/*Tree_AVL node_tree;
		Node_AVL* now;
		node_tree.Create(Rt, now);
		cout << "node tree" << endl;
		node_tree.print_tree(node_tree.root);
		cout << endl << "array tree" << endl;
		tree_print(Rt);
		cout << endl;*/
	}
	//cout << tans << "\n";
}

void Create(Node* t, int x) {
	int l = son[x][0];
	int r = son[x][1];
	if (l != 0) {
		t->left = new Node(val[l], avl.rank(val[l]));
		Create(t->left, l);
	}
	if (r != 0) {
		t->right = new Node(val[r], avl.rank(val[r]));
		Create(t->right, r);
	}
}

Tree *Base2GUI() {
	// crete root
	//avl.rank()
	Node* root = new Node(val[Rt], avl.rank(val[Rt]));
	Create(root, Rt);
	Tree* t = new Tree(root, WIDTH / 2, 50, 20, 50, 55, 105);
	return t;
}

void DeleteNode(Node* n) {
	if (n != NULL) {
		DeleteNode(n->left);
		DeleteNode(n->right);
		delete n;
	}
}

void DeleteTree(Tree *t){
	if (t->node != NULL) {
		DeleteNode(t->node);
	}
	delete t;
}

//void Create(int x, Node_AVL* now) {
//	if (x == Rt) {
//		Node_AVL* node = new Node_AVL(x, dep[x]);
//		root = node;
//		now = root;
//	}
//	// int ls = son[x][0];
//	// int rs = son[x][1];
//	if (ls != 0) {
//		Node_AVL* tmp = new Node_AVL(ls, dep[ls]);
//		now->left = tmp;
//		Create(ls, now->left);
//	}
//	if (rs != 0) {
//		Node_AVL* tmp = new Node_AVL(rs, dep[rs]);
//		now->right = tmp;
//		Create(rs, now->right);
//	}
//}