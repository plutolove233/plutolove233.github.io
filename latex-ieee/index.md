# IEEE模板说明


在这个章节中，我们将对IEEE模版进行逐条解析，并补充用IEEE Latex编写论文时，需要注意哪些细节。

和之前的章节一样，我们还是对主tex文件逐条进行解析，力图能够让你看这篇博客就能对模板有个基本的了解。
## 导言区
首先在导言区的第一行：
```latex
\documentclass[lettersize,journal]{IEEEtran}
```
明确了这个文档类型是IEEE的期刊（journal）文件类型，`lettersize`是用来定义文档大小，一般情况下，可以将其修改为：
```latex
\documentclass[a4paper,journal]{IEEEtran}
```
即定义文档大小为A4纸。

之后是引用其他包：
```latex
\usepackage{algorithmic}
\usepackage{algorithm}
\usepackage{array}
\usepackage[caption=false,font=normalsize,labelfont=sf,textfont=sf]{subfig}
\usepackage{textcomp}
\usepackage{stfloats}
\usepackage{url}
\usepackage{verbatim}
\usepackage{graphicx}
\usepackage{cite}
```
对于每个包，我们不需要都知道它们的意义，我们只需要在需要的时候引用即可。

之后有这么一段代码：
```latex
\hyphenation{op-tical net-works semi-conduc-tor IEEE-Xplore}
```
这段代码的含义是表示换行连词分词划分方式，一般情况下不需要对这里的内容进行修改。

## 标题、作者以及页眉页脚
首先是关于文档标题：
```latex
\begin{document}

\title{A Sample Article Using IEEEtran.cls\\ for IEEE Journals and Transactions}
```
作者：
```latex
\author{IEEE Publication Technology,~\IEEEmembership{Staff,~IEEE,}
\thanks{This paper was produced by the IEEE Publication Technology Group. They are in Piscataway, NJ.}
\thanks{Manuscript received April 19, 2021; revised August 16, 2021.}}
```
与Elsevier期刊不同的是，IEEE一般会在这里申明致谢标注基金信息，而不会单独拎一个章节来写。举个例子：
```latex
\author{
    ...
    \thanks{Financial supports from State Grid Corporation of China Company Science and Technology Project Research (No. xxxx). (Corresponding author: xxx.)}
    \thanks{...}
}
```
同时，会在第二个`\thanks{}`中申明作者的单位以及邮件信息，例如：
```latex
\author{
    aaa, bbb, ccc, ddd
    \thanks{...}
    \thanks{aaa, bbb, ccc, ddd are with the Peking University(email: aaa@126.com; bbb@139.com; ccc@qq.com; ddd\_2024@163.com)}
}
```
之后是论文页眉以及页脚，这里就不需要我们来修改，一般会由期刊主编对其进行修改调整：
```latex

\markboth{Journal of \LaTeX\ Class Files,~Vol.~14, No.~8, August~2021}%
{Shell \MakeLowercase{\textit{et al.}}: A Sample Article Using IEEEtran.cls for IEEE Journals}

\IEEEpubid{0000--0000/00\$00.00~\copyright~2021 IEEE}

\maketitle
```
## 摘要与关键字
摘要和关键字如下，只要和之前的章节一样替换内容即可：
```latex
\begin{abstract}
This document describes the most common article elements and how to use the IEEEtran class with \LaTeX \ to produce files that are suitable for submission to the IEEE.  IEEEtran can produce conference, journal, and technical note (correspondence) papers with a suitable choice of class options. 
\end{abstract}

\begin{IEEEkeywords}
Article submission, IEEE, IEEEtran, journal, \LaTeX, paper, template, typesetting.
\end{IEEEkeywords}
```
与ArXiv模板不同的是，关键字是用`IEEEkeywords`来定义

## 正文
正文部分就和之前博客中的内容一样了，可以按照自己的喜好进行排列编写，需要注意的是：
- IEEE论文正文的第一段话，一般需要采用`\IEEEPARstart{T}{his}`的方式来使第一个单词全部大写，并且第一个字母占两行
- 如果出现某个段落的文字内容与页脚重叠了，可以在这个段落前，加`\IEEEpubidadjcol`来避免这个问题，例如：
	```latex
	\IEEEpubidadjcol
	However, ...
	```

## 参考文献
模板中关于参考文献的定义如下：
```latex
\begin{thebibliography}{1}
\bibliographystyle{IEEEtran}

\bibitem{ref1}
{\it{Mathematics Into Type}}. American Mathematical Society. [Online]. Available: https://www.ams.org/arc/styleguide/mit-2.pdf

\bibitem{ref2}
T. W. Chaundy, P. R. Barrett and C. Batey, {\it{The Printing of Mathematics}}. London, U.K., Oxford Univ. Press, 1954.

\bibitem{ref3}
F. Mittelbach and M. Goossens, {\it{The \LaTeX Companion}}, 2nd ed. Boston, MA, USA: Pearson, 2004.

...

\end{thebibliography}
```

但是由于我们一般都是采用bib文件对参考文献进行管理，因此可以将这部分内容替换为：
```latex
\bibliographystyle{IEEEtran}
\bibliography{references}
```

## 作者信息
一般情况下，IEEE论文都会在参考文献后，放上作者的信息，以下这部分代码即为作者信息的参考：
```latex
\vspace{-33pt}
\begin{IEEEbiography}[{\includegraphics[width=1in,height=1.25in,clip,keepaspectratio]{fig1}}]{Michael Shell}
Use $\backslash${\tt{begin\{IEEEbiography\}}} and then for the 1st argument use $\backslash${\tt{includegraphics}} to declare and link the author photo.
Use the author name as the 3rd argument followed by the biography text.
\end{IEEEbiography}
\vfill
```
可以看到在`IEEEbiography`后的方括号中是作者照片的图片，之后第一个花括号中是作者的姓名，第二个花括号中是作者的介绍。

作者部分可以在第一个作者简介前加一个`\newpage`让作者信息另起一栏，这样会更美观些。

## 结语
IEEE模板不会像另外两类模板会出现很多奇怪的bug需要注意，大家可以在熟悉后，用该模板投递IEEE期刊论文，祝大家paper拿到手软～
