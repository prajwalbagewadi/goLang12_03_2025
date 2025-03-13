/*
package main
Every program must start with the package declaration. In Go language, packages are used to organize and reuse the code. In Go, there are two types of program available one is executable and another one is the library.
*/
package main

/*
import(
"fmt"
)
Here, import keyword is used to import packages in your program and fmt package is used to implement formatted Input/Output with functions.
*/
import (
	"fmt"
)

/*
func: It is used to create a function in Go language.
main: It is the main function in Go language, which doesn’t contain the parameter, doesn’t return anything, and call when you execute your program.
Println(): This method is present in fmt package and it is used to display “!… Hello World …!” string. Here, Println means Print line.

*/
func main(){
    var a,b int = 5,7
	fmt.Println("Hello World.")
    fmt.Printf("formatted str output:%d",a)
    fmt.Print("print without newline:",b)
}


/*
Step 1:  Create a program main.go and initialize the go.mod  file
go mod init gfg

Step 2: Build the program
go build

Run the executable, On Unix systems
./gfg

windows
gfg.exe
*/

/*
output :
PS C:\Users\bagew\Desktop\WebDev\GoLang1> cd .\helloWorld\
PS C:\Users\bagew\Desktop\WebDev\GoLang1\helloWorld> go mod init prog1
go: creating new go.mod: module prog1
go: to add module requirements and sums:
        go mod tidy
PS C:\Users\bagew\Desktop\WebDev\GoLang1\helloWorld> build go
build : The term 'build' is not recognized as the name of a cmdlet, function, script 
file, or operable program. Check the spelling of the name, or if a path was included, 
verify that the path is correct and try again.
At line:1 char:1
+ build go
+ ~~~~~
    + CategoryInfo          : ObjectNotFound: (build:String) [], CommandNotFoundExcep 
   tion
    + FullyQualifiedErrorId : CommandNotFoundException
 
PS C:\Users\bagew\Desktop\WebDev\GoLang1\helloWorld> go build
PS C:\Users\bagew\Desktop\WebDev\GoLang1\helloWorld> prog1.exe
prog1.exe : The term 'prog1.exe' is not recognized as the name of a cmdlet, function, 
script file, or operable program. Check the spelling of the name, or if a path was     
included, verify that the path is correct and try again.
At line:1 char:1
+ prog1.exe
+ ~~~~~~~~~
    + CategoryInfo          : ObjectNotFound: (prog1.exe:String) [], CommandNotFoundE  
   xception
    + FullyQualifiedErrorId : CommandNotFoundException


Suggestion [3,General]: The command prog1.exe was not found, but does exist in the current location. Windows PowerShell does not load commands from the current location by default. If you trust this command, instead type: ".\prog1.exe". See "get-help about_Command_Precedence" for more details.
PS C:\Users\bagew\Desktop\WebDev\GoLang1\helloWorld> .\prog1.exe
Hello World.
*/