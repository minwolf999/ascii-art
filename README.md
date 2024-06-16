# ASCII-ART

**Ascii-art** is a Golang program which consists in receiving a `string` as an argument and outputting the `string` in a graphic representation using ASCII.

What we mean by a graphic representation using ASCII, is to write the `string` received using ASCII characters, as you can see in the example below:

```
 ______
|___  /                                ___    _
   / /    ___    _ __     ___         / _ \  / |
  / /    / _ \  | '_ \   / _ \       | | | | | |
 / /__  | (_) | | | | | |  __/       | |_| | | |
/_____|  \___/  |_| |_|  \___|        \___/  |_|


```
1. This program is based on the [ASCII table](https://wikipedia.org/wiki/American_Standard_Code_for_Information_Interchange):

2. This program can handle an input with:
   -   `numbers`
   -   `letters`
   -   `spaces`
   -   `special characters`
   -   and `\n`

## How to run it ?
To run Ascii-Art :
   * Install or Update the latest version of [Go](https://go.dev/doc/install)
   * Clone the repo (Make sure you have the required authorizations)
   * Open the folder in an IDE
   * Open your terminal
   * Copy one of these commands followed by your input (myInput here):
```
$ go run . "myInput"
$ go run ./main.go "myInput"
$ go run main.go myInput 
```
   * If you want to see all flags possible copy one of these commands:
```
$ go run .
$ go run . -h
$ go run . --help
```

### Result :
```
$
                     _____                           _    
                    |_   _|                         | |
 _ __ ___    _   _    | |    _ __    _ __    _   _  | |_
| '_ ` _ \  | | | |   | |   | '_ \  | '_ \  | | | | | __|
| | | | | | | |_| |  _| |_  | | | | | |_) | | |_| | \ |_
|_| |_| |_|  \__, | |_____| |_| |_| | .__/   \__,_|  \__|
             __/ /                  | |
            |___/                   |_|
           
$
```
