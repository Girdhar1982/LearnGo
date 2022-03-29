package main   //first line must be package name and file name must end with  .go

import "fmt"

func main() {
value1,value2:=returntwostrings(); //assigning values from function retruning two values
fmt.Println(value2,value1);
}

func first() string{
	//declared variable must be used otherwise go will not compile
	fmt.Println("Hello! ");
	var  myname string;       //string variable
	myname="Manish";          //assigning value to variable
	var age int;             //int variable defaults to int64 if machine is  64 bit
	age=42;                  //assign int value
	lastname := "Girdhar"   //string without variable declaration
	fmt.Println(myname,lastname,age);
	return "done"
}

func returntwostrings() (string,string){ //in go function can retrun two values
valuefromfirst:=first() ;
return valuefromfirst, "secondstring"
}