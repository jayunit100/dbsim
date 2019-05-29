package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
	"github.com/gorilla/mux"
	"github.com/jayunit100/kafka-sim/pkg/components"
)


type Configuration struct {
	Port int
}

func newConfig() Configuration {
	_ = components.DefaultMQ{}

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration Configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return configuration
}

func main() {
	c := newConfig()
	r := newRouter()
	r.HandleFunc("/", r.handler)
	log.Info("starting with config %v",c.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",c.Port), r))
}

type Router struct {
	*mux.Router
}

func newRouter() *Router {
	r := mux.NewRouter()
	return &Router{r}
}

func (r *Router) handler(w http.ResponseWriter, req *http.Request) {
	log.Info("Starting app")
	fmt.Fprintf(w, "hello")
}
