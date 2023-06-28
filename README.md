# skiplist

`skiplist` implementation in go

```go
l := skiplist.New[int, string]()

l.Insert(1, "10")
l.Insert(1, "20")
l.Insert(1, "30")

v, ok := l.Get(3) // "30", true

l.Insert(3, "90") // updates
v, ok = l.Get(3) // "90", true

ok = l.Delete(3) // true

length := l.Length() // 2
```
