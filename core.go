package main

func main() {
	api, err := newDBAPI(PREFIX, PORT)
	if err != nil {
		panic(err)
	}
	api.Launch()

	// TEST
	// os.Remove("./Tables/test")
	// f, err := os.OpenFile("./Tables/test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// f.Write([]byte("KEY|APELLIDOS|CARRERA|MENSUALIDAD"))

	// engine, err := NewEngine()
	// if err != nil {
	// 	panic(err)
	// }

	// const timesTo = 1000000
	// for i := 0; i < timesTo-1; i++ {
	// 	engine.insertIntoTable(&Structure{
	// 		Key:     i,
	// 		Headers: []string{"KEY", "APELLIDOS", "CARRERA", "MENSUALIDAD"},
	// 		Attribs: map[string]interface{}{
	// 			"KEY":         i,
	// 			"APELLIDOS":   "asdas",
	// 			"CARRERA":     "ASDASD",
	// 			"MENSUALIDAD": "ASDADSADS",
	// 		},
	// 	}, "test")
	// }

	// start := time.Now()
	// engine.insertIntoTable(&Structure{
	// 	Key:     timesTo,
	// 	Headers: []string{"KEY", "APELLIDOS", "CARRERA", "MENSUALIDAD"},
	// 	Attribs: map[string]interface{}{
	// 		"KEY":         timesTo,
	// 		"APELLIDOS":   "asdas",
	// 		"CARRERA":     "ASDASD",
	// 		"MENSUALIDAD": "ASDADSADS",
	// 	},
	// }, "test")

	// elapsedInsert := time.Since(start)

	// log.Printf("Insert took %s", elapsedInsert)

	// t := engine.getTableByName("test")

	// startSearch := time.Now()
	// _, err = t.StructTree.Search(timesTo - 1000)
	// elapsedSearch := time.Since(startSearch)

	// log.Printf("Search took %s", elapsedSearch)

}
