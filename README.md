#  GoScala

利用 Go 1.18 (2022 年推出) 的 Generic，實作 Scala Monoid 功能。目前還在開發、測試、實驗。 API 還會再修改。請不要用在開發與正式環境。

目前已完成：

- Option
- Either
- Try
- Slice


## 編譯環境

請參考: [How to use type parameters for generic programming](https://www.jetbrains.com/help/go/how-to-use-type-parameters-for-generic-programming.html)。

我的環境： Mac OS Big Sur 11.5.2 / Go 1.17.1，安裝 Go 1.18 步驟如下：

1. `cd ~`
1. `git clone https://go.googlesource.com/go goroot`
1. `cd goroot`
1. `git checkout dev.typeparams`
1. `cd src`
1. `./all.bash`

Compile 時，還是需要用到 `-gcflags -G=3`

## 心得

因為版本還在開發中，還是會發生一些我覺得很奇怪的問題。不過比 Go 1.17 與 Go2Go 好一些。

語法上與 Go2Go 有些出入。如

Go2Go:

```go
type Number interface {
    type int, ...
}
```

在 Go 1.18 變成

```go
type Number interface {
    int, ....
}
```

相關資料，請見 [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)