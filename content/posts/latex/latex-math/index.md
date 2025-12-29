+++
title = 'LaTeX数学公式编写'
date = 2025-12-29T16:18:01+08:00
draft = false
categories = ["LaTeX论文写作"]
tags = ["Latex", "论文基本功"]
+++

latex中数学符号表达形式多样，各种运算符难以记忆。为此特立一个章节，用于整理以及对公式注意事项进行说明。
1. [数学符号表，包含各类运算符，希腊字母](https://blog.csdn.net/qq_31047111/article/details/105598163)
2. [箭头符号，以及在箭头上添加字母符号的方法](https://blog.csdn.net/J__Max/article/details/86549124)
3. [花体字母](https://blog.csdn.net/weixin_44550010/article/details/107864356)，需要注意的是，花体字母一般只作用于**单个大写**字母。
4. [LaTex中把下标置于文本正下方的方法](https://blog.csdn.net/da_kao_la/article/details/84836098)
5. [LaTeX 特殊符号、加帽子符号、横线和波浪线](https://blog.csdn.net/qq_17528659/article/details/82152530)
6. [latex写矩阵](https://surprisedcat.github.io/%E7%BD%91%E9%A1%B5%E8%B5%84%E6%96%99/Latex-%E5%A6%82%E4%BD%95%E7%94%A8latex%E7%BC%96%E5%86%99%E7%9F%A9%E9%98%B5%EF%BC%88%E5%8C%85%E6%8B%AC%E5%90%84%E7%B1%BB%E5%A4%8D%E6%9D%82%E3%80%81%E5%A4%A7%E5%9E%8B%E7%9F%A9%E9%98%B5%EF%BC%89.html)

## 公式对齐
有的公式会特别长，在双栏论文里根本放不下，超出所处那一栏的区域，可采用如下方法进行处理：
```latex
\begin{equation}
\begin{aligned}
L_{task} &= \lambda_1 L_{per}(G_s(x),G_t(x))  +\lambda_1  L_{CE}(y,\delta(z_s))\\ 
& + \lambda_1L_{Focal}(y,y_{out})
\end{aligned}
\end{equation}
```
其中，需采用aligned包裹区域，才能进行换行以及对齐。用`&`标注对齐的位置，用`\\`表示换行，呈现效果如下：

![img](./image.png)

## 公式大小调整
如果公式稍微大了些，用换行表示也不太方便，可以参考[这篇博客](https://blog.csdn.net/zou_albert/article/details/110532165)对公式字体大小进行调整。

## 一些类型符号的写法
|数学对象|Latex表达方式|例子|说明|
|:----:|:---:|:---:|:--:|
|向量|\vec{x}|$\vec{x}$|一般为小写|
|矩阵|\mathbf{X}|$\mathbf{X}$|一般为大写|
|集合|\mathcal{S}|$\mathcal{S}$|一般为大写|



