#pragma once
#include <graphics.h>
#include <conio.h>
#include "EasyButton.h"
#include "Tree.h"
#include <iostream>

int tans = 0;
int sz[1010100];
int cc;
int fa[1010100];

// button
EasyButton startButton;

// GUI Tree
Tree* T;

// declare 
void drawShowWindow();

void calculateTree() {

}

void drawTree() {

}

void On_btnStart_Click() {
	wchar_t line[256];
	bool in = InputBox(line, 256, L"Type your inputs:", L"INPUT", NULL, 0, 200);

	if (!in) {
		MessageBox(GetHWnd(), L"Invalid Input", L"INPUT", MB_OK);
		return;
	}
	if (in) {
		wchar_t* pm;
		wchar_t* pmt;
		wchar_t* pl, *plt;
		for (int i = 0; i < 5; i++) {
			pm = i == 0 ? wcstok_s(line, L"\n", &pmt) : wcstok_s(NULL, L"\n", &pmt);
			if (pm == NULL) {
				MessageBox(GetHWnd(), L"Invalid Input", L"INPUT", MB_OK);
				return;
			}
			int len = i == 0 ? 5 : 4;
			for (int j = 0; j < len; j++) {
				pl = j == 0 ? wcstok_s(pm, L" ", &plt) : wcstok_s(NULL, L" ", &plt);
				if (pl != NULL) {
					if (i == 0) {
						P1[j] = _wtoi(pl);
					}
					else {
						P2[4 * (i-1) + j] = _wtof(pl);
					}
					std::cout << _wtof(pl) << std::endl;
				}
				else {
					MessageBox(GetHWnd(), L"Invalid Input", L"INPUT", MB_OK);
					return;
				}
			}
			
		}
	}
	DynTreeMain();
	// 从code的计算结果中初始化一颗二叉树
	T = Base2GUI();
	BeginBatchDraw();
	T->draw();
	drawShowWindow();
	EndBatchDraw();
}

void On_btnInsert_Click() {
	wchar_t input[256];
	bool in = InputBox(input, 256, L"INSERT", L"INSERT");
	if (in) {

	}
	//MessageBox(GetHWnd(), L"INSERT", L"INSERT", MB_OK);
}

void On_btnDelete_Click() {
	MessageBox(GetHWnd(), L"DELETE", L"DELETE", MB_OK);
}

void drawShowWindow() {
	// draw show window contains tan value, insert button, delete button with text style Consolas
	settextstyle(16, 0, _T("Consolas"));
	int showWindowW = 160;
	int showWindowH = 90;
	// draw the windon with color 0xB0C4DE
	setfillcolor(0xB0C4DE);
	fillrectangle(WIDTH - showWindowW, 0, WIDTH, showWindowH);
	// draw insert button 
	wchar_t insert_text[15] = L"CLICK TO INPUT";
	startButton.Create(WIDTH - showWindowW * 0.9, showWindowH / 2 * 1.1, WIDTH - showWindowW * 0.1, showWindowH * 0.9, insert_text, On_btnStart_Click);
	// draw delete button
	wchar_t delete_text[7] = L"DELETE";
	//deleteButton.Create(WIDTH - showWindowW / 2 * 0.9, showWindowH / 2 * 1.1, WIDTH - showWindowW * 0.1, showWindowH * 0.9, delete_text, On_btnDelete_Click);
	// draw tan value
	wchar_t tans_text[256];
	swprintf_s(tans_text, L"tans: %d", tans);
	RECT r = { WIDTH - showWindowW , 0, WIDTH - 1, showWindowH / 2 - 1 };
	drawtext(tans_text, &r, DT_CENTER | DT_VCENTER | DT_SINGLELINE);
}

void On_btnOk_Click()
{
	MessageBox(GetHWnd(),L"II", L"II", MB_OK);
}



void connectNode(Node& a, Node& b) {
	Point from = a.getBottomPoint();
	Point to = b.getTopPoint();
	line(from.x, from.y, to.x, to.y);
}

EasyButton btnRank;

int main()
{	

	// 初始化绘图窗口
	initgraph(WIDTH, HEIGHT, EW_SHOWCONSOLE);
	// 获得窗口句柄
	HWND hWnd = GetHWnd();
	// 使用 Windows API 修改窗口名称
	SetWindowText(hWnd, L"Dynamic Tree");
	// 初始化， 背景色, 文字颜色
	setbkcolor(0xeeeeee);
	cleardevice();
	settextcolor(BLACK);
	setlinestyle(PS_SOLID, 1);
	setlinecolor(BLACK);

	// draw show window
	drawShowWindow();

	ExMessage msg;
	bool LB_DOWN = false;
	Point last;
	while (true)
	{
		msg = getmessage(EM_MOUSE);			// 获取消息输入

		if (msg.message == WM_LBUTTONDOWN)
		{	
			LB_DOWN = true;
			last.x = msg.x; last.y = msg.y;
			// 判断控件
			if (startButton.Check(msg.x, msg.y)) {
				startButton.OnMessage();
			}

		}
		// 鼠标拖动画布
		else if (msg.message == WM_MOUSEMOVE) {
			if (LB_DOWN) {
				T->moveTo(T->rootX + (msg.x - last.x), T->rootY + (msg.y - last.y));
				BeginBatchDraw();
				T->draw();
				drawShowWindow();
				FlushBatchDraw();
				EndBatchDraw();
				last.x = msg.x; last.y = msg.y;
			}
		}
		else if (msg.message == WM_LBUTTONUP) {
			LB_DOWN = false;
		}

	}

	// 按任意键退出
	_getch();

	closegraph();
	DeleteTree(T);
	return 0;
}