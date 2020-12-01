package main

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	
	bz "Week02/biz"
	"os"
)

func main() {
	mode  := ""
	arg_num := len(os.Args)
	if arg_num >1{
		mode = os.Args[1]
		log.Println("input mode is: " + mode)
	}
	
	_, err := bz.Biz(mode)
	if err!= nil {
		log.Printf("Main: %+v\n", err)
	}

}



