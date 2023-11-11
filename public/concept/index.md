# 强化学习数学知识总结

### 基本概念：

对于强化学习，我们一般会分成智能体（agent），环境（通过智能体的状态和动作反馈信息）两大部分，我们现在介绍一些名词，从而有利于之后学习的理解。在这一部分，我们会结合一个3×3的网格寻路来形象化介绍下述概念。

我们最初位置是在（1，1），我们希望能够在最短的时间内到达蓝色网格位置。

![env1](./env-example1.png)

- State：agent相对于environment的状态。例如以s1表示网格（1，1），s9表示网格（3，3）。所有的状态会构成S（State Set）。

- Action：每个状态可能采取的行动。例如向上（a1）、下（a3）、左（a4）、右（a2）走，这所有的动作会构成A（Action Set）。

- state transition：agent从s1→s4的过程。

- state transition probability：用条件概率的形式，表示转移的情况。例如p(s4|s1, a3)=0.5表示从状态s1执行动作a3（向下）到达状态s4的可能性是0.5（因为还可以向右到达s2的可能）。

- Reward：一个数字，用于表示一个state  transition行为所产生的激励。

- forbidden area：进入但是会产生损失的区域（Reward<0）。例如图示中的黄色区域。

- Policy：让agent知道在某个状态该如何行动，也是用条件概率的形式来表示，例如π(a1|s1)=0，π(a2|s1)=0.5。同时我们应该还满足下述条件：
  $$
  \sum_{a\in A}\pi(a|s_j) = 1
  $$

- Trajectory：是整个状态行为链，例如s1(a2)→s2(a3)→s5(a3)→s8(a2)→s9(a4)→s8...

- Return：表示整个Trajectory上Reward之和

- Discounted Rate：由上可知，一个Trajectory可能是无穷无尽的，所以我们需要定义γ来种折扣Reward，从而使整个过程的Return不会太大。

  Discounted Return：表示整个Trajectory的折扣回报总和，具体表示为：
  $$
  Discounted\ Return=\sum_{i=1}^\infin \gamma^{i-1}r_i
  $$
  我们可以知道，这个结果最终会收敛到一个值。并且γ越小，那么Discounted Return会与最早的reward相关（更近视），γ越大也就越远视

- Episode：一个有限步的Trajectory，并且最终状态时terminal state（目标状态，即图中蓝色方块）。

### 马尔可夫过程

##### 随机过程

随机过程研究的对象是，随时间改变的随机现象。例如本例子中，时刻1在状态s1，时刻2可能就会在s2获s4。这一过程是随机的，我们通常用
$$
P(s_{t+1}|a_{t+1},s_{t}, a_{t}, s_{t-1}...a_1, s_0)
$$
来表示下一个时刻，状态的可能性

##### 马尔可夫性质

当前状态只与上一个状态有关时，我们称这个过程具有马尔可夫性质，即
$$
P(s_{t+1}|a_{t+1},s_{t}, a_{t}, s_{t-1}...a_1, s_0)=P(s_{t+1}|a_{t+1}, s_t)
$$
在本例子中，每个时刻到达某个网格的状态可能完全取决于上一时刻所在状态和动作，所以我们这个网格寻路的例子是满足马尔可夫性质的。

##### 几个新的概率

$$
P(s'|s, a)表示s\stackrel{a}{\longrightarrow}s'即从s执行动作a到达s'的概率\\\\
P(r |s, a)表示s\stackrel{a}{\longrightarrow}s'这一过程获得reward=r的概率\\\\
$$

##### 新的概念

- State Value
  $$
  v_{\pi}(s)=\mathbb{E}(G_t|S=s)\\\\
  $$

- 回报矩阵G：表示当前时刻所有状态所获得的discounted return
  $$
  G_t=R_{t+1}+\gamma G_{t+1}
  $$

##### 贝尔曼公式

我们将上述式子进行展开以及归纳，可得贝尔曼公式：
$$
\forall s\in S，v_{\pi}(s)=\mathbb{E}(G_t|S_t=s)=\mathbb{E}(R_{t+1}|S_t=s)+\gamma\mathbb{E}(G_{t+1}|S_t=s)\\\\
=\sum_{a\in A}\pi(a|s)\sum_rp(r|s,a)r + \gamma\sum_{s'}\mathbb{E}(G_{t+1}|S_{t+1}=s',S_t=s)P(s'|s)\\\\
=\sum_{a\in A}\pi(a|s)\sum_rp(r|s,a)r +\gamma\sum_{s'}v_{\pi}(s')\sum_{a\in A}P(s'|s, a)\pi(a|s)
$$
我们将这个公式化成矩阵的形式
$$
\mathbf{v_{\pi}}=\mathbf{r_{\pi}}+\gamma\mathbf{P_{\pi}}\mathbf{v_{\pi}}
$$
其中：
$$
r_{\pi}(s)=\sum_{a\in A}\pi(a|s)\sum_rp(r|s,a)r\\\\
P_{\pi}(s'|s)=\sum_{a \in A}p(s'|s,a)\pi(a|s)，即s\rarr s'的可能性\\\\
\mathbf{P_{\pi(i,j)}} = P_{\pi}(s_j|s_i)
$$


我们该如何求解——求state value

- 求矩阵的逆

- 迭代的方式:

  - 先随机设置v0初始值

  - 利用
    $$
    \mathbf{v_{k+1}}=\mathbf{r_{\pi}}+\gamma\mathbf{P_{\pi}}\mathbf{v_{k}}
    $$
    不断进行迭代，当k→∞时，vk→vπ

