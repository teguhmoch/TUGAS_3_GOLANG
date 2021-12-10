package main

import "fmt"

func main() {

  var buah =[]string{"apel","jeruk","anggur","melon"}
  buah = append(buah,"pepaya")
  var jumlah =len(buah)
  var i=0
  fmt.Println("Jumlah element :",jumlah)
  fmt.Println("Isi element    :",buah)
  for i<jumlah{
    fmt.Println("element ke - :",i,buah[i])
    i++
  }
}
