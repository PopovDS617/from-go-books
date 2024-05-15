**primitives**

```go

    var x int   // 0
	var px *int // nil

```

**structs**

```go
	var p Person   // {Age: 0}

	p.PrintAge()   // 0

	var pp *Person // nil

	pp.Age        // panic: runtime error: invalid memory address or nil pointer dereference
	pp.PrintAge() // panic: runtime error: invalid memory address or nil pointer dereference

```

**maps**

```go
    var emap map[int]int // map[]

    val := emap[1] // 0

    emap[1] = 10 // panic: assignment to entry in nil map
```

**arrays and slices**

```go
    var arr [1]int // len 1 cap 1

    var sli []int // []

    var psli *[]int // nil

    psli[0] // cannot index psli (variable of type *[]int)

    sli[0] // panic: runtime error: index out of range [0] with length 0

    sli2 := make([]int, 1, 1) // len 1 cap 1

    sli3 := make([]int, 0, 1) // len 0 cap 1
```

**chans**

```go
    var ch1 chan int // nil

    ch1 = make(chan int) // len 0 cap 0

    ch1 <- 9 // deadlock

    ch2 := make(chan int, 10) // len 0 cap 10

    ch2 <- 10 // len 1 cap 10
```
