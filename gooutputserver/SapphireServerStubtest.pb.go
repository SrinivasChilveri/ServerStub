package test

import "github.com/golang/protobuf/proto"

/**
 *  This is a test comment on testobj 
 *  
 **/
type Testobj struct {
  Oid uint64
}

/**
 *  This is a test comment on func1  
 **/
func (Obj Testobj) Func1_Wrap(inParams  [] byte) (outParams  [] byte) {
  InPutStruct := &Msg1{} ; 
  _ = proto.Unmarshal(inParams, InPutStruct)
  OutPutStruct := &Msg2{} ; 
  OutPutStruct.Liststring, OutPutStruct.Idtostringmap =  Obj.Func1(InPutStruct.Name , InPutStruct.Id , InPutStruct.Floatingvalue )
  outParams, _ = proto.Marshal(OutPutStruct)

  return 
}

/**
 *  This is a multiline test 
 *  Comment on func2 
 *  
 **/
func (Obj Testobj) Func2_Wrap(inParams  [] byte) (outParams  [] byte) {
  InPutStruct := &Msg2{} ; 
  _ = proto.Unmarshal(inParams, InPutStruct)
  OutPutStruct := &Msg3{} ; 
  OutPutStruct.Nested =  Obj.Func2(InPutStruct.Liststring , InPutStruct.Idtostringmap )
  outParams, _ = proto.Marshal(OutPutStruct)

  return 
}

/**
 *  This is a different test 
 * comment on func3  
 **/
func (Obj Testobj) Func3_Wrap(inParams  [] byte) (outParams  [] byte) {
  InPutStruct := &Msg3{} ; 
  _ = proto.Unmarshal(inParams, InPutStruct)
  OutPutStruct := &Msg1{} ; 
  OutPutStruct.Name, OutPutStruct.Id, OutPutStruct.Floatingvalue =  Obj.Func3(InPutStruct.Nested )
  outParams, _ = proto.Marshal(OutPutStruct)

  return 
}

/**
 *  This is a different test 
 * comment on func4  
 **/
func (Obj Testobj) Func4_Wrap(inParams  [] byte) (outParams  [] byte) {
  InPutStruct := &Msg4{} ; 
  _ = proto.Unmarshal(inParams, InPutStruct)
  OutPutStruct := &Msg4{} ; 
  OutPutStruct.Params =  Obj.Func4(InPutStruct.Params )
  outParams, _ = proto.Marshal(OutPutStruct)

  return 
}

