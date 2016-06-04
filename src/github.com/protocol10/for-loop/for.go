package main

import "fmt"

func main(){

    i:= 1

    for i<3 {
      fmt.Println("Values of i = ",i)
      i += 1 // increment the counter value else it will be  infinite loop
    }

    for j:= 3; j<7; j++ {
      fmt.Println("Value for j", j)
      i:= 0
      for i < j{
          fmt.Println("Value of i and j ", i, j)
          i += 1
      }
    }

    i= 0
    for{
      fmt.Println("loop")

      if(i > 4){
        fmt.Println("Termination of for-loop")
        break;
      }
      i += 1;
    }
}
