## Prerequisite
Before start, some preparation are need.
- Run `make` to generate example configration file in `config/config.json`
- An example config.json is like below

```json
{
    "SymK": 128,
    "SymByte": 16,
    "PubK": 256,
    "Debug": false,
    "Root": "<Set Project Root Path Here>"
}
```

Note: for vscode user, set environment variable in workspace setting before using the go langugae plugin.
## Quick Start 

As the first taste of the project, we recommand to start with
```shell
> make test
```

This `Makefile` target launchs the test of the project and it is a good probe for you to make sure you are not downloading a fake project. XD

Then, everything going fine. Just RTFSC in `Makefile`.

example:
```go
package main
// max of 4 values, two from party 0. two from party 1.
func main() {
    var in_0 []sint8 =  i8n(i,2)
    var in_1 []sint8 =  i8n(i,2)
    var max sint8;
    if in_0[0] > in_0[1]{
        max = in_0[0]
    }
    if in_1[0] > max {
        max = in_1[0]
    }
    if in_1[1] > max {
        max = in_1[1]
    }
    var pub_max int8 = openi8(max)
    fmt.println(pub_max)
}

```

single private input from party i
```go
var in sint8 = i8(i)
```

private input array with length n from party i 
```go
var in []sint8 = i8n(i, n)
```
decl private var with init 0 
```go
var d sint8 = 0
```
decl private array with length n 
```go 
var c []sint8 = make([]sint8, n)
```
convert private to public
```go 
var p int8 = openi8(d)
```
for loop in go 
```go 
for (i :=0; i < m;i++){

}
```


