# 例子说明

演示 C struct 与 Go struct 内存分布一致

- main.cpp

  C++ 结构体内存按 1 对齐，并输出结构体对象内存分布

- main.go

  Go 结构体定义与 C++ 结构体一致，并做序列化，然后输出序列化数据的内存分布

# 输出结果

```shell
D:\golang\src\github.com\fananchong\cstruct-go\example>main.exe
len(b) = 76
struct data len =  47
struct data is:
127 56 231 3 0 0 0 1 2 3 4 184 34 154 153 177 66 56 231 3 0 0 0 1 2 3 4 56 231 3 0 0 0 1 2 3 4 56 231 3 0 0 0 1 2 3 4
D:\golang\src\github.com\fananchong\cstruct-go\example>main_cpp.exe
len(b) = 47
struct data len = 47
struct data is:
127 56 231 3 0 0 0 1 2 3 4 184 34 154 153 177 66 56 231 3 0 0 0 1 2 3 4 56 231 3 0 0 0 1 2 3 4 56 231 3 0 0 0 1 2 3 4
```

可以看到，原始 Go 结构体大小为 76

序列化后，大小为 47 ，且内存分布与 C++ 结构体内存分布一致
