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

// func SwapStudentId(id int, student models.Student, local_studentclass *[]models.StudentClass) {
// 	for _, e := range *local_studentclass {
// 		if student.Id == e.IdS {
// 			e.IdS = id
// 		}
// 	}
// 	student.Id = id
// }
// func SwapClassId(id int, class models.Class, local_studentclass *[]models.StudentClass) {
// 	for _, e := range *local_studentclass {
// 		if class.Id == e.IdC {
// 			e.IdS = id
// 		}
// 	}
// 	class.Id = id
// }
// func SwapMajorId(id int, major models.Majors, local_classes *[]models.Class) {
// 	for _, c := range *local_classes {
// 		if major.Id == c.IdMajor {
// 			c.IdMajor = id
// 		}
// 	}
// 	major.Id = id
// }
// func AddDataToDb(local_students *[]models.Student, local_classes *[]models.Class, local_studentClasses *[]models.StudentClass, local_majors *[]models.Majors) error {
// 	// todo: implementare una sorta di lock (se ho voglia)
// 	for _, m := range *local_majors {
// 		m.Id = majors.GetLastId()
// 		err := m.Save()
// 		if err != nil {
// 			return err
// 		}
// 		SwapMajorId(majors.GetLastId(), m, local_classes)
// 	}
// 	for _, s := range *local_students {
// 		err := s.Save()
// 		if err != nil {
// 			return err
// 		}
// 		SwapStudentId(students.GetLastId(), s, local_studentClasses)
// 	}
// 	for _, c := range *local_classes {
// 		err := c.Save()
// 		if err != nil {
// 			return err
// 		}
// 		SwapClassId(classes.GetLastId(), c, local_studentClasses)
// 	}
// 	// finisci questo ????
// 	for _, sc := range *local_studentClasses {
// 		err := sc.Save()
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

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
	// studs := make([]models.Student, 1)
	clss := classes.GetAll()
	// studClss := make([]models.StudentClass, 1)
	mjrs := majors.GetAll()

	currIdStudent := students.GetLastId() + 1
	currIdClass := classes.GetLastId() + 1
	currIdMajor := majors.GetLastId() + 1

	for {
		record, err := reader.Read()
		log.Println(record)
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

		student := models.Student{Name: record[0], LastName: record[1], Id: currIdStudent}
		student.Save()
		// studs = append(studs, student)
		currIdStudent++

		classInfo := strings.Split(record[2], " ")
		year := int(classInfo[0][0]) - 48
		if year < 0 || year > 5 {
			return fmt.Errorf("il formato del file non e' corretto, Anno errato")
		}

		log.Printf("ANNO CLASSE [%d] \n", year)

		section := classInfo[0][:1]
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
		// studClss = append(studClss, studCls)
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
