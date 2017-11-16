//file.go

package main 
import (
  "fmt"
  "io/ioutil"
  "os"
  "bufio"
)

func main(){
  //まとめて読み込み
  data, err := ioutil.ReadFile("resource/test.txt")
  if err != nil {
    fmt.Print("error!")
  }
  fmt.Print(string(data) + "\n")

  //個別読み込み
  file, err := os.Open("resource/test.txt")
  if err != nil{
    fmt.Print("os open error!")
  }
  defer file.Close()

  scan := bufio.NewScanner(file)
  for i := 1; scan.Scan(); i++ {
    if err := scan.Err(); err != nil {
      fmt.Print("scan error!")
      break
    }
    fmt.Printf("%4d行: %s\n", i, scan.Text())
  }


}
