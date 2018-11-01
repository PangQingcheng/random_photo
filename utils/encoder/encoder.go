package encoder

import (
	uuid "github.com/satori/go.uuid"
)

func GetUniqueString(name string)(string,error){
	uuid1,err := uuid.NewV1()
	if err!=nil {
		return "",err
	}
	uid := uuid.NewV5(uuid1,name)
	return uid.String(),nil
}
