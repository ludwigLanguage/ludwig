# Ludwig Language Description
## Expressions:
### Expressions are any complete statement in ludwig followed by a semicolon or a newline
```
<expr>; <expr>;...
<expr>
<expr>...
```

## Primitive Data Types:
### Primitive data types include booleans, numbers, and strings. All basic opperators can be used with these types.

Boolean:
```
true
false
```
Numbers:
```
<number> OR <number>.<number>
Examples: 10; 3.14; 21;
```
Strings:
```
"<text>" OR '<text>'
Examples: "Hello, World"; 'Hello, World'
```

## Lists:
### Lists are an array of values that can be accesed using standard slice syntax like you might see in python or go.
```
[<item>, <item>...]
Examples: [1, 2, 3]; [1, 1, 2, 3, 5, 8, 13]

<list>[<number>] #Returns value
[1, 2, 3][0]     #Returns the number one
```

## Declarations:
### Declarations are a way to assign an identifier to a value
```
<identifier> = <expression>
Examples: pi = 3.14
         name = "John"
```

## Functions:
### Function do not contain their identifier in there intial declaration. One must use a declaration, like the ones displayed in the previous section, to name functions
Function expression:
```
    func (<identifier>, <identifier>) <expr>`
```
Function Call:
```
    `<function>(<arguments>)`
```
Example:
```
add = func(x, y)
    x + y

ten = add(5, 5)


foreverPrint = func(value) {
    println(value)
    recurse(value)
}

foreverPrint("Hello, World")
```
### Notes:
1) A function returns the value of the expression immediately following the
arguments list<br/>
2) Every function will have the __recurse()__ function, which can be used for tail recursion

## Built In Functions
1) `typeOf(<expr>)`              
    Returns the type of a given value
2) `str(<expr>)`   
    Converts the given expression into a string
3) `num(<string>)`   
    Converts the given string into a number
4) `len(<expr>)`  
    Returns the length of the given string or list
5) `system(<bool>, <string>...)`  
    Executes the given command, and then prints the output if the boolean given is 'true'
    Returns an object containing the 'Error' and the 'Output' variables which contain the
    stdout and stderr values from that command as strings
6) `import(<string>)`
    Returns an object produced from the given file.


## Compound Expression:
### Scoped Compound Expression:
```
{ 
    <expr> 
    <expr> 
}
```
Example:
```
{
    a = 10
    b = 3
    a + b
}
```
Note: Any variables created will NOT be available outside the compound expression

### Un-scoped Compound Expression:
```
do <expr> <expr> end
```
Example:
```
do
    a = 10
    b = 2
    a + b
end
```
Note: Any variables created will be available outside the compound expression

## Structs:
```
struct <expr>
```

Example:
```
person = struct {
    __init__ = func(name, age) {
        self.name = name
        self.age = age
    }

    printInfo = func() {
        println(age, " ", name)
    }
}
john = person("John Doe", 22)
```

### Notes:
1) The '__init\__' function will be executed upon the initialization of an object, with the arguments given to the object in the initialization call<br/>
2) The '__self__' object will be present within the environment of every struct so that users can change the values of variables contained within the object<br/>
3) All code within the struct will be executed when the struct is created






























