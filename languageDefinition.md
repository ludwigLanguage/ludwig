# Ludwig Language Description
## Expressions:
Expressions are any complete statement in ludwig followed by a semicolon or a newline
```
<expr>; <expr>;...
<expr>
<expr>...
```

## Primitive Data Types:
Primitive data types include booleans, numbers, and strings. All basic opperators can be used with these types.

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
Empty Value:
```
nil
```

## Lists:
Lists are an array of values that can be accesed using standard slice syntax like you might see in python or go.
```
[<item>, <item>...]
Examples: [1, 2, 3]; [1, 1, 2, 3, 5, 8, 13]

<list>[<number>] #Returns value
[1, 2, 3][0]     #Returns the number one
```
Slicing:
```
<list>[<start>:<end>] OR <list>[<start>:] OR <list>[:<end>]
```

## Declarations:
Declarations are a way to assign an identifier to a value
### Note:
Declarations will return the value the compute on the right side (see example)
```
<identifier> = <expression>
```
Examples: 
```
pi = 3.14
name = "John"
x = y = 10 ## Both x and y will be assigned the value 10
```
## Block Expression:
### Note:
The last expression evaluated in the block will become the return value of the block
### Scoped Compound Expression:
```
{ 
    <expr> 
    <expr> 
}
```
Example:
```
c = {
    a = 10
    b = 3
    a + b
}

println(c) #prints 13
```
Note: Any variables created will NOT be available outside the block expression

### Un-scoped Block Expression:
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
Note: Any variables created will be available outside the block expression


## If - Else Expressions:
The standard control-flow feature. Except the return a value!!!!
```
if <expr> <expr> else <expr>
```
Example:
```
someValue = 10

division_statement =
    if someValue == (100/10)
        "Division Works!"
    else
        "Division Does Not Work :("

println(division_statement) ##Prints out 'Division Works!'
```

## While Loops:
Another standard control flow features, except it returns a list of the values evaluated
```
while <expr> <expr>
```
Example:
```
i = 0
listOfNumsUnder11 =
    while i < 10 {
        i = i + 1 
    }

println(listOfNumsUnder11) ##[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
```

## For Loops:
For looping through lists, this feature also returns a list containing the values evaluated.
```
for <ident>, <ident> in <list> <expr>
```
Example:
```
list_of_outputs =
    for location, iter in numbers {
        println(iter, "has location:", location, "in the list")
    }
```
### Notes:
The first identifier will be assigned the current index in the list that we are looping over. The second identifier will be assigned the value of that index in the list

## Functions:
Function do not contain their identifier in there intial declaration. One must use a declaration, like the ones displayed in the previous section, to name functions
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
2) Every function will have the __recurse()__ function, which can be used for tail recursion (see example).

## Built In Functions
1) `type_of(<expr>)`              
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
7) `println(<string>)` 
8) `print(<string>)`
9) `read(<string>, <string>)`
    prints first string and reads until second string is seen
10) `type_check(<type_id>, <value>)`
    panics if the type of the value does not match the type id.
    The available type ids include:
    Numbers:    "_num"
	Strings:    "_str"
	Booleans:   "_bool"
	Nil:        "_nil"
	Lists:      "_list"
	Functions:  "_func"
	Structures: "_struct"
	Objects:    "_object"
	Builtins:   "_builtin"
	Type IDs:   "_type"

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






























