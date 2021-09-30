# defer

```go
package main

func main() {
   println(f1())
   println(f2())
   println(f3())
}

func f1() (r int) {
   defer func() {
      r++
   }()
   return 0
}

func f2() (r int) {
   t := 5
   defer func() {
      t = t + 5
   }()
   return t
}

func f3() (r int) {
   defer func(r int) {
      r = r + 5
   }(r)
   return 1
}
```

输出:

```
1
5
1
```



## 为什么

1. 返回r，将0赋值给r，随后在defer中，对r进行+1，所以返回1
2. 返回t=5，将5赋值给r，随后对t进行+5，对r没有影响，所以返回5
3. 将1赋值给r，但defer中是函数变量r，与返回值没有关系，所以返回1



## 总结

1. 多个defer的执行顺序是`先进后出，后进先出`，与栈相同

2. **有返回值的且带有 defer 函数的方法中， return 语句执行顺序：**

   ```
   1. 返回值赋值
   2. 调用 defer 函数 (在这里是可以修改返回值的)
   3. return 返回值
   ```

3. 匿名返回值是在 return 执行时被声明，有名返回值则是在函数声明的同时被声明，因此在 defer 语句中只能访问有名返回值，而不能直接访问匿名返回值；