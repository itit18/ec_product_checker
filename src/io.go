//io.go

package main

import (
  "fmt"
  "os"
)

func main(){
  var stdin string
  //fmt.Scan(&stdin)
  stdin = os.Args[1]

  fmt.Println(stdin)
}