package viper

import (
	"fmt"

	"github.com/spf13/viper"
)

func ExampleViperSetDefault() {
	viper.SetDefault("hello", "world")

	fmt.Println(viper.GetString("hello"))
	// output:
	// world
}

// ExampleAddConfigPath
// 无覆盖效果, 根据 config path 的顺序依次查找, 找到后直接返回
// 例如如果 /etc/viper, 中有 config.yaml, 就不会去找 ./.viper
func ExampleAddConfigPath() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./etc/viper")
	viper.AddConfigPath("./.viper")
	viper.AddConfigPath(".")
	viper.ConfigFileUsed()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("hello"))
	// output:
	// world
}
