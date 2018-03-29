# cstruct-go

c-style binary data pack & unpack for golang

## TODO

实现类似以下用法：

- 定义

  ```go
    type mystruct1 struct {
        f0 bool      `c:bool`
        f1 int8      `c:int8`
        f2 float32   `c:float`
        f3 float64   `c:double`
        f4 string    `c:string(32)`
        f5 []byte    `c:binary(24)`
        f6 mystruct2 `c:struct`
    }
  ```

- pack

  ```go
    func pack(obj) []byte {

    }
  ```

- unpack

  ```go
    func unpack([]byte) interface{} {

    }
  ```
