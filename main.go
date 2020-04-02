// Programmer name: Vincent G.
// Date completed:  4/1/2020
// Description: Sect. 1 You Code it

package main

import "fmt"

//create a function that accepts a name and prints a greeting message using the name
func Greeting(){
  var Name string
  fmt.Println()
  fmt.Println("Please Enter Your Name: ")
  fmt.Scanln(&Name)
  fmt.Println()
  fmt.Println("Greetings",Name,"!")
}

func main() {
    //call your function
    Greeting()
}