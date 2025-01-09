# Struct
---
## 1. Struct type and struct type literal

## 2. Struct value literals and struct value manipulations

## 3. About struct value assignment
  - The requirement to the struct type can assignment to another one:
    + Both struct type is identical
    + One of both can implicit convert to the other one (the underlying type is identical and at least one is `unnamed type`)
    > Unnamed type is usually the composite type

## 4. Struct field addressability
  - Composite literal is unaddressable
  ```go
  type Book struct{
    page int
    content string
  }

   Book{}.page = 10 // Book{} is unaddressable -> fail to compile
    
```

## 5. Composite literal are unaddressable but can take addresses
  ```go
  type Book struct{
    page int
    content string
  }

  p := &Book{} // its ok
  ```
## 6. In field selectors, dereferences of receivers can be implicit
  ```go
  type Book struct{
    page int
    content string
  }

  p := &Book{} 
  p.page = 10 // equivalent to (*p).page = 10
  ```

## 7. About struct comparison
  - The comparable struct is the fields of the struct must be comparable 
  - The conditions:
    + Both can assign to each other and are comparable
    + The field's order of both structs must be the same
  - How? Comparing the corresponding fields of both structs and stop in advance when the first difference is found or a `panic` occurs

## 8. About struct value conversion
  - Both structs have the same underlying type -> `implicit conversion`

## 9. Anonymous structs type can be used in field declaration

