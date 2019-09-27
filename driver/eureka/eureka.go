//@author D-S
//Created on 2019-07-10 13:06

package eureka

import (
	"cortex3/conf"
	"fmt"
	"github.com/Shine-di/base-go/eureka"
)

//"http://wind:fa87d67b-e53c-4139-ad0b-44c4943b5e7a@47.56.47.110:80/eureka/",

func InitEureka(){
	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           conf.Yaml.Conf.Eureka.DefaultZone,
		App:                   conf.Yaml.Conf.Eureka.App,
		Port:                  80,
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
	if err := client.Start();err != nil{
		panic("eureka 无法连接 " + conf.Yaml.Conf.Eureka.DefaultZone + err.Error())
	}

	fmt.Println("注册成功",conf.Yaml.Conf.Eureka.App)
}
