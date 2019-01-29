package main

import (
	"bufio"
	"fmt"
	"github.com/pathao-eng/empleo/core"
	"github.com/spf13/viper"
	"html/template"
	"log"
	"os"
	"plugin"
	"strings"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		viper.AddConfigPath(os.Args[1])
	}
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var sources []core.EmpleoSource

	srcCfg := viper.GetStringMapString("sources")
	for k, v := range srcCfg {
		plug, err := plugin.Open(v)
		if err != nil {
			log.Println(err)
			continue
		}

		sym, err := plug.Lookup(strings.ToUpper(k))
		if err != nil {
			log.Println(err)
			continue
		}

		a, ok := sym.(core.EmpleoSource)
		if !ok {
			log.Println("Symbol is not type of EmpleoSource")
			continue
		}

		sources = append(sources, a)
	}

	var empleos []core.Empleo

	for _, s := range sources {
		if err := s.Init(); err != nil {
			log.Println(err)
			continue
		}
		data := fetchAll(s, []core.Empleo{})
		empleos = append(empleos, data...)
	}

	timeNow := time.Now().Format("2006-01-02")

	f, err := os.Create(fmt.Sprintf("./results/%s.md", timeNow))
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)

	t := template.New("EmpleoTemplate")
	t, err = t.ParseGlob("templates/*.tmpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, map[string]interface{}{
		"parsed_date": timeNow,
		"empleos":     empleos,
	})
	if err != nil {
		panic(err)
	}
}

func fetchAll(s core.EmpleoSource, init []core.Empleo) []core.Empleo {
	data, ok := s.Fetch()
	init = append(init, data...)
	if ok {
		return fetchAll(s, init)
	}
	return init
}
