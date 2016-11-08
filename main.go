package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

type instruction byte
type data struct {
  tape []byte
  i    int
}

const (
  inc    instruction = '+'
  dec    instruction = '-'
  incPtr instruction = '>'
  decPtr instruction = '<'
  read   instruction = ','
  write  instruction = '.'
  jmpZ   instruction = '['
  jmpNZ  instruction = ']'
)

func execute(bs []byte) {
  tape, i := make([]byte, 1), 0
  for p := 0; p < len(bs); p++ {
    switch instruction(bs[p]) {
    case inc:
      tape[i]++
    case dec:
      tape[i]--
    case incPtr:
      i++
      if len(tape) <= i {
        tape = append(tape, 0)
      }
    case decPtr:
      i--
    case read:
      fmt.Scanf("%c", &tape[i])
    case write:
      fmt.Print(string(tape[i]))
    case jmpZ:
      if tape[i] == 0 {
        for d := 1; d > 0; {
          p++
          switch instruction(bs[p]) {
          case jmpZ:
            d++
          case jmpNZ:
            d--
          }
        }
      }
    case jmpNZ:
      if tape[i] != 0 {
        for d := 1; d > 0; {
          p--
          switch instruction(bs[p]) {
          case jmpZ:
            d--
          case jmpNZ:
            d++
          }
        }
      }
    }
  }
}

func main() {
  var name string
  if len(os.Args) == 2 {
    name = os.Args[1]
  }

  bs, err := ioutil.ReadFile(name)
  if err != nil {
    panic(err)
  }
  execute(bs)
}
