## Numeric Data types in go

In Go, numeric data types are divided into several categories: integers, floating-point numbers, and complex numbers. Here’s an overview of each, along with examples:

## Integer Types

1. _Signed Integers_: Can hold both positive and negative values

- int: Platform-dependent size (either 32 or 64 bits).
- int8: 8-bit signed integer (-128 to 127).
- int16: 16-bit signed integer (-32,768 to 32,767).
- int32: 32-bit signed integer (-2,147,483,648 to 2,147,483,647).
- int64: 64-bit signed integer (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807).

2. _Unsigned Integers_: Can hold only non-negative values

- uint: Platform-dependent size (either 32 or 64 bits).
- uint8: 8-bit unsigned integer (0 to 255). Also known as byte.
- uint16: 16-bit unsigned integer (0 to 65,535).
- uint32: 32-bit unsigned integer (0 to 4,294,967,295).
- uint64: 64-bit unsigned integer (0 to 18,446,744,073,709,551,615).

3. _Alias Types_:

- byte : it is an alias for `unit8`, which means it represents and 8-bit unsigned integer.

- rune : it is an alias for `int32`, which means it represents a 32-bit signed integer.

## Floating Points:

- float32
- float64

## Complex Types:

- complex64 : Complex numbers with float34 and imaginary part

- complex128 : Complex number with float64 and imaginary part
