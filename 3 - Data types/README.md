# Integer

Integer type can be int8, int16, int32 and int64. The number at the end represents the number of bytes allocated to represent the number.
When creating variables using only `int` or with inferred type will cause the compiler to use the processor architecture as reference, i.e. x64 processor with auto create a int64.
`uint` can be used for unsigned integer numbers (only positive).

Golang supports alias for some types. 
- int32 = rune
- uint8 = byte
Example: var numberEx rune = 12345

# Float

float32 and float64

# String

Created using `string` type or charactes sequence between double quotes.
There is no `char` type in go. A single character between single quotes is converted to its number on ASCII table.

# Boolean

`bool`
`false` by default.

# Error

Type `error` exists to represent errors. Default value is `<nil>`.
Errors can be created using the native package `errors`, example: `var err error = errors.New("message goes here")`