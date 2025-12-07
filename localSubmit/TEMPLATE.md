FOLDER Structure

Current
.
|-- TEMPLATE.md
|-- go.mod
|-- inOut  (only use input.txt, output.txt files only)
|   |-- inOut.go  // not in use
|   |-- input.txt
|   `-- output.txt
`-- list             //dailySubmissions
|   |-- a271125lc.go
|   |-- generic.go   // has only IOES func
|-- snippet
    |-- snippet.go   // single generic code for every where (contest,lc,..etc)
    |-- generic.go   // for IOES (if req put this IOES func in snippet.go also)


Below is VOID (not in use)
.
|-- inOut (module)
|   |-- go.mod
|   |-- inOut.go
|   |-- input.txt
|   `-- output.txt
`-- list (main)
    |-- generic.go (always run this file)
    `-- a271125lc.go (only functions based on questions)



TODO:
1. make fully read write from file (if file not found then create the file):DONE
2. make a script that makes single main.go code :not need 
3.





Learning
module is collection of multiple packages
packages = directories
we import packages ,not modules
Go resolves imports via modules (or GOPATH)
A module = a project root.
A module is just a folder that Go recognizes as a project
go.mod tells Go the name of your project, This name becomes the import path for subfolders.
Modules allow Go to find your code anywhere

Without a module:
    Go does not know where your code lives
    You can't import your own folders
    You can't resolve dependencies
    You get errors like “cannot find main module”
Modules fix this.

Modules manage your dependencies
    And will download it into: $GOMODCACHE

Modules allow you to import your own other modules.


what this means:
Go resolves imports via modules (or GOPATH). If b isn’t in a module, it has no import path, so it can’t import amod normally.



WAYS of running code

1. go run ./list/*.go  => compiles all .go files
