package models

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestAddUserRequest_String(t *testing.T) {
	req := &AddUserRequest{}
	req.Name = "admin"
	req.Password = "123456"
	fmt.Println(req.GetName())
	fmt.Println(req.GetPassword())
	fmt.Println(req.String())

	// marshal
	body, err := proto.Marshal(req)
	if err!=nil{
		t.Error(err)
		return
	}
	fmt.Println(body)

	// unmarshal
	req_unmashal := &AddUserRequest{}
	err = proto.Unmarshal(body, req_unmashal)
	if err!=nil{
		t.Error(err)
		return
	}
	fmt.Println(req_unmashal)
}

