package main

import "fmt"

func main() {
     const n int = 25;
     
     var first, second, next int = 0, 1, 0;

     fmt.Print("fibonacci: ");

     for i := 0; i < n; i++ {
     	 if i <= 1 {
	    next = i;
	 } else {
            next = first + second;
            first = second;
            second = next;
         }
         fmt.Print(next, " ");
     }
     fmt.Println();
}