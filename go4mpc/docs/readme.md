# 总览
这个框架提供以下特性:
1. 将Go程序翻译为Go MPC程序
2. 多种MPC协议实现,统一协议API
3. 最优化混合协议
4. 2方计算

简单来说, 用高级抽象的Go语法编写的常规计算程序, 可以通过此框架转换成MPC程序,以提高MPC程序开发效率, 降低开发难度. 可以的话也提高下MPC程序的运行效率.

输入程序的处理: 在Go语法的基础上拓展MPC语义, 通过翻译工具来完成. 
翻译工具识别特定的语法模式为MPC语义, 然后翻译为等价的纯Go语义实现.
等价的Go实现会调用相对应的MPC协议库来完成MPC计算.

MPC协议: 已经实现了Yao和算术秘密共享两种协议.

最优化混合协议: Effitive MPC via program analysis

主要文件目录:
- cmd 
    - bristol-converter 将cbmc电路转换为bristol电路
    - calibration calib的主程序
    - gtransl 代码翻译的主程序
    - ote 压测ot extension
    - ssadump Not useful
- internal 
    - encrypt 底层密码原语
        - pef aes伪随机发生器
        -sym aes对称加密
    - misc 一些辅助函数
    - network 网络接口
    - ot Oblivious Transfer 
        - baseot  asharov-lindell OT protocol of paper "More Efficient Oblivious Transfer and Extensions for Faster Secure Computation"
        - cot correlated OT
        - rot random OT
        - simpleot simple OT
- pkg
    - always Not useful
    - bristol Bristol Circuit
    - calib 当前环境网络和计算性能的测量benchmark
    - cbmc-gc 用于手工生成和裁剪基础运算电路
    - fast faster XOR operation
    - iast 翻译工具的主要代码, 包含分析和transform pass
    - lz Not useful
    - primitive 乘法三元组
        - multriple 新的乘法三元组
        - triple 旧的乘法三元组, 一些代码还没有换用multiple
    - type 私有变量和公开变量类型
        -  ppub Not useful
        - pub 公开变量类型及其实现
        - pvt 私有变量类型及其实现
        - value pub和pvt的父类型
- template API模版文件, 需要由框架应用程序import
- test 测试,包含了绝大部分功能的单元测试

典型使用方式, 运行gtransl,将生成的文件用Go编译器编译后执行.
其他内部的使用方式参考test文件夹, 其中包含了几乎所有可用函数的测试例程.

## 主要参考文献

### Encrypt
elgamal public key encryption system \
128 bit aes \
fixed key 128 bit aes

### Oblivisou Transfer and extension 
More Efficient Oblivious Transfer and Extensions for Faster Secure Computation \
Random OT extension \
Correlated OT extension \
IKNP \
ALSZ 

### Hybrid Protocol
Effitive MPC via program analysis

### Yao Protocol
Progmatic introduction to mpc

### Arithmetic Protocol
Additive Arithmetic Secret Sharing \
ABY 

### Multiplicative triple
ABY's Multi triple

### Circuit
bristol fashion circuit \
cbmc-gc circuit generator

### example
cryptonet from cryptonet \
bio match \
gcd 

## 其他参考文档
Overview.md\
short_guide.md\
yao.md \
https://homes.esat.kuleuven.be/~nsmart/MPC/\
https://mp-spdz.readthedocs.io/en/latest/readme.html\
https://securecomputation.org/\
https://blog.csdn.net/weixin_44885334/article/details/127084970\
https://www.cs.cornell.edu/~asharov/slides/ALSZ15.pdf\
https://github.com/encryptogroup/ABY/blob/public/README.md\





