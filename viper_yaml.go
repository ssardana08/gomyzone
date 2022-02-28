package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("slice_config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	x := viper.GetStringSlice("name")
	fmt.Println(x)
}
