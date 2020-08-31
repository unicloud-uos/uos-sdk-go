package main

import (
	"fmt"
	"os"

)



func HandleError(err error) {
	fmt.Println("panic err:", err)
	//err = DeleteTestBucketAndObject()
	//if err != nil {
	//	fmt.Println("DeleteTestBucketAndObject err:", err)
	//}
	os.Exit(-1)
}
