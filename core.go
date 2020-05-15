package main

func main() {
	api, err := newDBAPI(PREFIX, PORT)
	if err != nil {
		panic(err)
	}
	api.Launch()

	// engine, err := NewEngine()
	// if err != nil {
	// 	panic(err)
	// }

	// var insertTime time.Duration
	// const timesTo = 10
	// for i := 0; i < timesTo; i++ {
	// 	start := time.Now()
	// 	engine.insertIntoTable(&Structure{
	// 		Key:     i,
	// 		Headers: []string{"KEY", "APELLIDOS", "CARRERA", "MENSUALIDAD"},
	// 		Attribs: map[string]interface{}{
	// 			"KEY":         i,
	// 			"APELLIDOS":   "asdas",
	// 			"CARRER":      "ASDASD",
	// 			"MENSUALIDAD": "ASDADSADS",
	// 		},
	// 	}, "test")

	// 	elapsed := time.Since(start)
	// 	insertTime += elapsed
	// }

	// log.Printf("Search took %s",insertTime / timesTo)

	// var searchTime time.Duration
	// t := engine.getTableByName("test")
	// for i := 0; i < 10; i++ {
	// 	startSearch := time.Now()
	// 	t.StructTree.Search(i)

	// 	elapsedSearch := time.Since(startSearch)
	// 	searchTime += elapsedSearch
	// }
	// log.Printf("Search took %s", searchTime/timesTo)

}
