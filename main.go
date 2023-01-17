package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

func printSpeed() string {
	return "speed is 10"
}

func splitUrlPath(file_path string) (string, string) {
	// if you don't know the initial value of a variable, use declaration
	// score := 0 no -> use var score init as it's already 0
	// if it's a package level variable -> declare!
	// group variables together to enhance readability
	// most of time we prefer using short decleration
	dir, file := path.Split(file_path)
	return dir, file
}

// function delaration

func checkValidUser() {
	const (
		username = "jack"
		password = "1888"
	)

	var InputUsrName, InputUsrPwd string

	if len(os.Args) != 3 {
		println("Usage: [username] [password]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	for _, s := range os.Args {
		fmt.Println(s)
	}

	InputUsrName, InputUsrPwd = os.Args[1], os.Args[2]

	if InputUsrName != username {
		fmt.Printf("Access denied for %q.\n", InputUsrName)
	} else if InputUsrName == username && InputUsrPwd != password {
		fmt.Printf("Invalid password for %q.", username)
	} else {
		fmt.Printf("Access granted to %q.", username)
	}

}

func convertFeetToMeter() {
	feet := os.Args[1]
	f, err := strconv.ParseFloat(feet, 64)
	if err == nil {
		meter := f * 0.304
		fmt.Printf("Converted feet %s to meter %f\n", feet, meter)
	} else {
		fmt.Println("Error during conversion: ", err)
	}

	// using simple statement
	// simple statement variable is only visible to the if statement
	if n, err := strconv.Atoi("52"); err == nil {
		fmt.Println(n)
	} else {
		fmt.Println(err)
	}

	//fmt.Println(n)

}

func shortStatementTest() {
	var n int
	var err error

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number")
		return
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot convert %q\n", a[1])
		return
	} else {
		fmt.Printf("%d * 2 is %d \n", n, n*2)
	}
	// being shadowed!
	fmt.Println(n, err)

}

func switchCity() {

	city := "Lyon"
	switch city {
	default:
		fmt.Println("Can't find")
	case "Paris", "Lyon":
		fmt.Println("France")
		// case block creates it's own space
		// variables are only visible at case level
	case "Tokyo":
		fmt.Println("China")
	}

}

func switchBool() {

	i := -100

	switch true {
	case i > 0:
		fmt.Println("Pos")
	case i < 0:
		fmt.Println("Neg")
	default:
		fmt.Println("zero")
	}

}

func switchFallThrough() {

	switch i := 1000; true {
	case i > 100:
		fmt.Println("Big")
		fallthrough
	case i > 0:
		fmt.Println("Pos")
		fallthrough
	default:
		fmt.Println("number")
	}

}

func partsOfDay() {
	switch current_hour := time.Now().Hour(); true {
	case current_hour >= 18:
		fmt.Println("good evening")
	case current_hour >= 12:
		fmt.Println("good afternoon")
	case current_hour >= 6:
		fmt.Println("good morning")
	default:
		fmt.Println("good night")
	}

}

func randomNumberGame() {
	const maxTurn = 10
	rand.Seed(time.Now().UnixNano())
	guess := 10
	for turn := 0; turn < maxTurn; turn++ {
		n := rand.Intn(guess + 1)
		if n == guess {
			fmt.Println("You win!")
			return
		}
	}
	fmt.Println("You Lost!")

}

func wordFinder() {
	const corpus = "" +
		"lazy cat jumps again and again and again"
	words := strings.Fields(corpus)
	query := os.Args[1:]
queries:
	for _, q := range query {
		for i, v := range words {
			if q == v {
				fmt.Println(i, v)
				// remove duplicates
				break queries
			}
		}
	}

}

func multiArray() {
	books := [...][3]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	sum := 0.0
	count := 0.0
	for _, grades := range books {
		for _, grade := range grades {
			sum += grade
			count++
		}
	}
	fmt.Println(sum / count)

}

func nammedtype() {
	// unnamed type
	// var books [3]int
	// named type
	// unnamed type can be compare to named type as long as underlying structure is the same
	// named type can be only compare to each other with same name
	type books3 [3]string

	var mybooks books3

	mybooks[0] = "a"
	fmt.Println(mybooks)

}

func fixProblem() {

	names := [...]string{
		"Einstein", "Shepard",
		"Tesla",
	}

	books := [3]string{
		"Stay Golden",
		"Fire",
		"Kafka's Revenge",
	}

	sort.Strings(books[:])

	// this time, do not change the nums array to a slice
	nums := [...]int{5, 1, 7, 3, 8, 2, 6, 9}

	// use the slicing expression to change the nums array to a slice below
	sort.Ints(nums[:])

	fmt.Printf("%q\n", names)

	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}

func sliceAppend() {
	nums := []int{}
loop:
	for i := 0; i <= 10; i++ {
		nums = append(nums, i)
		if i == 5 {
			break loop
		}
	}
	nums2 := []int{5, 5, 5}
	nums = append(nums, nums2...)

	for i, val := range nums {
		fmt.Println(i, val)
	}
}

func append2() {
	toppings := []string{"black olives", "green peppers"}

	var pizza []string
	pizza = append(pizza, toppings...)
	pizza = append(pizza, "onions", "extra cheese")

	fmt.Printf("pizza       : %s\n", pizza)
}

func sortNum() {
	input := os.Args
	numbers := []int{}
	if len(input) == 1 {
		fmt.Println("provide a few numbers")
		return
	}
	for _, i := range input {
		n, err := strconv.Atoi(i)
		if err == nil {
			numbers = append(numbers, n)
		} else {
			continue
		}
	}
	sort.Ints(numbers)
	fmt.Println(numbers)
}

func main() {

}
