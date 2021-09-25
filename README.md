# go-strcase

![gitlab tag (latest by date)](https://img.shields.io/gitlab/v/tag/zephinzer/go-strcase?sort=date)
![pipeline status](https://gitlab.com/zephinzer/go-strcase/badges/master/pipeline.svg)
![coverage report](https://gitlab.com/zephinzer/go-strcase/badges/master/coverage.svg)

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
  - [Constraints/design decisions](#constraintsdesign-decisions)
    - [Character delimiters](#character-delimiters)
    - [Casing delimiters](#casing-delimiters)
    - [Number delimiters](#number-delimiters)
    - [Additional trivia](#additional-trivia)
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

```go
fmt.Printf("camel case: %s", strcase.ToCamel("hello world"))
// output: helloWorld
```

> Useful for: **C-like sematics, JSON, YAML**

### `ToCapital`

Use this if you want your string to `LookLikeThis`

```go
fmt.Printf("capital case: %s", strcase.ToCamel("hello world"))
// output: HelloWorld
```

> Useful for: **C-like sematics**

### `ToLowerKebab`

Use this if you want your string to `look-like-this`

```go
fmt.Printf("lower kebab case: %s", strcase.ToCamel("hello world"))
// output: hello-world
```

> Useful for: **JSON, YAML**

### `ToLowerSnake`

Use this if you want your string to `look_like_this`

```go
fmt.Printf("lower snake case: %s", strcase.ToCamel("hello world"))
// output: hello_world
```

> Useful for: **Python, Perl, Shell**, JSON, YAML

### `ToUpperKebab`

Use this if you want your string to `LOOK-LIKE-THIS`

```go
fmt.Printf("upper kebab case: %s", strcase.ToCamel("hello world"))
// output: HELLO-WORLD
```

> Useful for: **?**

### `ToUpperSnake`

Use this if you want your string to `LOOK_LIKE_THIS`

```go
fmt.Printf("upper snake case: %s", strcase.ToCamel("hello world"))
// output: HELLO_WORLD
```

> Useful for: **Shell**, JSON, YAML

## Constraints/design decisions

### Character delimiters

- Only dashes, underscores, spaces are considered as character delimtiers (eg. `a-pizza`, `a_pizza`, and `a pizza` are seens as `a` and `pizza`)
- Character delimitations take precedence over the other delimiters (eg. `i want 1 Pizza` is `[i, want, 1, Pizza]`, but `iWant1Pizza` is `[i, want1, pizza]`)

### Casing delimiters

- An upper-cased character is considered a delimiter only when the character after that is a lower-cased character (eg. `APizza` is seen as `A` and `Pizza`, but `sendAPIRequest` is seen as `[send, api, request]`)

### Number delimiters

- The end of any series of numbers is considered a delimiter. This decision was made since variable names usually cannot begin with a number (eg. `1pizza` is `[1, pizza]`, but `pizza1` is `[pizza1]`)

### Additional trivia

- After writing the above, the world pizza doesn't hold any more meaning in my mind. Send help.

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
