package main

import "fmt"

func Fibonacci(x float64) float64 { 
   
   if x==0 { 
   return 0
   } 
  
  if x==1 || x ==2 {
      return 1
   } 
   
   return Fibonacci(x-1) + Fibonacci(x-2)
} 
func main() {
  
var i float64
var input float64

fmt.Print("Enter the no for which the Fibonacci Series should be generated :");
fmt.Scanf("%f",&input)

  var fin float64 

for i=0; i<=input; i++  {
	
   fin  =Fibonacci(i)  	
   //   fmt.Println(Fibonacci(i))
  }
fmt.Println(fin)
}