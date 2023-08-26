## `goroutine` 辅助包

### 重要提醒

当前包目录下，不要存放其它文件，否则，可能在交叉编译时出现问题。

```go
# runtime/goroutine
./runtime/goroutine/goid.s:6: unrecognized instruction "get_tls"
./runtime/goroutine/goid.s:7: unrecognized instruction "MOVQ"
./runtime/goroutine/goid.s:8: unrecognized instruction "MOVQ"
./runtime/goroutine/goid.s:9: unrecognized instruction "LEAQ"
./runtime/goroutine/goid.s:10: unrecognized instruction "MOVQ"
./runtime/goroutine/goid.s:11: unrecognized instruction "MOVQ"
asm: assembly of ./runtime/goroutine/goid.s failed
```

### 参考

- [Go 高级编程：Goroutine ID](https://github.com/chai2010/advanced-go-programming-book/blob/master/ch3-asm/ch3-08-goroutine-id.md)
- [曹春晖 goroutineid 类库](https://github.com/cch123/goroutineid)
- [曹春晖 plan9 assembly 完全解析](https://github.com/cch123/golang-notes/blob/master/assembly.md)
- [曹春晖 汇编分享](https://github.com/cch123/asmshare)