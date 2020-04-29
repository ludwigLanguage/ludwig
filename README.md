# Ludwig
Ludwig Programming Language
The ludwig programming language is a fun little project I started in 2018. The goal was to create and interpreter for a programming langauge in which everything is an expression.

### Note: This implementation of Ludwig is still under heavy development. As a result, this interpreter may be buggy or hard to install. Note that at this point, there are very few contibutors to this project. Having fewer eyes on the project results in messier and more inconsistant code!

## Installation
Dependacies: In order to install this program at this point, you must have the compiler for golang installed.

Linux & MacOs
```
$ mkdir ~/go
$ cd go
$ mkdir src
$ cd src
$ git clone https://github.com/ludwigLanguage/ludwig.git
$ cd ./ludwig/bin
$ bash ./install.sh
```

## Usage
```
ludwig -e <filename>.ldg #executes the the file
```

## Hello World Program
```
println("Hello, World!")
```
A more advanced version of a greeter program might run as such:
```
mkGreeting = func(greet) {
	func(name) {
		print(greet + ", " + name) #Newline given at the end of read() call
	}
}

sayHello = mkGreeting("Hello")
name = read("Enter your name: ", "\n")
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
