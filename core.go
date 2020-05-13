package main

func main() {
	api, err := newDBAPI(PREFIX, PORT)
	if err != nil {
		panic(err)
	}
	api.Launch()
}

