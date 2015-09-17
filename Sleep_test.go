package main

import "testing"
import "time"

type TimeOut struct {
	
        seconds1 int
        actual time.Duration
}

var TimeOutTests = []TimeOut {
	
        {3, 3*time.Second}, {4, 4*time.Second}, 
}

func TestFunctionSleep(t *testing.T) {

	for _, counter := range TimeOutTests {
                start := time.Now().UTC()
				FunctionSleep(counter.seconds1)
				stop := time.Now().UTC()
				
				var timeElapsed time.Duration = stop.Sub(start)
				
				if timeElapsed < counter.actual{
					
				    t.Errorf("Fail ")
                }
    
	}
}