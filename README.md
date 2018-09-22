# cstruct-go

a fast c-style struct packer & unpacker for golang

## 用法

1. 定义举例

  ```go
  type mystruct1 struct {
      F1  bool
      F2  float32
      F3  float64
      F4  string
      F5  []byte
      F6  int8
      F7  int16
      F9  int32
      F11 int64
      F12 uint8
      F13 uint16
      F15 uint32
      F17 uint64
      F18 [20]byte
      F19 [16]uint32
  }
  ```

2. 使用举例

  ```go
  a := &mystruct1{}
  // ...(初始化a 代码略)...

  // 序列化代码如下，返回 []byte类型数据
  buf_l, _ := cstruct.Marshal(a)

  // 反序列化代码如下
  b := &mystruct1{}
  if err := cstruct.Unmarshal(buf_l, b); err != nil {
      fmt.Println(err)
      return
  }
  ```

  详细例子可以参考：

  - [x_test.go](tests/x_test.go)
  - [array_test.go](tests/array_test.go)

## 字节序

`小端`字节序。

## 基本类型

go类型    | 内存说明
------- | ----------------------------
bool    | 1 byte
int8    | 1 byte
uint8   | 1 byte
int16   | 2 byte
uint16  | 2 byte
int32   | 4 byte
uint32  | 4 byte
int64   | 8 byte
uint64  | 8 byte
float32 | 4 byte
float64 | 8 byte
string  | [2 byte] + [len(字符串) byte]
[]byte  | [2 byte] + [len(2进制数据) byte]


## 数组类型

- 支持基本类型的数组（string 、 []byte 除外，因为它们不定长）
- 支持定长struct的数组


## 指针类型

- 支持struct指针

go类型        | 内存说明                                |byte含义说明
------------ | ---------------------------------------|--------------
struct ptr   | [1 byte] + [len(struct内存占用) byte]   | 1byte值为0，表示指针为nil


## 复杂类型

- 支持struct嵌套
- 支持基本类型的Slice
- 支持struct指针的Slice

go类型  | 内存说明
----- | -----------------------------
slice | [2 byte] + [len(元素内存占用) byte]

## 基准测试

```dos
D:\golang\src\github.com\fananchong\cstruct-go\benchmarks>call go test -test.bench=".*" -count=1
goos: windows
goarch: amd64
pkg: github.com/fananchong/cstruct-go/benchmarks
Benchmark_CStructGO-4            1000000              1877 ns/op
Benchmark_Protobuf-4              300000              5003 ns/op
Benchmark_GoGoProtobuf-4         1000000              2038 ns/op
PASS
ok      github.com/fananchong/cstruct-go/benchmarks     5.899s
```

基准测试代码：[cstrucgo_test.go](benchmarks/cstrucgo_test.go)

## TODO

- 增加定长数组类型

## 参考项目

- <https://github.com/golang/protobuf>
- <https://github.com/gogo/protobuf>
