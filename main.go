package main

import (

"fmt"
"log"
"github.com/joho/godotenv"

)

func main(){
	fmt.Println("Hello World");

	godotenv.Load(".env")
	if(portString == ""){
		log.Fatal("PORT is not set in enviornemnt")
	}
	fmt.Println("Port:%v", godotenv)
}