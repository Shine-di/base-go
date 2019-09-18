package cmd

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var env string

func init() {
	RootCmd.AddCommand(start)
	start.Flags().StringVarP(&env, "env", "e", "", "要启动的环境 dev/uat/hk/hkpub")
}

var start = &cobra.Command{
	Use:   "start",
	Short: "s",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start cmd,args为", args)
		if len(env) == 0 {
			cmd.Help()
			fmt.Println("需指定启动环境 dev/uat/hk/hkpub")
			return
		}

		StartWithEnv(env)

		//连接测试
		//redis.InitRedis()
		//mongo.InitMongo()
		////注册中心
		//eureka.InitEureka()
		//api.Run()
		return
	},
}

func StartWithEnv(env string) {

	configFile, err := ioutil.ReadFile("conf/" + env + ".yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}
	//yaml.Unmarshal(configFile, &conf.Yaml)
	//d, _ := yaml.Marshal(&conf.Yaml)
	//fmt.Println("配置文件 :", string(d))
}
