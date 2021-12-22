# Go内存分配

## 小于32KB内存块的分配策略

当程序里发生了`32kb`以下的小块内存申请时，Go会从一个叫做的`mcache`的本地缓存给程序分配内存。这个本地缓存`mcache`持有一系列的大小为`32kb`的内存块，这样的一个内存块里叫做`mspan`，它是要给程序分配内存时的分配单元。

![image-20211222153644273](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211222153644.png)

在Go的调度器模型里，每个线程`M`会绑定给一个处理器`P`，在单一粒度的时间里只能做多处理运行一个`goroutine`，每个`P`都会绑定一个上面说的本地缓存`mcache`。当需要进行内存分配时，当前运行的`goroutine`会从`mcache`中查找可用的`mspan`。从本地`mcache`里分配内存时不需要加锁，这种分配策略效率更高。

那么有人就会问了，有的变量很小就是数字，有的却是一个复杂的结构体，申请内存时都分给他们一个`mspan`这样的单元会不会产生浪费。其实`mcache`持有的这一系列的`mspan`并不都是统一大小的，而是按照大小，从8字节到32KB分了大概70类的`msapn`。

![image-20211222154003875](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211222154004.png)

如果结构体的大小是32字节，正好32字节的这种`mspan`能满足需求，那么分配内存的时候就会给它分配一个32字节大小的`mspan`。

![image-20211222154151132](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211222154151.png)



现在，我们可能会好奇，如果分配内存时`mcachce`里没有空闲的32字节的`mspan`了该怎么办？`Go`里还为每种类别的`mspan`维护着一个`mcentral`。

`mcentral`的作用是为所有`mcache`提供切分好的`mspan`资源。每个`central`会持有一种特定大小的全局`mspan`列表，包括已分配出去的和未分配出去的。每个`mcentral`对应一种`mspan`，当工作线程的`mcache`中没有合适（也就是特定大小的）的`mspan`时就会从`mcentral` 去获取。`mcentral`被所有的工作线程共同享有，存在多个`goroutine`竞争的情况，因此从`mcentral`获取资源时需要加锁。



`mcentral`的定义如下：

```
//runtime/mcentral.go

type mcentral struct {
    // 互斥锁
    lock mutex 
    
    // 规格
    sizeclass int32 
    
    // 尚有空闲object的mspan链表
    nonempty mSpanList 
    
    // 没有空闲object的mspan链表，或者是已被mcache取走的msapn链表
    empty mSpanList 
    
    // 已累计分配的对象个数
    nmalloc uint64 
}
```

`mcentral`里维护着两个双向链表，**nonempty**表示链表里还有空闲的`mspan`待分配。**empty**表示这条链表里的`mspan`都被分配了`object`。

![图片](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211222154531)

如果上面我们那个程序申请内存的时候，`mcache`里已经没有合适的空闲`mspan`了，那么工作线程就会像下图这样去`mcentral`里去申请。

简单说下`mcache`从`mcentral`获取和归还`mspan`的流程：

- 获取 加锁；从`nonempty`链表找到一个可用的`mspan`；并将其从`nonempty`链表删除；将取出的`mspan`加入到`empty`链表；将`mspan`返回给工作线程；解锁。
- 归还 加锁；将`mspan`从`empty`链表删除；将`mspan`加入到`nonempty`链表；解锁。

![image-20211222154842718](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211222154842.png)





当`mcentral`没有空闲的`mspan`时，会向`mheap`申请。而`mheap`没有资源时，会向操作系统申请新内存。`mheap`主要用于大对象的内存分配，以及管理未切割的`mspan`，用于给`mcentral`切割成小对象。

![图片](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211222155531)

同时我们也看到，`mheap`中含有所有规格的`mcentral`，所以，当一个`mcache`从`mcentral`申请`mspan`时，只需要在独立的`mcentral`中使用锁，并不会影响申请其他规格的`mspan`。

上面说了每种尺寸的`mspan`都有一个全局的列表存放在`mcentral`里供所有线程使用，所有`mcentral`的集合则是存放于`mheap`中的。`mheap`里的`arena` 区域是真正的堆区，运行时会将 `8KB` 看做一页，这些内存页中存储了所有在堆上初始化的对象。运行时使用二维的 runtime.heapArena 数组管理所有的内存，每个 runtime.heapArena 都会管理 64MB 的内存。



![图片](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211222155625)

如果 `arena` 区域没有足够的空间，会调用 runtime.mheap.sysAlloc 从操作系统中申请更多的内存。



## 大于32KB内存块的分配策略

Go没法使用工作线程的本地缓存`mcache`和全局中心缓存`mcentral`上管理超过32KB的内存分配，所以对于那些超过32KB的内存申请，会直接从堆上(`mheap`)上分配对应的数量的内存页（每页大小是8KB）给程序。

![图片](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211222155723)



## 总结

我们把内存分配管理涉及的所有概念串起来，可以勾画出Go内存管理的一个全局视图：

![图片](https://gitee.com/lzw657434763/pictures/raw/master/Blog/20211222155735)

总结起来关于Go内存分配管理的策略有如下几点：

- Go在程序启动时，会向操作系统申请一大块内存，由`mheap`结构全局管理。
- Go内存管理的基本单元是`mspan`，每种`mspan`可以分配特定大小的`object`。
- `本地缓存mcache`, `全局中心缓存mcentral`, `mheap`是`Go`内存管理的三大组件，
  - `mcache`管理线程在本地缓存的`mspan`；
  - `mcentral`管理全局的`mspan`供所有线程使用；
  - `mheap`管理`Go`的所有动态分配内存。
- 一般小对象通过`mspan`分配内存；大对象则直接由`mheap`分配内存。





本文节选自：wx网管叨bi叨

