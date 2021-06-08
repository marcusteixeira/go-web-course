package main 

import "fmt"

func main(){
	ok, err := say("Hello :)")
	if err != nil {
		panic((err.Error))
	}
	switch ok {
	case true:
		fmt.Println("OK")
	default:
		fmt.Println("Deu merda")		
	}
}

func say(what string) (bool,error){
	if what == ""{
		return false,fmt.Errorf("Empty String")
	}
	fmt.Println(what)
	return true, nil 
}