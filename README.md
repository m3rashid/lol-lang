## Lol Lang

Lol lang is an assembly style interpreted language with its syntax inspired from assembly

- Currently works only with integers
- supports the following features
- Refer the `examples` folder for code examples

| Feature   | Description                                       |
| --------- | ------------------------------------------------- |
| PUSH      | Push to the stack                                 |
| POP       | Pop from the stack                                |
| ADD       | add last 2 values from the stack                  |
| SUB       | subtracts last 2 values from the stack            |
| PRINT     | prints things to the screen                       |
| READ      | read input from the user                          |
| JUMP.EQ.0 | jump if top of the stack is zero                  |
| JUMP.GT.0 | jump if the top of the stack is greater than zero |
| HALT      | marks the end of the program                      |

## Usage

```bash
# build the project
make build

# run the project
bin/lol <file_name>.lol
```
