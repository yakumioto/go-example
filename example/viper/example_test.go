package viper

import (
	"fmt"
	"os"
	"strings"

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

// ExampleAutomaticEnv
// 设置 Automatic env 使用 Get 相关的方法会依次查询, env 的优先级高. 所以会覆盖参数.
// 查询环境变量时需要使用 '_', 如果想使用 '.' 需要使用 SetEnvKeyReplacer 进行替换.
func ExampleAutomaticEnv() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./etc/viper")
	viper.ConfigFileUsed()
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	_ = os.Setenv("HELLO", "emmm")
	fmt.Println(viper.GetString("hello"))
	fmt.Println(viper.GetString("o1.o2.hello"))
	_ = os.Setenv("O1_O2_HELLO", "emmm")
	fmt.Println(viper.GetString("o1.o2.hello"))
	fmt.Println(viper.GetString("o1_o2_hello"))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fmt.Println(viper.GetString("o1.o2.hello"))
	// output:
	// emmm
	// world
	// world
	// emmm
	// emmm
}
