# Go-Lang Notes

## Three Dots/Variadic Parameter Syntax (`...`)

### 1. Using (`...`) in function parameter

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

- `...int` means `nums` can take any number of `int` arguments.

call `sum` with any numbers of arguments

```go
sum(1, 2, 3, 4)
```

### 2. Using (`...`) when passing a slice to a variadic function

`...` is used to **"unpack"** a slice when passing it to a variadic function.

```go
func main() {
    numbers := []int{1, 2, 3, 4}

    fmt.Println(sum(numbers...))  // Pass the slice with `...` to "spread" it
}
```

### 3. Using (`...`) in array declaration

`...` can also be used to let Go determine the array’s length based on the number of elements provided.

```go
arr := [...]int{1, 2, 3, 4}
```
