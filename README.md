# Kata Machine

## Generating files

Run `make gen` to generate a basic skeleton of the implementation files the tests will be ran against. A build will fail after generating because
some of the genric functions are not returning the correct thing.

For those functions you can use the following code to get pass the building errors, change the type being returned if needed:
```go
var ret T
return ret
```

## Testing

Run `make test` for all tests and `make test TEST=TestArrayList` to specify a specifc test name to run

