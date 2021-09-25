# go-strcase

A string case converter for use in Golang applications.

- [go-strcase](#go-strcase)
- [Usage](#usage)
  - [Importing](#importing)
  - [Methods](#methods)
    - [`ToCamel`](#tocamel)
    - [`ToCapital`](#tocapital)
    - [`ToLowerKebab`](#tolowerkebab)
    - [`ToLowerSnake`](#tolowersnake)
    - [`ToUpperKebab`](#toupperkebab)
    - [`ToUpperSnake`](#touppersnake)
- [Code design (how it works)](#code-design-how-it-works)
  - [Requirements](#requirements)
  - [Tokenization](#tokenization)
  - [Reconciliation](#reconciliation)
- [Contributing](#contributing)
- [License](#license)

# Usage

## Importing

```go
import "github.com/zephinzer/go-strcase"
```

## Methods

### `ToCamel`

Use this if you want your string to `lookLikeThisFormat`

> Useful for: **C-like sematics, JSON, YAML**

### `ToCapital`

Use this if you want your string to `LookLikeThis`

> Useful for: **C-like sematics**

### `ToLowerKebab`

Use this if you want your string to `look-like-this`

> Useful for: **JSON, YAML**

### `ToLowerSnake`

Use this if you want your string to `look_like_this`

> Useful for: **Python, Perl, Shell**, JSON, YAML

### `ToUpperKebab`

Use this if you want your string to `LOOK-LIKE-THIS`

> Useful for: **?**

### `ToUpperSnake`

Use this if you want your string to `LOOK_LIKE_THIS`

> Useful for: **Shell**, JSON, YAML

# Code design (how it works)

## Requirements

This package was created mainly for usage within a tool that needs to handle cross-runtime naming conventions such as between `SHELL_VARIABLES` and `javascriptVars` or `python_vars`

## Tokenization

This package works by first tokenizing the input string into separate words first before applying any transformations. This looks like:

```
HelloWorld => [hello, world]
helloWorld => [hello, world]
10PizzasPlease => [10, pizzas, please]
1-for-1 => [1, for, 1]
1for1 => [1, for1]
```

## Reconciliation

After the above tokenization, the transformations are added:

```
[hello, world] =(ToCamel)=> helloWorld
[10, pizzas, please] =(ToCamel)=> 10PizzasPlease
[1, for, 1] =(ToCamel)=> 1For1
[1, for1] =(ToCamel)=> 1For1
```

# Contributing

The working repository is at Gitlab at [https://gitlab.com/zephinzer/go-strcase](https://gitlab.com/zephinzer/go-strcase), if you are seeing this on Github, it's just for SEO since y'know all the cool new kids are on Github 😂

# License

Use this anywhere you need to. Licensed under [the MIT license](./LICENSE)
