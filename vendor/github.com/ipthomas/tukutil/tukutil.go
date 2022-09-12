package tukutil

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
)

var (
	SeedRoot   = "1.2.40.0.13.1.1.3542466645."
	IdSeed     = getIdIncrementSeed(5)
	CodeSystem = make(map[string]string)
)

// TemplateFuncMap returns a functionMap of tukutils for use in templates
func TemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"dtday":      Tuk_Day,
		"dtmonth":    Tuk_Month,
		"dtyear":     Tuk_Year,
		"mappedid":   GetCodeSystemVal,
		"prettytime": PrettyTime,
		"newuuid":    NewUuid,
		"newid":      Newid,
		"tuktime":    Tuk_Time,
	}
}

// SplitXDWKey takes a string input (xdw key) and returns the pathway and nhs id for the xdw
func SplitXDWKey(xdwkey string) (string, string) {
	var pwy string
	var nhs string
	if len(xdwkey) > 10 {
		pwy = xdwkey[:len(xdwkey)-10]
		nhs = strings.TrimPrefix(xdwkey, pwy)
	}
	log.Printf("Pathway = %s NHS ID = %s", pwy, nhs)
	return pwy, nhs
}

// SetCodeSystem takes a map input and sets the codesystem map with the input
func SetCodeSystem(cs map[string]string) {
	CodeSystem = cs
}

// InitCodeSystem loads the codesystem json file and sets the codesystem map from the json file values
func InitCodeSystem(codesystemFile string) error {
	file, err := os.Open(codesystemFile)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if err = json.NewDecoder(file).Decode(&CodeSystem); err != nil {
		log.Println(err.Error())
		return err
	}
	log.Printf("Loaded %v code system key values", len(CodeSystem))
	Log(CodeSystem)
	return nil
}
func Log(i interface{}) {
	b, _ := json.MarshalIndent(i, "", "  ")
	log.Println(string(b))
}

// GetCodeSystemVal takes a string input (key) and returns the string value corresponding to the input (key) from the codesystem
func GetCodeSystemVal(key string) string {
	val, ok := CodeSystem[key]
	if ok {
		return val
	}
	return key
}

// returns unique id in format '1.2.40.0.13.1.1.3542466645.20211021090059143.32643'
// idroot constant - 1.2.40.0.13.1.1.3542466645.
// + datetime	   - 20211021090059143.
// + 5 digit seed  - 32643
// if state is maintained the seed is incremented after each call to newid() to ensure a unique id is generated.
// If state is not maintained the `new` datetime will ensure a unique id is generated.
func Newid() string {
	id := SeedRoot + dt_yyyyMMddhhmmSSsss() + "." + GetStringFromInt(IdSeed)
	IdSeed = IdSeed + 1
	return id
}

// Tuk_Pretty_Time returns a pretty version of the current time for location Europe/London (strips everything after the `.` in Tuk_Time)
func TUK_Pretty_Time() string {
	return PrettyTime(Tuk_Time())
}

// Tuk_Time returns the current time for location Europe/London.
func Tuk_Time() string {
	location, err := time.LoadLocation("Europe/London")
	if err != nil {
		log.Println(err.Error())
		return time.Now().String()
	}
	return time.Now().In(location).String()
}

// PrettyTime strips everything after the `.` from the input time string
func PrettyTime(time string) string {
	return strings.TrimSuffix(strings.ReplaceAll(strings.Split(time, ".")[0], "T", " "), "Z")
}

// TUK_Hour returns the current hour as a 2 digit string
func Tuk_Hour() string {
	return fmt.Sprintf("%02d",
		time.Now().Local().Hour())
}

// TUK_Min returns the current minute as a 2 digit string
func Tuk_Min() string {
	return fmt.Sprintf("%02d", time.Now().Local().Minute())
}

// TUK_Sec returns the current second as a 2 digit string
func Tuk_Sec() string {
	return fmt.Sprintf("%02d",
		time.Now().Local().Second())
}

// TUK_MilliSec returns the current milliseconds as a 3 digit int
func Tuk_MilliSec() int {
	return GetIntFromString(Substr(GetStringFromInt(time.Now().Nanosecond()), 0, 3))
}

// TUK_Day returns the current day as a 2 digit string
func Tuk_Day() string {
	return fmt.Sprintf("%02d",
		time.Now().Local().Day())
}

// TUK_Year returns the current year as a 4 digit string
func Tuk_Year() string {
	return fmt.Sprintf("%d",
		time.Now().Local().Year())
}

// TUK_Month returns the current month as a 2 digit string
func Tuk_Month() string {
	return fmt.Sprintf("%02d",
		time.Now().Local().Month())
}

// NewUuid returns a random UUID as a string
func NewUuid() string {
	u := uuid.New()
	return u.String()
}

// GetStringFromInt takes a int input and returns a string of that value.
func GetStringFromInt(i int) string {
	return strconv.Itoa(i)
}

// GetIntFromString takes a string input with an integer value and returns an int of that value. If the input is not numeric, 0 is returned
func GetIntFromString(input string) int {
	i, _ := strconv.Atoi(input)
	return i
}

// Substr takes a string input and returns the rune (Substring) defined by the start and length in th start and length input values
func Substr(input string, start int, length int) string {
	asRunes := []rune(input)
	if start >= len(asRunes) {
		return ""
	}
	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}
	return string(asRunes[start : start+length])
}

// GetXMLNodeList takes an xml message as input and returns the xml node list as a string for the node input value provide
func GetXMLNodeList(message string, node string) string {
	if strings.Contains(message, node) {
		var nodeopen = "<" + node
		var nodeclose = "</" + node + ">"
		log.Println("Searching for XML Element: " + nodeopen + ">")
		var start = strings.Index(message, nodeopen)
		var end = strings.Index(message, nodeclose) + len(nodeclose)
		m := message[start:end]
		log.Println("Extracted XML Element Nodelist")
		return m
	}
	log.Println("Message does not contain Element : " + node)
	return ""
}

// PrettyAuthorInstitution takes a string input (XDS Author.Institution format) and returns a string with just the Institution name
func PrettyAuthorInstitution(institution string) string {
	if strings.Contains(institution, "^") {
		return strings.Split(institution, "^")[0] + ","
	}
	return institution
}

// PrettyAuthorPerson takes a string input (XDS Author.Person format) and returns a string with the person last and first names
func PrettyAuthorPerson(author string) string {
	if strings.Contains(author, "^") {
		authorsplit := strings.Split(author, "^")
		if len(authorsplit) > 2 {
			return authorsplit[1] + " " + authorsplit[2]
		}
		if len(authorsplit) > 1 {
			return authorsplit[1]
		}
	}
	return author
}

// GetFolderFiles takes a string input of the complete folder path and returns a fs.DirEntry
func GetFolderFiles(folder string) ([]fs.DirEntry, error) {
	var err error
	var f *os.File
	var fileInfo []fs.DirEntry
	f, err = os.Open(folder)
	if err != nil {
		log.Println(err)
		return fileInfo, err
	}
	fileInfo, err = f.ReadDir(-1)
	f.Close()
	if err != nil {
		log.Println(err)
	}
	return fileInfo, err
}
func getIdIncrementSeed(len int) int {
	return GetIntFromString(Substr(GetStringFromInt(time.Now().Nanosecond()), 0, len))
}
func dt_yyyyMMddhhmmSSsss() string {
	return Tuk_Year() + Tuk_Month() + Tuk_Day() + Tuk_Hour() + Tuk_Min() + Tuk_Sec() + strconv.Itoa(Tuk_MilliSec())
}
