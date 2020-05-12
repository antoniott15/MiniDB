package main

import "errors"

var keyExist = errors.New("key already exists")
var keyNotFound = errors.New("key not found")
var treeIsEmpty = errors.New("tree is empty")
var tableNameFound = errors.New("table already exist")