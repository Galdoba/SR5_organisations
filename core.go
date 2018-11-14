package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
)

//TakeOptions - takes Q, slice of A and returns number of chosen A and string of that A
func TakeOptions(question string, options ...string) (int, string) {
	fmt.Println(question)
	for i := range options {
		prefix := "[" + strconv.Itoa(i+1) + "] - "
		fmt.Println(prefix + options[i])
	}
	answer := 0
	gotIt := false
	for !gotIt {
		answer = InputInt()
		if answer < len(options)+1 && answer > 0 {
			gotIt = true
		} else {
			fmt.Println("Answer is incorrect...")
			fmt.Println(question)
		}
	}
	return answer, options[answer-1]
}

func describe(descr []string) {
	if len(descr) > 0 {
		fmt.Print(descr[0])
	}
}

//InputInt - takes Integer from User
func InputInt(descr ...string) int {
	describe(descr)
	var dataVal int
	fmt.Scan(&dataVal)
	return dataVal
}

//InputFloat64 - takes Float64 from User
func InputFloat64(descr ...string) float64 {
	describe(descr)
	var dataVal float64
	fmt.Scan(&dataVal)
	return dataVal
}

//InputString - takes Float64 from User
func InputString(descr ...string) string {
	describe(descr)
	var dataVal string
	fmt.Scan(&dataVal)
	return dataVal
}

//Str2Float64 - convert String to Float64
func Str2Float64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}

//Str2Int -
func Str2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

//Float64ToStr -
func Float64ToStr(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 0, 64)
}

//ClearScreen - clearing comand console (for Windows)
func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func askYesNo(str string) bool {
	gotAnswer := false
	for !gotAnswer {
		fmt.Print(str + "(y/n) ")
		answer := InputString()
		switch answer {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Println("Error: Answer is incorrect. (Type 'y' or 'n')")
		}
	}
	return false

}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	//test
	return float64(round(num*output)) / output
}

func randInt(min int, max int) int {
	return min + rand.Intn(max)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func sr5SimpleTest(val0 ...int) (result string, glitch string, hits int) {
	val := val0
	dp := 0
	limit := 999
	tn := 1
	for i := range val {
		if i == 0 {
			dp = val[i]
		}
		if i == 1 {
			limit = val[i]
		}
		if i == 2 {
			tn = val[i]
		}
		if i > 2 {

			break
		}
	}
	var outarray []int
	for i := dp; i > 0; i-- {
		dice := randInt(1, 6)
		outarray = append(outarray, dice)
	}
	var ones int
	for i := range outarray {
		switch outarray[i] {
		case 1:
			ones++
		case 5:
			hits++
		case 6:
			hits++
		default:
		}
	}
	if hits > limit {
		hits = limit
	}
	if hits >= tn {
		result = "SUCCESS!"
	} else {
		result = "FAILURE!"
	}
	if len(outarray)/2 < ones {
		glitch = "GLITCH!"
		if hits == 0 {
			glitch = "CRITICAL GLITCH!!!"
		}
	}
	return result, glitch, hits
}

func analyzeSR5SimpleTest(val0 ...int) (totalSuc int, totalGl int, averageHits float64) {
	var totalHits float64
	val := val0
	dp := 0
	limit := 999
	tn := 1
	for i := range val {
		if i == 0 {
			dp = val[i]
		}
		if i == 1 {
			limit = val[i]
		}
		if i == 2 {
			tn = val[i]
		}
		if i > 2 {

			break
		}
	}
	for i := 0; i < 1000000; i++ {

		result, glitch, hits := sr5SimpleTest(dp, limit, tn)
		if result == "SUCCESS!" {
			totalSuc++
		}
		if glitch != "" {
			totalGl++
		}
		totalHits = totalHits + float64(hits)

	}
	totalSuc = totalSuc / 10000
	averageHits = totalHits / 1000000
	fmt.Println("Test:", strconv.Itoa(dp)+"["+strconv.Itoa(limit)+"] ("+strconv.Itoa(tn)+")")
	fmt.Println("RESULT:", "totalSuc =", totalSuc, "   |||   Hits(average) =", averageHits)
	return totalSuc, totalGl, averageHits
}

func shuffleStringSlice(slice []string) []string {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
