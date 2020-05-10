package models

type UserInfo struct{
	AccessKey string `json:"access_key"`
	SecreteKey string `json:"secrete_key"`
	Name string `json:"name"`
}
