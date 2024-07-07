# Variables in Go

- Variables name must begin with letter or an underscore(_). And the names may contain letter 'a-z' or 'A-Z' or digit 0-9 or _.

- A variable name should not start with an digit

- the name of the variable is case sensitive

## Declaring Variable

- In Go language, variables are created using var keyword of a particular type, connected with name and provide its initial value.

```
var variable_name type = expression
```

1. **Short variable declaration** : this is used inside a function to declare and initialize variable
   syntax :

```
variable_name := value
```

2. **Multiple Variable Declaration** : multiple variable can be declared in a single statement using var keyword but the type of the variable should be same

```
var variable1,variable2 type
```

3. **Grouped Variable Declaration** : multiple variable can be declared in a block using the 'var' keyword

```
var (
    variableName1 type
    variableName2 type
)
```

4. **Declaration with Initializer** : When u declare a variable with an initial value go can infer the type

```
var variableName = value
```

5. **Constant Variable** : Constant are created using const keyword and cannot be changed after declaration.

```
const Const_name type = value
```
