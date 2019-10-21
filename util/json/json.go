//@author D-S
//Created on 2019/10/21 6:46 下午
package json

import (
	"encoding/json"
	"log"
)

func logJSON(dto interface{})  {
	s,_ := json.Marshal(dto)
	log.Println(string(s))
}
