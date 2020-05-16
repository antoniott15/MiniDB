package main

import "strings"

func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func getPositions(line string) []string {
	return SplitAny(line, "|")
}

func (head *Header) toStudentFormat(name, lastName, career, monthlyPayment string, pos []string) string {
	student := make([]string, len(pos))
	for index, elements := range pos {
		if elements == "KEY" {
			student[index] = name
		} else if elements == "Apellidos" {
			student[index] = lastName
		} else if elements == "Carrera" {
			student[index] = career
		} else if elements == "Mensualidad" {
			student[index] = monthlyPayment
		}
	}
	var finalStudent string
	for i := 0; i < len(student); i++ {
		finalStudent += student[i] + "|"
	}
	return finalStudent[:len(finalStudent)-1]
}
