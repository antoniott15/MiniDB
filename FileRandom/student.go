package main

import (
	"bufio"
	"fmt"
	"time"

	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Student struct {
	KEY            string
	LastName       string
	Career         string
	MonthlyPayment string
}

type Header struct {
	totalCount int
	students   []*Student
	NameFile   string
	Position   []string
}

func newFile(file string) (*Header, error) {
	rawFile, err := os.Open(file)
	log.Info("Open file ", file)
	if err != nil {
		return nil, err
	}
	defer rawFile.Close()

	scanner := bufio.NewScanner(rawFile)
	var students []*Student

	var pos []string
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			pos = getPositions(scanner.Text())
		} else {
			students = append(students, getStudentByPos(scanner.Text(), pos))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &Header{
		totalCount: len(students),
		students:   students,
		NameFile:   file,
		Position:   pos,
	}, nil
}

func (head *Header) GetAllStudents() []*Student {
	return head.students
}
func (head *Header) SearchStundet(key string) *Student {
	for _, elements := range head.students {
		if elements.KEY == key {
			return elements
		}
	}
	return nil
}

func getStudentByPos(line string, pos []string) *Student {
	studentSplited := SplitAny(line, "|")
	var student Student
	for index, elements := range studentSplited {
		if pos[index] == "KEY" {
			student.KEY = elements
		} else if pos[index] == "Apellidos" {
			student.LastName = elements
		} else if pos[index] == "Carrera" {
			student.Career = elements
		} else if pos[index] == "Mensualidad" {
			student.MonthlyPayment = elements
		}
	}

	return &student
}

func (head *Header) CreateNewStudent(name, lastName, career, monthlyPayment string) error {
	student := &Student{
		KEY:            name,
		LastName:       lastName,
		Career:         career,
		MonthlyPayment: monthlyPayment,
	}
	head.students = append(head.students, student)
	head.totalCount += 1

	file, err := os.OpenFile(head.NameFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	stu := head.toStudentFormat(name, lastName, career, monthlyPayment, head.Position)

	dataWrite := bufio.NewWriter(file)

	_, err = io.WriteString(dataWrite, "\n"+stu)
	if err != nil {
		return err
	}
	dataWrite.Flush()

	return nil
}

func main() {
	os.Remove("./datos.txt")
	f, err := os.OpenFile("datos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.Write([]byte("KEY|Apellidos|Carrera|Mensualidad"))

	head, err := newFile("./datos.txt")
	if err != nil {
		log.Error(err)
	}

	const timesTo = 1000000

	for i := 0; i < timesTo-1; i++ {
		if err := head.CreateNewStudent(fmt.Sprint(i), "toche", "cs", "1600"); err != nil {
			log.Error(err)
		}
	}
	start := time.Now()
	if err := head.CreateNewStudent(fmt.Sprint(timesTo), "toche", "cs", "1600"); err != nil {
		log.Error(err)
	}
	elapsedInsert := time.Since(start)

	log.Printf("Insert took %s", elapsedInsert)

	searchT := time.Now()
	_ = head.SearchStundet(fmt.Sprint(timesTo - 10000))
	elapsedSearch := time.Since(searchT)

	log.Printf("Search took %s", elapsedSearch)

}
