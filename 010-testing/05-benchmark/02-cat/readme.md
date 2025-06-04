# Remember to BET
- Benchmark
- Example
- Test

```
BenchmarkCat(b *testing.B)
ExampleJoin()
TestCat(t *testing.T)
```

# Commands
```
godoc -http=:8080  
go test 
go test -bench . 
go test -cover
go test -coverprofile c.out
go tool cover -html=c.out
```