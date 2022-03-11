# Go AutoMapper

*This repository is an extension or fork edition of [stroiman/go-automapper](https://github.com/stroiman/go-automapper)*

Go-automapper can automatically map data between different types with identically named fields in GO. Useful for initializing DTO types from build in data.



**Extension**

* This package only use `MapLoose` function of [stroiman/go-automapper](https://github.com/stroiman/go-automapper) to avoid any kind failure when mapping into destination type.
* This package allows to pass a callback function while invoking `Map` function to map unknown fields of destination object



***

### How to use

Install automapper in your go project

```
go get github.com/rafiulgits/go-automapper
```



* Object - Object Mapping

```go
type A struct {
  Name  string
  Phone string
}

type B struct {
  Name  string
  Phone string
  Email string
}

a := &A{Name: "Tony", Phone: "12355"}
b := &B{}

automapper.Map(a, b, func(src *A, dst *B) {
  dst.Email = "hello@world.com"
})
```



* Slice - Slice Mapping

```go
type A struct {
  Name string
}

type B struct {
  ID   int
  Name string
}

a1 := &A{Name: "One"}
a2 := &A{Name: "Two"}

aSlice := []*A{a1, a2}
bSlice := []*B{}

automapper.Map(aSlice, &bSlice, func(idx int, src *A, dst *B) {
  t.Log(idx, src, dst)
  dst.ID = idx + 1
})

```
