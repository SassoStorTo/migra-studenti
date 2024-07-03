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

func SwapStudentId(id int, student models.Student, local_studentclass *[]models.StudentClass) {
	for _, e := range *local_studentclass {
		if student.Id == e.IdS {
			e.IdS = id
		}
	}
	student.Id = id
}

func SwapClassId(id int, class models.Class, local_studentclass *[]models.StudentClass) {
	for _, e := range *local_studentclass {
		if class.Id == e.IdC {
			e.IdS = id
		}
	}
	class.Id = id
}

func SwapMajorId(id int, major models.Majors, local_classes *[]models.Class) {
	for _, c := range *local_classes {
		if major.Id == c.IdMajor {
			c.IdMajor = id
		}
	}
	major.Id = id
}

func AddDataToDb(local_students *[]models.Student, local_classes *[]models.Class, local_studentClasses *[]models.StudentClass, local_majors *[]models.Majors) error {
	// todo: implementare una sorta di lock (se ho voglia)

	for _, m := range *local_majors {
		err := m.Save()
		if err != nil {
			return err
		}
		SwapMajorId(majors.GetLastId(), m, local_classes)
	}

	for _, s := range *local_students {
		err := s.Save()
		if err != nil {
			return err
		}
		SwapStudentId(students.GetLastId(), s, local_studentClasses)
	}

	for _, c := range *local_classes {
		err := c.Save()
		if err != nil {
			return err
		}
		SwapClassId(classes.GetLastId(), c, local_studentClasses)
	}

	// finisci questo ????
	for _, sc := range *local_studentClasses {
		err := sc.Save()
		if err != nil {
			return err
		}

	}

	return nil
}

func ParseFile(path string, startYear int) (*[]models.Student, *[]models.Class, *[]models.StudentClass, *[]models.Majors, error) {
	log.Println("Parsing file")
	fd, error := os.Open(path)
	if error != nil {
		fmt.Println(error)
	}
	log.Println("Successfully opened the CSV file")
	defer fd.Close()

	reader := csv.NewReader(fd)
	reader.Comma = ','
	students := make([]models.Student, 1)
	classes := make([]models.Class, 1)
	studentClasses := make([]models.StudentClass, 1)
	majors := make([]models.Majors, 1)

	currIdStudent := 0
	currIdClass := 0
	currIdMajor := 0

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
			return nil, nil, nil, nil, fmt.Errorf("il formato del file non e' corretto")
		}

		if len(record) != 4 {
			return nil, nil, nil, nil, fmt.Errorf("il formato del file non e' corretto")
		}

		if record[2] != "Frequenta" {
			continue
		}

		students = append(students, models.Student{Name: record[0], LastName: record[1], Id: currIdStudent})
		currIdStudent++

		classInfo := strings.Split(record[3], " ")
		year := int(classInfo[0][0])
		if year < 49 || year > 53 {
			return nil, nil, nil, nil, fmt.Errorf("il formato del file non e' corretto, Anno errato")
		}

		section := classInfo[0][:1]
		major := classInfo[1]

		idxMajor := searchMajor(major, &majors)
		if idxMajor == -1 {
			majors = append(majors, models.Majors{Name: major, Id: currIdMajor})
			currIdMajor++
			idxMajor = len(majors) - 1
		}

		idxClass, err := searchClassFromStirng(year, section, idxMajor, &classes, &majors)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		if idxClass == -1 {
			classes = append(classes, models.Class{Year: year, Section: section, IdMajor: -1,
				ScholarYearStart: time.Now().Year(), Id: currIdClass})
			currIdClass++
			idxClass = len(classes) - 1
		}

		//Todo: qua non setto l'id
		studentClasses = append(studentClasses, *models.NewStudentClass(currIdStudent-1, idxClass, time.Now()))
	}

	return &students, &classes, &studentClasses, &majors, nil
}

func searchClassFromStirng(year int, section string, idxMajor int, classes *[]models.Class, majors *[]models.Majors) (int, error) {
	for idx, c := range *classes {
		if c.Year == year && c.Section == section && (*majors)[idxMajor].Name == getMajorFormIdMajor(c.IdMajor, majors) {
			return idx, nil
		}
	}
	return -1, nil
}

func getMajorFormIdMajor(id int, majors *[]models.Majors) string { // Todo: ricordare di aggiungere gli id algi oggetti quando li dichiaro
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
