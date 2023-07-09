# Infix Calulator
## Requires
* Without using C++ libraries for containers and algorithms, including \<array\>, \<list\>, \<map\>, \<queue\>, \<set\>, \<stack\>, \<vector\>, \<algorithm\>, \<tuple\>
******
## Input
### Input specifications
EBNF
* Expression = Expression BinaryOp Expression | "(" Expression ")" | Numerical
* Numerical = Integer | Floating
* Integer = Digits
* Floating = Digits "." Digits | "." Digits | Digits "."
* Digits = Digit Digits | Digit
* Digit = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
* BinaryOp = "+" | "-" | "*" | "/"
* All whitespace (space or tab) should be ignored.
*Integers will not exceed the limit of [-2147483648,2147483647].


## Output
### Output specifications
Output the result in one line. For example, if the input were to be (3 + 4) * 8 - 2, the output would be 54.
