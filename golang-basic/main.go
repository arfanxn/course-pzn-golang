package main

import (
	"container/list"
	"container/ring"
	"flag"
	"fmt"
	"golang-basic/database"
	"golang-basic/helper"
	"math"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	// variable declarations
	const name string = "Muhammad arfan"
	var (
		gender = "male"
		age    = "20"
	)
	// prevent warning by unused variables
	_ = gender
	_ = age

	// Display to console
	fmt.Println(name)
	fmt.Println("Length of name is :", len(name))

	/* Data Type conversions */
	var number32 int32 = 10000
	var number16 int16 = int16(number32)
	println("Data Type conversion :", number16)

	/* Array */
	var employees = [10]string{
		"jack",
		"dorsey",
	}
	_ = employees

	/* Array and slices */
	arrayOfMonths := [...]string{
		"Jan", "Feb", "Mar", "Apr",
		"May", "Jun", "Jul", "Aug",
		"Sep", "Oct", "Nov", "Dec",
	}
	sliceOfMonths := arrayOfMonths[4:9]
	fmt.Println("Slice of months :", sliceOfMonths)
	fmt.Println("Capacity of \"Slice of months\" is", cap(sliceOfMonths))

	sliceOfDays := make([]string, 2, 7)
	sliceOfDays[0] = "Wednesday"
	sliceOfDays[1] = "Thursday"
	fmt.Println("Slice of days :", sliceOfDays)
	fmt.Println("Len and cap :", len(sliceOfDays), cap(sliceOfDays))
	// Copy "sliceOfDays" to new variable
	copiedSliceOfDays := make([]string, len(sliceOfDays), cap(sliceOfDays))
	copy(copiedSliceOfDays, sliceOfDays)
	fmt.Println(copiedSliceOfDays)

	/* Different between Array and Slice declaration */
	arrayDeclaration := [...]string{"Hello", "World"}
	sliceDeclaration := []string{"Hello", "World"}
	fmt.Println("Declared slice len and cap :",
		len(sliceDeclaration), cap(sliceDeclaration))
	_ = arrayDeclaration

	/* Map data type */
	mapCar := map[string]string{
		"model": "E36 320i",
		"brand": "BMW",
	}
	mapCar["price"] = "EUR 60000"
	fmt.Println(mapCar["model"])
	fmt.Println(mapCar)
	delete(mapCar, "price")
	fmt.Println(mapCar)

	/* If statements */
	if false {
		fmt.Println("not executed")
	} else if length := len("Hello"); length > 0 {
		fmt.Println("lenght is more than ", length)
	}

	/* Switch statements */
	switch length := len("Hello"); length > 0 {
	case true:
		fmt.Println("lenght is more than ", length)
	default:
		fmt.Println("lenght is less than ", length)
	}

	/* For loops */
	for i := 0; i < 5; i++ {
		fmt.Println("Index loop ", i)
	}
	for key, car := range mapCar {
		fmt.Println("Property :", key, "=", car)
	}

	/* Functions */
	_ = sayHello("Arfan")
	firstName, _ := getFullname()
	_ = firstName
	// Variadic function
	numbers := []int32{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8))
	fmt.Println(sumAll(numbers...))

	/* Type assertions */
	var typeString interface{} = "Hello world"
	fmt.Println("Value of \"typeString\" :", typeString.(string))
	// fmt.Println("Value data type of \"typeString\" :", typeString.(type))

	/* Using imported custom package */
	helper.SayMorning()
	fmt.Println(database.GetMysql())

	/* Communicate with OS via Go-Lang Package OS */
	osArguments := os.Args
	osHostname, _ := os.Hostname()
	fmt.Println("List of OS arguments : ")
	fmt.Println(osArguments)
	fmt.Println("OS Hostname : " + osHostname)
	fmt.Println("OS Environment - USERNAME : " + os.Getenv("USERNAME"))

	/* Package FLAG */
	host := flag.String("host", "localhost", "Put your host")
	flag.Parse()
	fmt.Println("Package FLAG host :", *host)

	/* Package String */
	fmt.Println("Result of string comparison :", strings.Compare("ass", "ssa"))

	/* Package String Conversion */
	parsedStrToBool, _ := strconv.ParseBool("true")
	fmt.Println("Result of \"parsedStrToBool\" :", parsedStrToBool)

	/* Package Math */
	fmt.Println("Go-Lang Maxiumun 64Int :", math.MaxInt64)

	/* Package Container List */
	fullnameLinkList := list.New()
	fullnameLinkList.PushBack("Muhammad")
	fullnameLinkList.PushBack("Arfan")

	fmt.Println("Package Container List :")
	fmt.Println("\"fullnameLinkList\" Front : ", fullnameLinkList.Front().Value)
	fmt.Println("\"fullnameLinkList\" Back : ", fullnameLinkList.Back().Value)

	for name := fullnameLinkList.Front(); name != nil; name = name.Next() {
		fmt.Println(name.Value)
	}

	/* Package Container Ring */
	var numbersLinkRing *ring.Ring = ring.New(10)
	for i := 0; i < numbersLinkRing.Len(); i++ {
		numbersLinkRing.Value = strconv.FormatInt(int64(i*100), 10)
		numbersLinkRing.Next()
	}
	numbersLinkRing.Do(func(number interface{}) {
		fmt.Println(number)
	})

	/* Package Time */
	now := time.Now()
	fmt.Println("Time Now :", now)
	fmt.Println("Day Now :", now.Day())

	timeLayout := "2006-01-02"
	parsedTime, _ := time.Parse(timeLayout, "2020-10-10")
	fmt.Println("Parsed time :", parsedTime)

	/* Package Reflect */
	personReflect := Person{name: "Jack", age: 21}
	personReflectType := reflect.TypeOf(personReflect)
	fmt.Println("personReflectType :", personReflectType)
	fmt.Println("personReflectType Field :", personReflectType.Field(0).Name)

	/* Package Regexp */
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])n")
	fmt.Println("Is regex match string :", regex.MatchString("arn"))
	

}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

func sayHello(name string) string {
	fmt.Println("hello", name)
	return "hello " + name
}

func getFullname() (firstName string, lastName string) {
	return "Muhammad", "Arfan"
}

func sumAll(numbers ...int32) int32 {
	var result int32 = 0
	for _, number := range numbers {
		result += number
	}
	return result
}

type Person struct {
	name string
	age  uint8
}
