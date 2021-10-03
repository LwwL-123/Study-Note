# Goroutine



## 1. 什么是Goroutine

**Goroutine是一个由Go运行时管理的轻量级线程，一般称其为”协程“**

```go
go f(a,b,c)
```

操作系统本身无法感知到Goroutine的存在，Goroutine的操作和切换归属于”用户态“

Goroutine 由特定的调度模式来控制，以 “多路复用” 的形式运行在操作系统为 Go 程序分配的几个系统线程上。

同时创建 Goroutine 的开销很小，初始只需要 2-4k 的栈空间。Goroutine 本身会根据实际使用情况进行自伸缩，非常轻量。



## 2.调度是什么

既然有了用户态的代表 Goroutine，操作系统又看不到他。必然需要有某个东西去管理他，才能更好的运作起来。

这指的就是 Go 语言中的调度，最常见、面试最爱问的 GMP 模型。



### 调度基础知识

Go scheduler 的主要功能是针对在处理器上运行的 OS 线程分发可运行的 Goroutine，而我们一提到调度器，就离不开三个经常被提到的缩写，分别是：

- G：Goroutine，实际上我们每次调用 `go func` 就是生成了一个 G。
- P：Processor，处理器，一般 P 的数量就是处理器的核数，可以通过 `GOMAXPROCS` 进行修改。
- M：Machine，系统线程。