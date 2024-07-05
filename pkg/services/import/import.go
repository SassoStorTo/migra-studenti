package impo_service

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/SassoStorTo/migra-studenti/pkg/models"
	"github.com/SassoStorTo/migra-studenti/pkg/services/classes"
	"github.com/SassoStorTo/migra-studenti/pkg/services/majors"
	"github.com/SassoStorTo/migra-studenti/pkg/services/students"
)

func ParseFile(path string, startYear int) error {
	log.Println("Parsing file")
	fd, error := os.Open(path)
	if error != nil {
		fmt.Println(error)
	}
	log.Println("Successfully opened the CSV file")
	defer fd.Close()

	reader := csv.NewReader(fd)
	reader.Comma = ','
	clss := classes.GetAll()
	mjrs := majors.GetAll()

	currIdStudent := students.GetLastId() + 1
	currIdClass := classes.GetLastId() + 1
	currIdMajor := majors.GetLastId() + 1

	for {
		record, err := reader.Read()
		if err != nil {
			if err == csv.ErrFieldCount {
				fmt.Println("Il formato del file non e' corretto!")
				continue
			}
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading record:", err)
			return fmt.Errorf("il formato del file non e' corretto")
		}

		if len(record) != 3 {
			return fmt.Errorf("il formato del file non e' corretto, numero di campi errato")
		}

		dateOfBirth, err := time.Parse("02/01/2006", record[3])
		if err != nil {
			return fmt.Errorf("il formato del file non e' corretto, data di nascita non valida")
		}
		student := models.Student{Name: record[0], LastName: record[1], Id: currIdStudent, DateOfBirth: dateOfBirth}
		student.Save()
		currIdStudent++

		classInfo := strings.Split(record[2], " ")
		year := int(classInfo[0][0]) - 48
		if year < 0 || year > 5 {
			return fmt.Errorf("il formato del file non e' corretto, Anno errato")
		}

		section := classInfo[0][1:]
		major := classInfo[1]

		idxMajor := searchMajor(major, mjrs)
		if idxMajor == -1 {
			mjr := models.Majors{Name: major, Id: currIdMajor}
			mjr.Save()
			*mjrs = append(*mjrs, mjr)
			currIdMajor++
			idxMajor = len(*mjrs) - 1
		}

		idxClass, err := searchClassFromStirng(year, section, idxMajor, clss, mjrs)
		if err != nil {
			return err
		}
		if idxClass == -1 {
			cls := models.Class{Year: year, Section: section, IdMajor: (*mjrs)[idxMajor].Id,
				ScholarYearStart: time.Now().Year(), Id: currIdClass}
			cls.Save()
			*clss = append(*clss, cls)
			currIdClass++
			idxClass = len(*clss) - 1
		}

		studCls := *models.NewStudentClass(currIdStudent-1, (*clss)[idxClass].Id, time.Now())
		studCls.Save()
	}

	return nil
}

func searchClassFromStirng(year int, section string, idxMajor int, classes *[]models.Class, majors *[]models.Majors) (int, error) {
	for idx, c := range *classes {
		if c.Year == year && c.Section == section && (*majors)[idxMajor].Name == getMajorFormIdMajor(c.IdMajor, majors) {
			return idx, nil
		}
	}
	return -1, nil
}

func getMajorFormIdMajor(id int, majors *[]models.Majors) string {
	for _, m := range *majors {
		if m.Id == id {
			return m.Name
		}
	}
	return ""
}

func searchMajor(major string, majors *[]models.Majors) int {
	for index, m := range *majors {
		if m.Name == major {
			return index
		}
	}
	return -1
}
