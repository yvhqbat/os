package models

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestAddUserRequest_String(t *testing.T) {
	u := &UserInfo{}
	u.AccessKey = "ak_123456"
	u.SecretKey = "sk_123456"
	fmt.Println(u.GetAccessKey())
	fmt.Println(u.GetSecretKey())
	fmt.Println(u.String())

	// marshal
	body, err := proto.Marshal(u)
	if err!=nil{
		t.Error(err)
		return
	}
	fmt.Println(body)

	// unmarshal
	req_unmashal := &UserInfo{}
	err = proto.Unmarshal(body, req_unmashal)
	if err!=nil{
		t.Error(err)
		return
	}
	fmt.Println(req_unmashal)
}

