# Code packages and package imports
---

## 1. Introduction of package import

## 2. More about `fmt.Printf` format verbs
  - There are some format verbs:
    + `%v` general string of corresponding arguments
    + `%T` type of the arguments
    + `%x` the hex string
    + `%s` string of corresponding arguments, the arguments must be string or byte slice
     `%%` percent sign

## 3. Package folder, package import path and package dependencies
  - The special `.../path/to/internal/` folder.
  - `main` package is an entry point program (only one).

## 4. The `init` function
  - Do not have any arguments or return results.
  - Can re-define
  - `init` is a keyword

## 5. Resource initialization

## 6. Full package import forms

## 7. Each non-anonymous package must use at least one
  - Except the `blank` import (just using the effect of `init` function of imported packages)

## 8. Module
 - The collection of packages


