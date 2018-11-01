package main

import (
	"fmt"
	"random_photo/utils/encoder"
)

func main() {
	res,err := encoder.GetUniqueString("pangqingcheng");
	if err!= nil {
		panic(err)
	}
	fmt.Println(res)
}
