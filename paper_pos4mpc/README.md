- main.tex: 论文的tex源码
- section/: 论文每个章节的tex源码
- img.ipynb: matplot绘图代码
- img/: tex中用到的图片
- mpc.bib: tex中的引用文献

# Review

- [x] 图片 

## Introduction
- [x] MPC introduction\application
- [x] MPC efficiency problem, other's solution
- [x] our method, why important
- [x] out expr result
- [x] contribution
- [x] section layout
## Motivation
- [x] motivation example
- [x] 3  challenges
## Preliminary
- [x] MPC define
- [x] adv
- [x] security model
- [x] syntax
- [x] semantic
- [x] type system
- [x] information leakage
- [x] information leakage lower bound 
## Method
- [x] overview
- [x] optimization
- [x] leakage collection
- [x] verification
- [ ] [**施工中**]proof
## Impl and expr 
- [x] impl
- [x] expr

## Related work

- [x] tradeoff
- [x] circuit level optimization
- [x] specified protocol
- [x] hybrid protocol



- [x] Conclusion



------



# ~~一稿到二稿rewrite: Guideline~~

## ~~Introduction~~

~~intro的逻辑与后面章节的无关~~

- [x] ~~什么是MPC，MPC的定义，自然语言+简洁形式描述；~~
- [x] ~~MPC的作用，应用场景~~
- [x] ~~MPC的问题，效率，为什么是效率，哪些方法解决效率问题~~
- [x] ~~我们的方法解决了哪方面效率问题，什么方法，为什么重要~~
- [x] ~~我们的结果(部分)~~
- [ ] ~~我们的结果~~
- [x] ~~contribution~~
- [x] ~~文章结构~~

## ~~Motivation~~

- [x] ~~challenges凝练到三个，展开讲讲为什么是challenges~~

## ~~Preliminary~~

- [x] ~~MPC~~
  - ~~MPC define, usage; GC,SS,FHE~~
- [x] ~~Adv~~ 
- [x] ~~Security~~
- [x] ~~MPC framework~~
  - ~~MPC lang model-syntax~~
- [x] ~~Information leak~~
  - [x] ~~privacy information~~
  - [x] ~~adversary knowledge define~~
- [x] ~~type system~~
- [x] ~~leakage semantic~~ 
  - [x] ~~leakage define~~
  - [x] ~~semantic with leakage~~
- [x] ~~leakage lower bound~~

~~MPC理论部分的内容参考A Pragmatic Introduction to Secure Multi-Party Computation, Secure Multiparty Computation; MPC框架部分参考Fairplay~~

### ~~Method~~

~~overview+分节讲每一个模块的作用+结合motivation example~~

- [x] ~~overview~~ 
  - [x] ~~我们方法在MPC开发流程中的位置，~~
  - [x] ~~整体的算法~~
- [x] ~~optimization~~
- [x] ~~leakage trace collection~~ 
  - [x] ~~sym收集信息的算法~~
- [x] ~~verification~~
  - [x] ~~leakage secure define~~
  - [x] ~~verify~~

### ~~Impl and expr~~

- [ ] ~~research question~~
  - [x] ~~effeteness？ GC~~ 
  - [x] ~~efficacy? verification time~~
  - [x] ~~universal? SS~~

### ~~Related work~~

~~性能相关的，和introduction中的讨论呼应~~

- [x] ~~circuit level，half gate、free-xor等等~~
- [ ] ~~专用协议，PSI~~
- [ ] ~~混合协议~~
- [x] ~~leak-performance-tradeoff~~

