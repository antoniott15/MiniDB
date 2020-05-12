package main

import (
	"os"
	"os/signal"
)

func main() {
	if err := Run(); err != nil{
		panic(err)
	}
}


func Run() error {
	api, err := newDBAPI("/api", ":4200")
	if err != nil {
		panic(err)
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		api.done <- nil
	}()
	go func() {
		if api.engine != nil {
			api.done <- api.Launch()
		}
	}()

	return <-api.done
}