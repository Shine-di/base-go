//@author D-S
//Created on 2019/9/18 3:41 下午
package main

import (
	eureka_client "base-go/eureka"
	"encoding/json"
	"fmt"
	"net/http"
)
func main() {


	// create eureka client
	client :=  eureka_client.NewClient(&eureka_client.Config{
		DefaultZone:           "http://127.0.0.1:8080/eureka/",
		App:                   "go-example",
		Port:                  10000,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
		Metadata: map[string]interface{}{
			"VERSION":              "0.1.0",
			"NODE_GROUP_ID":        0,
			"PRODUCT_CODE":         "DEFAULT",
			"PRODUCT_VERSION_CODE": "DEFAULT",
			"PRODUCT_ENV_CODE":     "DEFAULT",
			"SERVICE_VERSION_CODE": "DEFAULT",
		},
	})
	// start client, register、heartbeat、refresh
	client.Start()

	// http server
	http.HandleFunc("/v1/services", func(writer http.ResponseWriter, request *http.Request) {
		// full applications from eureka server
		apps := client.Applications

		b, _ := json.Marshal(apps)
		_, _ = writer.Write(b)
	})

	// start http server
	if err := http.ListenAndServe(":10000", nil); err != nil {
		fmt.Println(err)
	}
}


//删除  idea命令
// git rm --cached -r .idea
// git rm -r --cached .