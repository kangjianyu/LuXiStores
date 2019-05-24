package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

type Postgres struct {
	Mysql    string
	Password string
}
type Config struct {
	Postgres Postgres
}

func main() {
	doc, err := ioutil.ReadFile("./user/config/config.toml")
	fmt.Println(err)
	config := Config{}
	_ = toml.Unmarshal(doc, &config)
	fmt.Println("user=", config.Postgres.Mysql)
	fmt.Println(config.Postgres)
}
