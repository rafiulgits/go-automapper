# Go AutoMapper

*This repository is an extension or fork edition of [stroiman/go-automapper](https://github.com/stroiman/go-automapper)*

Go-automapper can automatically map data between different types with identically named fields in GO. Useful for initializing DTO types from build in data.



**Extension**

* This package only use `MapLoose` function of [stroiman/go-automapper](https://github.com/stroiman/go-automapper) to avoid any kind failure when mapping into destination type.
* This package allows to pass a callback function while invoking `Map` function to map unknown fields of destination object



***

### How to use

Install automapper in your project

```
go get github.com/rafiulgits/automapper
```



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

