# Marshaling and Unmarshalling JSON
---

## 1. The usage
  - The `encoding/json` that standard package use to encode/decode from unstructed/structed data to JSON

## 2. Unmarshalling (parsing) JSON 
  - Can convert from the JSON data to the unstructed/structed data type in do.
  - Using `json.Unmarshal([]byte, &target) error`
  - There are two types to unmarshalling to:
    + Structed data type: `primitive data` type, `slice`, `struct`, `time`, `map`
    + Unstructed data type: always using `map[string]any` type to `unmarshalling` to
  - Custom `Unmarshal`, implement `UnmarshalJSON([]byte) error`

## 3. Marshalling (stringify) JSON
  - `json.Marshal(any) ([]byte, error)`
  - Using tag field `json:"-"` to omit marshalling this field
  - Using tag field `json:"omitempty"` to omit empty field value
  - Custom `Marshal`, implement `MarshalJSON() ([]byte, error)`

## 4. Pretty JSON using `json.MarshalIndent([]byte, prefix string, indent string)`

## 5. Check valid JSON using `json.Valid([]byte) bool`

## 6. Validate JSON using `validator` pacakge

## 7. `json.Decoder` vs `json.Unmarshal`

## 8. `json.Encoder` vs `json.Marshal`
