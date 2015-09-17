package main

import  (
         "testing"
		// "fmt"
		)
		
type Cir_Test struct {  

   radius float64
   perimeter float64
} 		

type Rect_Test struct {  

   height,width float64
   perimeter float64
} 		

var Circle_Cases = []Cir_Test {  // Test cases for Circle {Radius, Perimeter} 

  {3.4,21.362830044410593}, 
  {4.9,30.787608005179976}, 
  {4.0,25.132741228718345}, 
  {1.9,11.938052083641214}, 
  {5.4,33.929200658769766},
} 

var Rect_Cases = []Rect_Test { 
   {2,3,10}, 
   {7,7,28},
   {6,4,20},
} 
		
func TestGetPerimeter(t *testing.T) {
  
//  Circle_Variable := Circle{radius: 3.4} 
//  Rect_Variable := Rectangle{height : 3,width: 5}

  for _,check := range Circle_Cases { 

      if check.perimeter !=  getPerimeter(&Circle{radius:check.radius}) { 
		     
		  //  fmt.Println(getPerimeter(&Circle(radius:check.radius)  ))
		//	fmt.Println(check.perimeter)
		
			 t.Errorf("Test Fail")
       }
	 }
	
	for _,params := range Rect_Cases { 

      if params.perimeter !=  getPerimeter(&Rectangle{height:params.height,width:params.width}) { 
		    
			 t.Errorf("Test Fail")
       }
	 	
  }
} 


 
 


   