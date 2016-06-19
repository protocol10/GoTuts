package main

import "fmt"

func main(){

  var a[5] int; // Need to explicity define data type of array, else it gives warning

  fmt.Println("Array is ", a);

  a[4] = 200;

  fmt.Println("Array is ",a);
  fmt.Println("Value at index 4 is ", a[4]);

  fmt.Println("Length of array is ", len(a));

  b:=[5] int {1, 2, 3, 4, 5};

  fmt.Println("Array B is ",b);

  var two_D[2][3] int;

  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      two_D[i][j] = i + j;
    }
  }

  fmt.Println("Two Dimensional Array is ",two_D);
}
