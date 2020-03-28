# Ludwig
Ludwig Programming Language
The ludwig programming language is a fun little project I started in 2018. The goal was to create and interpreter for a programming langauge in which everything is an expression.

## Installation
Dependacies: In order to install this program at this point, you must have the compiler for golang installed.

Linux & MacOs
```
$ cd ./ludwig/bin
$ bash ./install.sh
```

## Usage
```
ludwig -e <filename>.kgo #executes the the file
```

## Hello World Program
```
println("Hello, World!")
```
A more advanced version of a greeter program might run as such:
```
mkGreeter = func(prefix) {
  func(name) {
    println(prefix, ", ", name)
  }
}

sayHello = mkGreeter("Hello")
name = readln("Enter your name: ")
sayHello(name)
```

## License
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.