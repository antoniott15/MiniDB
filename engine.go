package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
	"strings"
)

type Engine struct {
	TablesTree []*Table
	TotalDir int
	Dirs []string
}

type Table struct {
	StructureKeys      []int
	Name       string
	Headers []string
	StructTree *BPlussTree
}

type Structure struct {
	Key      int
	Headers []string
	Attribs map[string]interface{}
}

func NewEngine() (*Engine, error) {
	dirs := WalkDir(WorkingDir)

	var count int
	var tables []*Table
	for i, element := range dirs {
		if i != 0 {
			table, err := createTableByName(element)
			if err!= nil {
				return nil ,err
			}
			tables = append(tables, table)
		}
		count++
	}

	return &Engine{
		TablesTree: tables,
		TotalDir: count,
		Dirs: dirs,
	}, nil

}


func createTableByName(fileName string) (*Table, error) {
	tree := NewPlusTree()

	rawFile, err := os.Open(fileName)
	log.Info("Open file ", fileName)
	if err != nil {
		return nil, err
	}
	defer rawFile.Close()
	scanner := bufio.NewScanner(rawFile)
	var headers []string
	var keys []int
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			headers = getPositions(scanner.Text())
		} else {
			newStr := getValueByHeaders(headers, scanner.Text())
			str, err := json.Marshal(newStr)
			if err != nil {
				return nil, err
			}
			keys = append(keys, newStr.Key)
			if err := tree.Insert(newStr.Key, str); err != nil {
				return nil, err
			}
		}
	}
	return &Table{
		StructureKeys:      keys,
		Name:       fileName,
		Headers: headers,
		StructTree: tree,
	}, nil
}



func getValueByHeaders(headers []string, line string) *Structure {
	attributes, pk := getPositionsAndPK(line)

	attribs := make(map[string]interface{})
	for i,elements := range attributes {
		attribs[headers[i]] = elements
	}

	key, err := strconv.Atoi(fmt.Sprint(attribs["KEY"]))
	if err != nil {
		key = pk
	}

	return &Structure{
		Key:      key,
		Headers: headers,
		Attribs: attribs,
	}
}


func (e *Engine) createNewTable(name string, header []string) ([]*Table,error) {
	name = "./Tables/" + strings.ToLower(name)
	for _,elements := range e.Dirs {
		if strings.Contains(elements, "./") {
			elements = elements
		}else{
			elements = "./" + elements
		}
		if elements == name {
			return nil, tableNameFound
		}
	}
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()


	tree := NewPlusTree()

	head := headers(header)

	if err := Save(name,head); err != nil {
		return nil,err
	}

	table := &Table{
		StructureKeys:       nil,
		Name:       name,
		Headers: header,
		StructTree: tree,
	}

	e.TablesTree = append(e.TablesTree, table)
	e.Dirs = append(e.Dirs, name)
	e.TotalDir += 1
	log.Info("Table created correct ", name)
	return e.TablesTree, nil
}


func (e *Engine) insertIntoTable(elem *Structure, tableName string)  error {
	if !strings.Contains(tableName, "Tables") {
		tableName = "./Tables/" + tableName
	}

	for _,element := range e.TablesTree {
		if  strings.Contains(element.Name, "./") {
			if element.Name == tableName {
				element.StructureKeys = append(element.StructureKeys, elem.Key)
				element.Headers = elem.Headers
				c,err := json.Marshal(elem)
				if err != nil {
					return err
				}

				if err := element.StructTree.Insert(elem.Key,c); err != nil{
					return err
				}

				str := toFormat(*elem)
				if err := Save(tableName,str); err != nil{
					return err
				}
				return nil
			}
		}else{
			if "./" + element.Name == tableName {
				c,err := json.Marshal(elem)
				if err != nil {
					return err
				}
				element.StructureKeys = append(element.StructureKeys, elem.Key)
				element.Headers = elem.Headers
				if err := element.StructTree.Insert(elem.Key,c); err != nil{
					return err
				}
				str := toFormat(*elem)
				if err := Save(tableName,str); err != nil{
					return err
				}
				return nil

			}
		}
	}
	return nil
}


func Save(name string, elem interface{}) error {
	if !strings.Contains(name, "Tables") {
		name = "./Tables/" + name
	}
	fmt.Println(name)
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return  err
	}
	defer file.Close()

	dataWrite := bufio.NewWriter(file)
	_, err = io.WriteString(dataWrite, fmt.Sprintf("%s\n",elem))
	if err != nil {
		return err
	}
	dataWrite.Flush()
	return  nil
}

func (e *Engine) getTableByName(name string) *Table {
	for _,elements := range e.TablesTree {
		if strings.Contains(elements.Name, name) {
			return elements
		}
	}
	return nil
}