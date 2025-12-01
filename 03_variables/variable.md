# Variables vs Constants - Unused Declaration Behavior

## Key Difference

In Go:
- **Unused local variables** → Compilation ERROR 
- **Unused constants** → No error 
- **Unused package-level variables** → No error ✅
- **Unused imports** → Compilation ERROR 

## Why this difference?

**Constants are compile-time values** with zero runtime cost. They're often used for:
1. Configuration values
2. Documentation/reference
3. API definitions where not all constants may be used in every file

**Variables allocate memory** at runtime, so unused local variables are considered wasteful and likely a mistake.

## Example

```go
package main

func main() {
    //  ERROR: unused variable
    // name := "harsh"
    
    //  NO ERROR: unused constant
    const port = 5000
    const host = "localhost"
}
```

If you uncomment `name := "harsh"`, you'll get:
```
./constants.go:X:X: name declared and not used
```

But the constants won't complain! This is intentional design by the Go team to make constants more flexible for configuration and API design scenarios.

## Pro Tip

If you want to temporarily "use" a variable to avoid the error during development, you can assign it to the blank identifier:
```go
name := "harsh"
_ = name  // This satisfies the compiler
```

