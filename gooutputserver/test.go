

package main

import "CodeGenTest/test"
import "github.com/golang/protobuf/proto"
import (
	//"github.com/golang/protobuf/proto"
	"fmt"
	//"gostubs/gooutput"
	//"gostubs/gooutput"
	//"gostubs/gooutput"
	//"gostubs/gooutput"
)

//import "CodeGenTest/test"

func main() {
	fmt.Printf("Hello\n")
	t := test.Testobj {Oid: 4}

	msg := &test.Msg1{
		Name: "abccc",
		Id: 3,
		Floatingvalue: 5.0,
	}
	msgbuffer, _ := proto.Marshal(msg)

	fmt.Println("proto.Marshal",msgbuffer)

 a := t.Func1_Wrap(msgbuffer)
  fmt.Println("output",a)
}
