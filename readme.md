# cal
cal is a simple commmand-line calculator based on cobra and pflag. A demo for learning those two libs.

# introduction
cal 是一个实现了加减乘除的命令行计算器，它使用四个子命令`add`,`sub`,`mul`,`div`实现了加减乘除操作。它简易的演示了
* 如何使用cobra搭建命令行程序
* 如何使用cobra生成子命令
* 如何在父子命令之间传参
* 如何声明可以父子命令行传递的选项或独有选项
* 如何声明选项的短版本

`cal`的`div`子命令包含一个独有的参数`--precision`或`-p`来指示结果保留小数点后几位

具体的cobra使用方法参考总结：[cobra.md](https://github.com/DaoiestFire/Go-learning-summary/blob/main/cobra.md)

# Structure
```
│  go.mod
│  go.sum
│  main.go                       //main函数入口，实例化rootCmd并启动
│  readme.md
│
└─cmd
    │  cal.go                    //实现rootCmd和四个计算子命令
    │
    └─options
            options.go           //定义options类型存放两个操作数，演示如何进行复杂传参
```

# Usage
生成二进制文件
```shell
cd ./cal
go build
```

用法
```shell
# add
./cal.exe add --argone 2 --argtwo 3
result: 5

# sub
./cal.exe sub --argone 2 --argtwo 3
result: -1

# mul
./cal.exe mul --argone 2 --argtwo 3
result: 6

# div
./cal.exe div --argone 2 --argtwo 3
result: 5

# div
./cal.exe add --argone 2 --argtwo 3 -p 4
result: 0.6667

# -v --version
./cal.exe -v
cal version 0.1

# ./cal
./cal.exe
cal is simple calculator app based on cobra library
it supports add/sub/div/mul four operation.
hurry to try it!

Usage:
  cal [command]

Available Commands:
  add
  completion  Generate the autocompletion script for the specified shell
  div
  help        Help about any command
  mul
  sub

Flags:
      --argone int   assign a value to arg1
      --argtwo int   assign a value to arg2
  -h, --help         help for cal
  -v, --version      version for cal

Use "cal [command] --help" for more information about a command.
```