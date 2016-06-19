package main

import "fmt"

func main(){

  if 6%2 ==0{
    fmt.Println("Number is Even");
  }else{
    fmt.Println("Number is Odd");
  }

  if num := 15; num > 0 && num == 15{
    fmt.Println(num, " is greater than 0 and equals", num);
  }else if num < 10 {
    fmt.Println(num, " is less than 10");
  }else {
    fmt.Println(num, "Number is has multiple Digits ");
  }
}
