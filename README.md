# cstruct-go

c-style binary data pack & unpack for golang

## 用法

1.  定义举例

    ```go
    type mystruct1 struct {
        F1  bool    `c:"bool"`
        F2  float32 `c:"float"`
        F3  float64 `c:"double"`
        F4  string `c:"string"`
        F5  []byte `c:"binary"`
        F6  int8   `c:"int8"`
        F7  int16  `c:"int16"`
        F8  int32  `c:"int24"`
        F9  int32  `c:"int32"`
        F10 int64  `c:"int40"`
        F11 int64  `c:"int64"`
        F12 uint8  `c:"uint8"`
        F13 uint16 `c:"uint16"`
        F14 uint32 `c:"uint24"`
        F15 uint32 `c:"uint32"`
        F16 uint64 `c:"uint40"`
        F17 uint64 `c:"uint64"`
    }
    ```

2.  使用举例

    ```go
    a := &mystruct1{}
    // ...(初始化a 代码略)...

    // 序列化代码如下，返回 []byte类型数据
    buf_l := cstruct.PackLE(a)

    // 反序列化代码如下
    b := &mystruct1{}
    if err := cstruct.UnpackLE(buf_l, b); err != nil {
        fmt.Println(err)
        return
    }
    ```

## 字节序

| API      | 说明              |
| -------- | --------------- |
| PackLE   | 序列化为 `小端` 数据    |
| UnpackLE | 将数据视为 `小端`序，反序列 |
| PackBE   | 序列化为 `大端` 数据    |
| UnpackBE | 将数据视为 `大端`序，反序列 |

## 复杂类型

-   支持嵌套struct
-   支持不带`Tag`字段（不带`Tag`字段，不参与序列化、反序列化）

## 基本类型

| cstruct类型 | go类型    | 内存说明                       |
| --------- | ------- | -------------------------- |
| bool      | bool    | 1 byte                     |
| int8      | int8    | 1 byte                     |
| uint8     | uint8   | 1 byte                     |
| int16     | int16   | 2 byte                     |
| uint16    | uint16  | 2 byte                     |
| int24     | int32   | 3 byte                     |
| uint24    | uint32  | 3 byte                     |
| int32     | int32   | 4 byte                     |
| uint32    | uint32  | 4 byte                     |
| int40     | int64   | 5 byte                     |
| uint40    | uint64  | 5 byte                     |
| int64     | int64   | 8 byte                     |
| uint64    | uint64  | 8 byte                     |
| float     | float32 | 4 byte                     |
| double    | float64 | 8 byte                     |
| string    | string  | len(字符串)[2 byte] + 字符串     |
| binary    | \[]byte | len(2进制数据)[2 byte] + 2进制数据 |


## TODO

  - 参考go protobuf实现，优化buffer
  - 参考go protobuf实现，优化反射代码
  
  
