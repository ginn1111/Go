# Reflection

---

## 1. Overview of go reflection

## 2. The `reflect.Type` Type and Value

- `reflect.Type` is a type of `reflect.TypeOf` (at runtime) or `reflect.TypeFor` (at compile time). The two functions can be used for interface or non-interface type.
- Some methods of `reflect.Type` type just suitable for specific types, otherwise, the `panic` will occurs.
- With the `struct` type, the method `NumMethod` just count the exported method, but the `NumField` count all field of `struct` type
- The above behavior do not apply for non-exported method interface type.
- Can use `reflect.StructTag` with two functions: `Lookup` and `Get`
- There are some limitations:
  - Do not defined `interface`.
  - Do not use `embedding` field.
  - Do not define new type

## 3. The `reflect.Value` Type and Value
  - The same as `reflect.Type`, `reflect.ValueOf` function will return the `reflect.Value` type.
  - The represented by a `reflect.Value` of value `v` often called the underlying value of `v`
  - The suitable methods the same as `reflect.Type`'s methods
  - The `CanSet` function will check whether the value can be set. If it's ok, we will use `Set` function to set the value.
  - The `non-exported` field must not modify. If modify, the `panic` will occurs.
  - Using `reflect.Indirect(reflect.Value)` the same as `reflect.Value.Elem()`. If the `reflect.Value` do not value be represented of pointer type, will return the copy of value.
  - The `reflect` standard package declare some functions the same as builtin function or non-reflection functionality:
    + Slice type: `reflect.Value.Len()`, `reflect.Value.MakeSlice(reflect.Type, int, int)`, `reflect.Index()`, `reflect.Append(reflect.Value, ...reflect.Value)`.
    + Function type: `reflect.MakeFunc(reflect.Type, function([]reflect.Value) []reflect.Value)`
    + Map type: `reflect.MapRange()`, `reflect.SetMapIndex(reflect.Value, reflect.Value)`
    + Channel type: `reflect.Send(reflect.Value)`, `reflect.TrySend(reflect.Value)`, `reflect.Trycv()`
      > `TrySend` and `Trycv` is the same as `select-case` with one-case-one-default
    + `select-case`, `reflect.SelectCase([]struct{Dir: reflect.Select<Default|Send|Recv>, Chan: reflect.Value, Send: reflect.Value}`)`, 
    + Type assertion: `reflect.Interface().(type)`,
      > Do not call `Interface` with the non-exported fields of struct type
    + Convert `slice` to `pointer array`: `reflect.ConvertibleTo(reflect.Type)`, return `false` if the len of pointer array is greater than len of slice
  - If the `reflect.Value` is function underlying value, using `Call` to call the function
    > The arguments of `Call` must not the non-exported struct fields, if not the panic will occurs

