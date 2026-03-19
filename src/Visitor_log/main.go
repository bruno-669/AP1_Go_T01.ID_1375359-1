package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	PatientNotFoundError = errors.New("patient not found")
)

type Visits struct {
	ID        uint64
	VisitorID uint64
	DoctorID  uint64
	Date      time.Time
}
type IdCount struct {
	VisitsId  uint64
	DoctorsId uint64
	PersonId  uint64
}

func ErrMassage(message string) error {
	if message == "" {
		return errors.New("invalid input")
	}
	return errors.New(message)
}

func AddPersonIsNotExist(PersonName string, PersonList map[string]uint64, Counting *IdCount) {
	if _, exist := PersonList[PersonName]; exist == false {
		PersonList[PersonName] = Counting.PersonId
		Counting.PersonId++
	}
}

func AddDoctorIsNotExist(DoctorName string, DoctorsList map[string]uint64, Counting *IdCount) {
	if _, exist := DoctorsList[DoctorName]; exist == false {
		DoctorsList[DoctorName] = Counting.DoctorsId
		Counting.DoctorsId++
	}
}

func CreateNewVisit(Counting *IdCount, PersonID, DoctorID uint64, date time.Time) Visits {
	NewVisit := Visits{
		ID:        Counting.VisitsId,
		VisitorID: PersonID,
		DoctorID:  DoctorID,
		Date:      date,
	}
	Counting.VisitsId++
	return NewVisit
}

func SerchName(PersonList map[string]uint64, id uint64) string {
	for name, v := range PersonList {
		if v == id {
			return name
		}
	}
	return ""
}

func PrintAllVisitList(VisitList []Visits, DoctorsList, PersonList map[string]uint64) string {
	var builder strings.Builder
	for _, visit := range VisitList {
		builder.WriteString("[" + strconv.Itoa(int(visit.ID)) + "][" + strconv.Itoa(int(visit.VisitorID)) + "][" + strconv.Itoa(int(visit.DoctorID)) + "][" + SerchName(PersonList, visit.VisitorID) + "][" + SerchName(DoctorsList, visit.DoctorID) + "][" + visit.Date.Format("2006-01-02") + "]\n")
	}
	return builder.String()
}

func SaveVisitor(input io.Reader, VisitList *[]Visits, DoctorsList, PersonList map[string]uint64, Counting *IdCount) error {
	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		return ErrMassage("")
	}
	PersonName := scanner.Text()
	AddPersonIsNotExist(PersonName, PersonList, Counting)

	if !scanner.Scan() {
		return ErrMassage("")
	}

	DoctorName := scanner.Text()

	AddDoctorIsNotExist(DoctorName, DoctorsList, Counting)

	if !scanner.Scan() {
		return ErrMassage("")
	}
	datestr := scanner.Text()
	date, err := time.Parse("2006-01-02", datestr)
	if err != nil {
		return ErrMassage("Error Invalid Date")
	}

	*VisitList = append(*VisitList, CreateNewVisit(Counting, PersonList[PersonName], DoctorsList[DoctorName], date))

	return nil
}

func CheckEndMassage(massage string) bool {
	switch strings.ToLower(massage) {
	case "exit", "end", "esc", "q":
		return true
	}
	return false
}

func ParseString(line string, VisitList *[]Visits, DoctorsList, PersonList map[string]uint64, Counting *IdCount) error {
	var UserName, UserLastName, UserFirstName, DoctorName, datestr string
	var date time.Time
	if _, err := fmt.Sscanf(line, "%s %s %s %s %s", &UserName, &UserLastName, &UserFirstName, &DoctorName, &datestr); err != nil {
		return ErrMassage("Invalid Format Line In file")
	}
	date, err := time.Parse("2006-01-02", datestr)
	if err != nil {
		return ErrMassage("Invalid Date Format In file")
	}
	PersonName := UserName + " " + UserLastName + " " + UserFirstName
	AddPersonIsNotExist(PersonName, PersonList, Counting)
	AddDoctorIsNotExist(DoctorName, DoctorsList, Counting)
	*VisitList = append(*VisitList, CreateNewVisit(Counting, PersonList[PersonName], DoctorsList[DoctorName], date))
	return nil
}

func AddInFile(input io.Reader, VisitList *[]Visits, DoctorsList, PersonList map[string]uint64, Counting *IdCount) (string, error) {
	scanner := bufio.NewScanner(input)
	filename := ""
	for {
		if !scanner.Scan() {
			return "", ErrMassage("")
		}
		text := scanner.Text()
		if text == "" {
			text = "data.txt"
		}

		if CheckEndMassage(text) {
			return "", nil
		}
		file, err := os.Open(text)
		if err != nil {
			fmt.Println("Invalid File name")
			continue
		}
		filename = text
		file.Close()
		break
	}
	file, _ := os.Open(filename)
	defer file.Close()
	FileScanner := bufio.NewScanner(file)
	var allcount, passcount int

	for FileScanner.Scan() {
		allcount++
		if ParseString(FileScanner.Text(), VisitList, DoctorsList, PersonList, Counting) != nil {
			continue
		}
		passcount++
	}

	return "Export " + strconv.Itoa(passcount) + " in " + strconv.Itoa(allcount) + " file data", nil
}
func InputName(input io.Reader) string {
	scanner := bufio.NewScanner(input)
	name := "'"
	for {
		if !scanner.Scan() {
			return ""
		}
		name = scanner.Text()
		if name == "" {
			fmt.Println("Invalid name")
			continue
		}
		break
	}
	return name
}

func SaveFile(input io.Reader, VisitList []Visits, DoctorsList, PersonList map[string]uint64) error {
	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		return ErrMassage("")
	}
	filename := scanner.Text()
	if filename == "" {
		filename = "data.txt"
	}
	file, err := os.Create(filename)
	if err != nil {
		return ErrMassage("Error Save File")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, visit := range VisitList {
		writer.WriteString(SerchName(PersonList, visit.VisitorID) + " " + SerchName(DoctorsList, visit.DoctorID) + " " + visit.Date.Format("2006-01-02") + "\n")
	}
	writer.Flush()
	return nil
}

func GetHistoryVisitor(input io.Reader, VisitList []Visits, DoctorsList, PersonList map[string]uint64) (string, error) {
	var builder strings.Builder
	username := InputName(input)
	_, exist := PersonList[username]
	if !exist {
		return "", PatientNotFoundError
	}
	count := 0
	for _, visit := range VisitList {
		if visit.VisitorID == PersonList[username] {
			builder.WriteString(SerchName(DoctorsList, visit.DoctorID) + " " + visit.Date.Format("2006-01-02") + "\n")
			count++
		}
	}
	if count == 0 {
		return "", PatientNotFoundError
	}
	return builder.String(), nil
}

func GetLastVisit(input io.Reader, VisitList []Visits, DoctorsList, PersonList map[string]uint64) (string, error) {
	username := InputName(input)
	_, exist := PersonList[username]
	if !exist {
		return "", PatientNotFoundError
	}
	doctorname := InputName(input)
	SortDate := []time.Time{}

	for _, visit := range VisitList {
		if visit.VisitorID == PersonList[username] && visit.DoctorID == DoctorsList[doctorname] {
			SortDate = append(SortDate, visit.Date)
		}
	}
	if len(SortDate) == 0 {
		fmt.Println("Records Not Found")
		return "", nil
	}
	sort.Slice(SortDate, func(i, j int) bool {
		return SortDate[j].Before(SortDate[i])
	})
	return SortDate[0].Format("2006-01-02") + "\n", nil
}

func Handler(comand string, input io.Reader, VisitList *[]Visits, DoctorsList, PersonList map[string]uint64, Counting *IdCount) (string, error) {
	switch strings.ToLower(comand) {
	case "save", "s":
		if err := SaveVisitor(input, VisitList, DoctorsList, PersonList, Counting); err != nil {
			return "", err
		}
	case "gethistory", "ghs":
		out, err := GetHistoryVisitor(input, *VisitList, DoctorsList, PersonList)
		if err != nil {
			return "", err
		}
		return out, nil
	case "getlastvisit", "glv":
		out, err := GetLastVisit(input, *VisitList, DoctorsList, PersonList)
		if err != nil {
			return "", err
		}
		return out, nil
	case "addfile", "a":
		out, err := AddInFile(input, VisitList, DoctorsList, PersonList, Counting)
		if err != nil {
			return "", err
		}
		return out, nil
	case "printall", "p":
		return PrintAllVisitList(*VisitList, DoctorsList, PersonList), nil
	case "savefile":
		SaveFile(input, *VisitList, DoctorsList, PersonList)
		return "", nil
	default:
		return "", ErrMassage("Error type operation")
	}
	return "", nil
}

func ComandScan(input io.Reader) (string, error) {
	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		return "", ErrMassage("Error scan input")
	}
	return scanner.Text(), nil
}

func main() {
	VisitList := []Visits{}
	DoctorsList := make(map[string]uint64)
	PersonList := make(map[string]uint64)
	var Counting IdCount

	for {
		var err error
		comand := ""
		fmt.Println("enter the command:")
		if comand, err = ComandScan(os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else if CheckEndMassage(comand) {
			break
		} else {
			if text, err := Handler(comand, os.Stdin, &VisitList, DoctorsList, PersonList, &Counting); err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Print(text)
			}

		}
	}

}
