# Linux常见命令

## nohup命令

nohup命令用于不挂断地运行命令（关闭当前session不会中断改程序，只能通过kill等命令删除）。
使用nohup命令提交作业，如果使用nohup命令提交作业，那么在缺省情况下该作业的所有输出都被重定向到一个名为nohup.out的文件中，除非另外指定了输出文件。



1. 示例

```bash
nphup ./ttchain daemon >provider.log 2>&1 &
```



2. 2>&1

```bash
bash中：
0 代表STDIN_FILENO 标准输入（一般是键盘），
1 代表STDOUT_FILENO 标准输出（一般是显示屏，准确的说是用户终端控制台），
2 三代表STDERR_FILENO (标准错误（出错信息输出）。
```

```
> 直接把内容生成到指定文件，会覆盖原来文件中的内容[ls > test.txt],
>> 尾部追加，不会覆盖原有内容 [ls >> test.txt],
< 将指定文件的内容作为前面命令的参数[cat < text.sh]
```

2>&1就是用来将标准错误2重定向到标准输出1中的。此处1前面的&就是为了让bash将1解释成标准输出而不是文件1。至于最后一个&，则是让bash在后台执行。





## lsof

l