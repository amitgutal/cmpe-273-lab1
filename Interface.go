package main

import ("math";"fmt")

type Shape interface { 

 Perimeter() float64
} 

type Rectangle struct {
 
 x1,y1,x2,y2 float64
 height float64
 width float64
} 

type Circle struct { 

 x , y , radius float64
} 

func (r *Rectangle) Perimeter() float64 { 

 fmt.Print("Perimeter (Rectangle) :")
 return  2 * (r.height + r.width)
} 

func (c *Circle) Perimeter() float64  {

 fmt.Print("Perimeter (Circle) :")
 return 2 * math.Pi * c.radius
} 

func getPerimeter(s Shape) float64  {  
 
 return s.Perimeter()
} 

func main() { 

  Circle_Variable := Circle{radius: 5.4} 
  Rect_Variable := Rectangle{height : 3,width: 5} 
 
  fmt.Println(getPerimeter(&Circle_Variable))
  fmt.Println(getPerimeter(&Rect_Variable))
} 



