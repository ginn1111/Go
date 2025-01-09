# Type embedding
---

## 1. What does types embedding look like?

## 2. What types can be embedded?
  - The defined name type 
  - The alias type if:
    + The underlying type is a `unnamed pointer type`, the base type must not be `interface` or `pointer` type.
  - The `unnamed pointer` type if the base type is not `interface` or `pointer` type.

  > Do not embedding the type and corresponding the pointer type. The not embedding struct type in itself or its alias.

## 3. What is the meaningfulness of type embedding?
  - Like inheritance of OOP

## 4. Does the embedding type obtain fields and methods of the embedded type?
  - Consider the methods are defined for value type or pointer type

## 5. Shorthands of Selectors

## 6. Selector shadowing and colliding
  - The difference when using the private method in the package

## 7. Implicit method for embedding types
  - For defined type `type Z X`. `Z` type obtain the method sets of `embedded types in X` but not the method sets of `X`

## 8. Normalization and evaluation of promoted method values
  - Like method `value evaluation`
  > Note the receiver

## 9. Interface types Embed all kind of types

## 10. An inserting type embedding example
