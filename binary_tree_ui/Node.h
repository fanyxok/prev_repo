#pragma once
#include <graphics.h>
#include <conio.h>
#include "Point.h"
#include <cmath>
#include "Global.h"
#include "TreeBase.h"


// 树节点：矩形
class Node
{
private:
	int date;
	int rank; // rank
public:
	Node* left = NULL;
	Node* right = NULL;

private:
	int centerX = 0;
	int centerY = 0;
	int halfW = 0;
	int halfH = 0;
	wchar_t* text = NULL;							// 控件内容
	void (*userfunc)() = NULL;						// 控件消息

public:
	Node(int d, int r) {
		date = d;
		rank = r;
		left = NULL;
		right = NULL;
	}

	~Node()
	{
		if (text != NULL)
			delete[] text;
	}

	void Create(int x, int y, int w, int h, wchar_t* title, void (*func)())
	{
		text = new wchar_t[wcslen(title) + 1];
		wcscpy_s(text, wcslen(title) + 1, title);
		centerX = x; centerY = y; halfH = h; halfW = w;
		userfunc = func;
		// 绘制用户界面
		Show();
	}


	// 绘制界面
	void Show()
	{
		int oldlinecolor = getlinecolor();
		int oldbkcolor = getbkcolor();
		int oldfillcolor = getfillcolor();

		setlinecolor(BLACK);			// 设置画线颜色
		setbkcolor(LIGHTGRAY);				// 设置背景颜色
		setfillcolor(LIGHTGRAY);			// 设置填充颜色


		fillrectangle(centerX-halfW, centerY+halfH, centerX+halfW, centerY - halfH);
		line(centerX - halfW, centerY, centerX + halfW, centerY);
		line(centerX, centerY, centerX, centerY+halfH);

		wchar_t val_str[256];
		swprintf_s(val_str, L"%d", date);
		outtextxy(centerX-halfW + (2 * halfW - textwidth(val_str) + 1) / 2, centerY - halfH + (halfH - textheight(val_str) + 1) / 2, val_str);

		wchar_t rank_str[256];
		swprintf_s(rank_str, L"%d", rank);
		outtextxy(centerX-halfW + (halfW - textwidth(rank_str) + 1) / 2, centerY + (halfH - textheight(rank_str) + 1) / 2, rank_str);

		outtextxy(centerX + (halfW - textwidth(text) + 1) / 2, centerY + (halfH - textheight(text) + 1) / 2, text);


		setlinecolor(oldlinecolor);
		setbkcolor(oldbkcolor);
		setfillcolor(oldfillcolor);
	}

	Point getTopPoint() {
		return Point{ centerX, centerY - halfH };
	}
	Point getBottomPoint() {
		return Point{ centerX, centerY + halfH };
	}
	
	void OnMessage()
	{
		if (userfunc != NULL)
			userfunc();
	}
};

// y axis of root node
	// x axis of root node
	// half of one node height
	// half of one node width
	// interval height of two near level (centra to centra)
	// interval width of two near node in same level (centra to centra) of the leaf level with max depth

class Tree
{	
public:
	Node* node;
	int rootX;
	int rootY;

	int nodeH;
	int nodeW;

	int intervalH;
	int intervalW;

	wchar_t info[4]= L"ABC";

	int h; // height/depth start from 1

public:
	Tree(Node *n, int x, int y, int H, int W, int interH, int interW) {
		node = n;
		rootX = x;
		rootY = y;
		nodeH = H;
		nodeW = W;
		intervalH = interH;
		intervalW = interW;
		h = -1;
		
	}

	int getHeight(Node* node)
	{
		if (node == NULL)
			return 0;
		else {
			/* compute the height of each subtree */
			int lheight = getHeight(node->left);
			int rheight = getHeight(node->right);

			/* use the larger one */
			if (lheight > rheight) {
				return (lheight + 1);
			}
			else {
				return (rheight + 1);
			}
		}
	}
	void draw() {
		if (h == -1) {
			h = getHeight(node);

		}
		clearrectangle(0, 0, WIDTH - 160 - 1, HEIGHT);
		clearrectangle(WIDTH - 160 - 1, 90 + 1, WIDTH, HEIGHT);
		drawInorder(node, 1, 1);
		
	}
	void drawInorder(Node *cur, int height, int nth)
	{
		if (cur == NULL)
			return;

		/* first recur on left child */
		drawInorder(cur->left, height + 1, nth * 2 );

		/* then print the data of node */
		int Y = rootY + (height - 1) * intervalH;
		// balance
		int num_interval = pow(2, (height - 1)) - 1;
		int interval_width = intervalW * pow(2, h - height);
		int total_width = num_interval * interval_width;
		int nth_interval = nth - pow(2, height - 1);
		int X = rootX - total_width / 2 + nth_interval * interval_width;
		wchar_t val_str[256];
		swprintf_s(val_str, L"%d", height-1);
		if (X >= -intervalW && X <= WIDTH + intervalW
			&& Y >= -intervalH && Y <= HEIGHT + intervalH) {
			cur->Create(X, Y, nodeW, nodeH, val_str, NULL);
			cur->Show();
		}
		Point from = Point{ X, Y + nodeH };
		if (cur->left) {
			Y = rootY + (height)*intervalH;
			num_interval = pow(2, (height)) - 1;
			interval_width = intervalW * pow(2, h - height - 1);
			total_width = num_interval * interval_width;
			nth_interval = nth * 2 - pow(2, height);
			X = rootX - total_width / 2 + nth_interval * interval_width;
			Point to = Point{ X, Y - nodeH };
			if (to.x >= 0 && to.x <= WIDTH 
				&& to.y >= 0 && to.y <= HEIGHT ) {
				line(from.x, from.y, to.x, to.y);
			}
			else if (from.x >= 0 && from.x <= WIDTH
				&& from.y >= 0 && from.y <= HEIGHT) {
				line(from.x, from.y, to.x, to.y);
			}
			else {
				bool topline =    max(from.x, to.x) < 0     || max(from.y, to.y) < 0      || WIDTH < min(from.x, to.x) || 0 < min(from.y, to.y);
				bool bottomline = max(from.x, to.x) < 0     || max(from.y, to.y) < HEIGHT || WIDTH < min(from.x, to.x) || HEIGHT < min(from.y, to.y);
				bool leftline =   max(from.x, to.x) < 0     || max(from.y, to.y) < 0      || 0 < min(from.x, to.x)     || HEIGHT < min(from.y, to.y);
				bool rightline =  max(from.x, to.x) < WIDTH || max(from.y, to.y) < 0      || WIDTH < min(from.x, to.x) || HEIGHT < min(from.y, to.y);
				if (!(topline || bottomline || leftline || rightline)) {
					;
				}
				else {
					line(from.x, from.y, to.x, to.y);
				}

			}
			
		}
		/* now recur on right child */
		drawInorder(cur->right, height + 1, nth * 2 + 1);

		if (cur->right) {
			Y = rootY + (height)*intervalH;
			num_interval = pow(2, (height)) - 1;
			interval_width = intervalW * pow(2, h - height - 1);
			total_width = num_interval * interval_width;
			nth_interval = nth * 2 + 1 - pow(2, height);
			X = rootX - total_width / 2 + nth_interval * interval_width;
			Point to = Point{ X, Y - nodeH };
			if (to.x >= 0 && to.x <= WIDTH
				&& to.y >= 0 && to.y <= HEIGHT) {
				line(from.x, from.y, to.x, to.y);
			}
			else if (from.x >= 0 && from.x <= WIDTH
				&& from.y >= 0 && from.y <= HEIGHT) {
				line(from.x, from.y, to.x, to.y);
			}
			else {
				bool topline = max(from.x, to.x) < 0 || max(from.y, to.y) < 0 || WIDTH < min(from.x, to.x) || 0 < min(from.y, to.y);
				bool bottomline = max(from.x, to.x) < 0 || max(from.y, to.y) < HEIGHT || WIDTH < min(from.x, to.x) || HEIGHT < min(from.y, to.y);
				bool leftline = max(from.x, to.x) < 0 || max(from.y, to.y) < 0 || 0 < min(from.x, to.x) || HEIGHT < min(from.y, to.y);
				bool rightline = max(from.x, to.x) < WIDTH || max(from.y, to.y) < 0 || WIDTH < min(from.x, to.x) || HEIGHT < min(from.y, to.y);
				if (!(topline || bottomline || leftline || rightline)) {
					;
				}
				else {
					line(from.x, from.y, to.x, to.y);
				}

			}
		}
		
	}

	void moveTo(int x, int y) {
		rootX = x;
		rootY = y;
	}
};