# cstruct-go

c-style struct pack & unpack for golang


## 用法

1.  定义举例

    ```go
    type mystruct1 struct {
	    F1  bool    `c:"bool"`
	    F2  float32 `c:"float"`
	    F3  float64 `c:"double"`
	    F4  string  `c:"string"`
	    F5  []byte  `c:"binary"`
	    F6  int8    `c:"int8"`
	    F7  int16   `c:"int16"`
	    F9  int32   `c:"int32"`
	    F11 int64   `c:"int64"`
	    F12 uint8   `c:"uint8"`
	    F13 uint16  `c:"uint16"`
	    F15 uint32  `c:"uint32"`
	    F17 uint64  `c:"uint64"`
    }
    ```

2.  使用举例

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

## 字节序

默认`小端`字节序。

可以通过定义下列语句，改变字节序：

```go
cstruct.CurrentByteOrder = cstruct.LE
```

或者

```go
cstruct.CurrentByteOrder = cstruct.BE
```


## 基本类型

| cstruct类型 | go类型    | 内存说明                 |
| --------- | ------- | -------------------------- |
| bool      | bool    | 1 byte                     |
| int8      | int8    | 1 byte                     |
| uint8     | uint8   | 1 byte                     |
| int16     | int16   | 2 byte                     |
| uint16    | uint16  | 2 byte                     |
| int32     | int32   | 4 byte                     |
| uint32    | uint32  | 4 byte                     |
| int64     | int64   | 8 byte                     |
| uint64    | uint64  | 8 byte                     |
| float     | float32 | 4 byte                     |
| double    | float64 | 8 byte                     |
| string    | string  | [2 byte] + [len(字符串)]    |
| binary    | \[]byte | [2 byte] + [len(2进制数据)] |


## 基准测试

```dos
D:\golang\src\github.com\fananchong\cstruct-go\benchmarks>go test -test.bench=".*" -count=1
goos: windows
goarch: amd64
pkg: github.com/fananchong/cstruct-go/benchmarks
Benchmark_CStructGO-4            1000000              1206 ns/op
Benchmark_Protobuf-4              500000              2708 ns/op
PASS
ok      github.com/fananchong/cstruct-go/benchmarks     2.975s
```

## 参考项目

  - https://github.com/golang/protobuf
  
  
