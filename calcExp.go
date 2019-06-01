package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readAndConvLevel(input *bufio.Reader) float64 {
	var levelArray = []float64{
		0,
		1000,
		2136.3636363636365,
		3427.6859504132235,
		4895.097670924118,
		6562.610989686497,
		8457.51248828011,
		10610.809645772853,
		13057.738233832788,
		15838.338902082713,
		18998.112388730355,
		22588.764078102675,
		26669.05008875304,
		31305.738737219363,
		36574.70311047655,
		42562.16262554153,
		49366.09389266083,
		57097.83396893276,
		65883.9022374236,
		75868.07072434499,
		87213.71673221022,
		100106.49628660252,
		114757.3821438665,
		131406.11607257556,
		150325.13190065406,
		171824.01352347052,
		196254.56082212558,
		224016.54638877907,
		255564.2572599762,
		291413.9287045184,
		332152.19170968,
		-1}
	fmt.Println("Enter your level:")
	in, err := input.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	rawLevel, err := strconv.ParseFloat(strings.Trim(in, "\n\r"), 64)
	if err != nil {
		log.Fatal(err)
	}
	level, decimal := math.Modf(rawLevel)
	totalExp := levelArray[int(level)]
	diff := levelArray[int(level)+1] - levelArray[int(level)]
	totalExp += diff * decimal
	return totalExp
}

func readAndConvProj(input *bufio.Reader) float64 {
	var tierArray = []float64{
		2.2,
		8.8330,
		20.02,
		33.88,
		56.10,
		79.2,
		93.50,
		132}

	fmt.Println("Enter the tier of project:")
	in, err := input.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	tier, err := strconv.ParseInt(strings.Trim(in, "\n\r"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enter mark:")
	in, err = input.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	mark, err := strconv.ParseInt(strings.Trim(in, "\n\r"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return tierArray[tier] * float64(mark)
}

func convToLevel(exp float64) float64 {
	var levelArray = []float64{
		0,
		1000,
		2136.3636363636365,
		3427.6859504132235,
		4895.097670924118,
		6562.610989686497,
		8457.51248828011,
		10610.809645772853,
		13057.738233832788,
		15838.338902082713,
		18998.112388730355,
		22588.764078102675,
		26669.05008875304,
		31305.738737219363,
		36574.70311047655,
		42562.16262554153,
		49366.09389266083,
		57097.83396893276,
		65883.9022374236,
		75868.07072434499,
		87213.71673221022,
		100106.49628660252,
		114757.3821438665,
		131406.11607257556,
		150325.13190065406,
		171824.01352347052,
		196254.56082212558,
		224016.54638877907,
		255564.2572599762,
		291413.9287045184,
		332152.19170968,
		-1}
	lvl := float64(sort.Search(len(levelArray), func(i int) bool { return levelArray[i] >= exp }))
	diff := levelArray[int(lvl)+1] - levelArray[int(lvl)]
	expLeft := exp - levelArray[int(lvl)]
	lvl += expLeft / diff
	return lvl
}

func main() {
	input := bufio.NewReader(os.Stdin)
	exp := readAndConvLevel(input)
	point := readAndConvProj(input)
	fmt.Printf("your level will be:\n%f", convToLevel(exp+point))
}

//sal, example of sales manager.
