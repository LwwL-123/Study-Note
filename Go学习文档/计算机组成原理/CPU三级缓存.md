# **CPU的三级缓存**



CPU缓存的定义为**CPU**与**内存**之间的临时数据交换器

它的出现是为了解决CPU运行处理速度与内存读写速度不匹配的矛盾——缓存的速度比内存的速度快多了。

CPU缓存一般直接跟CPU芯片集成或位于主板总线互连的独立芯片上。（现阶段的CPU缓存一般直接集成在CPU上）CPU往往需要重复处理相同的数据、重复执行相同的指令，如果这部分数据、指令CPU能在CPU缓存中找到，CPU就不需要从内存或硬盘中再读取数据、指令，从而减少了整机的响应时间。



## 缓存时间差距

我们来简单地打个比方：如果CPU在L1一级缓存中找到所需要的资料要用的时间为3个周期左右，那么在L2二级缓存找到资料的时间就要10个周期左右，L3三级缓存所需时间为50个周期左右；如果要到内存上去找呢，那就慢多了，可能需要几百个周期的时间。



## **三级缓存（L1、L2、L3）是什么？**

以近代CPU的视角来说，三级缓存（包括L1一级缓存、L2二级缓存、L3三级缓存）都是集成在CPU内的缓存，它们的作用都是作为CPU与主内存之间的高速数据缓冲区，

- L1最靠近CPU核心；L2其次；L3再次
- 运行速度方面：L1最快、L2次快、L3最慢；
- 容量大小方面：L1最小、L2较大、L3最大。
- CPU会先在最快的L1中寻找需要的数据，找不到再去找次快的L2，还找不到再去找L3，L3都没有那就只能去内存找了。L1、L2、L3各有特点



### **一级缓存（L1 Cache）**

CPU一级缓存，就是指CPU的第一层级的高速缓存，主要当担的工作是缓存指令和缓存数据。一级缓存的容量与结构对CPU性能影响十分大，但是由于它的结构比较复杂，又考虑到成本等因素，一般来说，CPU的一级缓存较小，通常CPU的一级缓存也就能做到256KB左右的水平。



### 二级缓存（L2 Cache）

CPU二级缓存，就是指CPU的第二层级的高速缓存，而二级缓存的容量会直接影响到CPU的性能，二级缓存的容量越大越好。例如intel的第八代i7-8700处理器，共有六个核心数量，而每个核心都拥有256KB的二级缓存，属于各核心独享，这样二级缓存总数就达到了1.5MB。



### 三级缓存（L3 Cache）

CPU三级缓存，就是指CPU的第三层级的高速缓存，其作用是进一步降低内存的延迟，同时提升海量数据量计算时的性能。和一级缓存、二级缓存不同的是，三级缓存是核心共享的，能够将容量做的很大。



例：下列哪个效率高

```go
func a(){
  for i:=0;i<100;i++{
    for j:=0;j<100;j++{
      A[i][j] = A[i][j] + B[i][j]
    }
  }
}
```

```go
func b(){
  for i:=0;i<100;i++{
    for j:=0;j<100;j++{
      A[i][j] = A[i][j] + B[j][i]
    }
  }
}
```



答：a()的效率高，涉及到内存的cache，通常CPU都会有三级缓存，a中读取的变量都是在连续的内存里面，b中每次读取都是不连续的，会涉及到重新加载，从内存读取数据到cache中，如果cache不充足，还会涉及到cache的一些修改，没有充分利用到cache的速度



#  MESI 协议

多核 CPU 里的每一个 CPU 核，都有独立的属于自己的 L1 Cache 和 L2 Cache。多个 CPU 之间，只是共用 L3 Cache 和主内存。

我们说，CPU Cache 解决的是内存访问速度和 CPU 的速度差距太大的问题。而多核 CPU 提供的是，在主频难以提升的时候，通过增加 CPU 核心来提升 CPU 的吞吐率的办法。我们把多核和 CPU Cache 两者一结合，就给我们带来了一个新的挑战。因为 CPU 的每个核各有各的缓存，互相之间的操作又是各自独立的，就会带来[**缓存一致性**](https://en.wikipedia.org/wiki/Cache_coherence)（Cache Coherence）的问题。

![img](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211027151608.jpeg)

MESI 是指4中状态的首字母。每个Cache line有4个状态，可用2个bit表示，它们分别是：

> **缓存行（Cache line）**:缓存存储数据的单元。

- M：代表已修改（Modified）
- E：代表独占（Exclusive）
- S：代表共享（Shared）
- I：代表已失效（Invalidated）



我们先来看看 “已修改” 和 “已失效”，这两个状态比较容易理解。所谓的 “已修改”，就是我们所说的 “脏” 的 Cache Block。Cache Block 里面的内容我们已经更新过了，但是还没有写回到主内存里面。而所谓的 “已失效 “，自然是这个 Cache Block 里面的数据已经失效了，我们不可以相信这个 Cache Block 里面的数据。



然后，我们再来看 “独占” 和 “共享” 这两个状态。这就是 MESI 协议的精华所在了。无论是独占状态还是共享状态，缓存里面的数据都是 “干净” 的。这个 “干净”，自然对应的是前面所说的 “脏” 的，也就是说，这个时候，Cache Block 里面的数据和主内存里面的数据是一致的。



那么 “独占” 和 “共享” 这两个状态的差别在哪里呢？这个差别就在于，在独占状态下，对应的 Cache Line 只加载到了当前 CPU 核所拥有的 Cache 里。其他的 CPU 核，并没有加载对应的数据到自己的 Cache 里。这个时候，如果要向独占的 Cache Block 写入数据，我们可以自由地写入数据，而不需要告知其他 CPU 核。

在独占状态下的数据，如果收到了一个来自于总线的读取对应缓存的请求，它就会变成共享状态。这个共享状态是因为，这个时候，另外一个 CPU 核心，也把对应的 Cache Block，从内存里面加载到了自己的 Cache 里来。



而在共享状态下，因为同样的数据在多个 CPU 核心的 Cache 里都有。所以，当我们想要更新 Cache 里面的数据的时候，不能直接修改，而是要先向所有的其他 CPU 核心广播一个请求，要求先把其他 CPU 核心里面的 Cache，都变成无效的状态，然后再更新当前 Cache 里面的数据。这个广播操作，一般叫作 RFO（Request For Ownership），也就是获取当前对应 Cache Block 数据的所有权。



![img](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211027152314.jpeg)

### 例：



假设有三个CPU A、B、C，对应三个缓存分别是cache a、b、 c。在主内存中定义了x的引用值为0。

![在这里插入图片描述](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211027153429.png)

- 单核读取
  那么执行流程是：
  CPU A发出了一条指令，从主内存中读取x。
  从主内存通过bus读取到缓存中（远端读取Remote read）,这是该Cache line修改为E状态（独享）.

![在这里插入图片描述](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211027153512.png)



- 双核读取
  那么执行流程是：
  CPU A发出了一条指令，从主内存中读取x。
  CPU A从主内存通过bus读取到 cache a中并将该cache line 设置为E状态。
  CPU B发出了一条指令，从主内存中读取x。
  CPU B试图从主内存中读取x时，CPU A检测到了地址冲突。这时CPU A对相关数据做出响应。此时x 存储于cache a和cache b中，x在chche a和cache b中都被设置为S状态(共享)。

![在这里插入图片描述](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211027153527.png)

- 修改数据
  那么执行流程是：
  CPU A 计算完成后发指令需要修改x.
  CPU A 将x设置为M状态（修改）并通知缓存了x的CPU B, CPU B将本地cache b中的x设置为I状态(无效)
  CPU A 对x进行赋值。

![在这里插入图片描述](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211027153544.png)

- 同步数据
  那么执行流程是：
  CPU B 发出了要读取x的指令。
  CPU B 通知CPU A,CPU A将修改后的数据同步到主内存时cache a 修改为E（独享）
  CPU A同步CPU B的x,将cache a和同步后cache b中的x设置为S状态（共享）。

![在这里插入图片描述](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211027153556.png)

