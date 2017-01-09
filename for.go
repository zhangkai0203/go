package main

import "fmt"

var a int = 10

func main(){

        numbers := [...]int{1,2,3,4,5,6,7,8,9}

        /*for i := 0;i < 10;i++{
		    fmt.Println(i)
		}*/
		
		/*for a <= 10 {
		    a++;
			fmt.Println(a)
		}*/
		
		/*for i := 0;i < len(numbers); i++{
		    fmt.Println(numbers[i])
		}*/
		
		for i := range numbers {
		    fmt.Println(numbers[i])
		}
}