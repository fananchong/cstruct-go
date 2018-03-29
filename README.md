# cstruct-go

c-style binary data pack & unpack for golang

## TODO

实现类似以下用法：

- define

  ```go
    type mystruct1 struct {
        f0 bool      `c:bool`
        f1 int8      `c:int8`
        f2 float32   `c:float`
        f3 float64   `c:double`
        f4 string    `c:string`
        f5 []byte    `c:binary`
        f6 mystruct2 `c:struct`
    }
  ```

- pack (le && be)

  ```go
    func pack(obj interface{}) []byte {

    }
  ```

- unpack (le && be)

  ```go
    func unpack([]byte) interface{} {

    }
  ```
