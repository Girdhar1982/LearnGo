package main //first line must be package name and file name must end with  .go

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func main() {
	//excerciseVariableAndPointers()
//	variablelocal()
//usingStruct()
//usingMap()
//usingSlices()
//usingDecision()
usingforLoops()
}


func usingforLoops(){
for i:=0; i<=100; i++{
i=i+10;
fmt.Println(i)
}

animals:=[]string{"dog","cat","horse",
//"cow","ox","fish",
"lion"}

for i,_:=range animals{
	log.Println(i,animals[i])
}
for _,animal:=range animals{
	log.Println(animal)
}

tags := make(map[string]string)
tags["name"]="Test"
tags["enviornment"]="nonprod"
for key,value:=range tags{
	log.Println(key,value)
}
}

func usingDecision(){
isTrue:=true
if isTrue {log.Println("Hello ",isTrue)}else{
	log.Println("Bye ",isTrue)
}

if isTrue != true {log.Println("Hello ",isTrue)}else{
	log.Println("Bye ",isTrue)
}
num:=5
name:="manish"
if isTrue && name == "manish"{log.Println("Hello - ", name)}
if isTrue && name != "manish"{log.Println("Hello! ", name)}
if num > 2 && isTrue {log.Println("Hello! -", name)}
usingSwitch("manish")
usingSwitch("nisha")
usingSwitch("manisha")
}
func usingSwitch(name string){
switch(name){
case "manish":
	log.Println("Hello ", name)
case "nisha":
	log.Println("Hello new user ", name)
default:
	log.Println("Hello unknown user ", name)
}
}


func usingMap(){
tag := make(map[string]string)
tag["name"]="manish"
tag["environment"]="prod"
log.Println(tag["name"])
usersMap := make(map[string]User)
usersMap["user1"]=User{FirstName: "Manish",LastName: "Girdhar"}
log.Println(usersMap["user1"])
}

var s = "seven" //go is smart enough to know it is string type variable; this variable is outside all the functions so value is available everywhere (global)
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
} //struct type is used to store data about an object
func (u *User) printFirstName() string{
	return u.FirstName;
	}
func variablelocal(){
	var s2 = "six" //local variable to function
	log.Println("s is",s); //display global variable
	log.Println("s2 is",s2); // display local variable
}

func usingStruct(){ //if function name is starting with lowercase letter function is private to the package
	user := User{	FirstName: "Manish",LastName: "Girdhar"}
	user.PhoneNumber="0435899999"
	log.Println(user.FirstName,user.LastName)
	user2:=User{FirstName: "Nisha",LastName: "Girdhar",PhoneNumber: "0406199999"}
	log.Println(user2.printFirstName(),user2.LastName)
}
func usingSlices(){
var myArrayorSlice []string
myArrayorSlice=append(myArrayorSlice, "Javin")
myArrayorSlice=append(myArrayorSlice, "Manish")
myArrayorSlice=append(myArrayorSlice, "Nisha")
sorted:=sort.StringSlice(myArrayorSlice)
log.Println(myArrayorSlice,sorted)
numbers:=[]int{1,7,2,8,5}
log.Println(numbers[0:4])
}

func excerciseVariableAndPointers(){
	value1,value2:=returntwostrings(); //assigning values from function retruning two values
	fmt.Println(value2,value1);
	pointers();
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

func pointers(){
mystring:="Green"
log.Println("before: mystring is set to ", mystring)
changeUsingPointer(&mystring) //& is added in front to take the memory pointer
log.Println("after: mystring is set to ", mystring)

}

func changeUsingPointer(s *string){ //this function takes memory reference
log.Println("memory address of RAM is shown here - ",s);
newValue:="Red"
*s=newValue;
}