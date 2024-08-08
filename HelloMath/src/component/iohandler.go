package component

import (
	"fmt"
	"strconv"
	"strings"
)

type Question struct {
	thisRelation [][2]int
	statement1   string
	statement2   string
	solution     string
	answer       string
	questionType int
	parseType    int
	level        int
}

func setOutput(slice []int) string {
	var setString []string
	for _, value := range slice {
		setString = append(setString, strconv.Itoa(value))
	}
	return "{" + strings.Join(setString, ", ") + "}"
}

func relationOutput(relation [][2]int) string {
	var relationString []string
	for _, value := range relation {
		singleRelation := "(" + strconv.Itoa(value[0]) + ", " + strconv.Itoa(value[1]) + ")"
		relationString = append(relationString, singleRelation)
	}
	return "{" + strings.Join(relationString, ", ") + "}"
}
func relationAlpOutput(relation [][2]string) string {
	var relationString []string
	for _, value := range relation {
		singleRelation := "(" + value[0] + ", " + value[1] + ")"
		relationString = append(relationString, singleRelation)
	}
	return "{" + strings.Join(relationString, ", ") + "}"
}
func interfaceTwoOutput(product [][]interface{}) string {
	var productString []string
	for _, pair := range product {
		var pairString []string
		for _, value := range pair {
			pairString = append(pairString, fmt.Sprint(value))
		}
		productString = append(productString, "("+strings.Join(pairString, ", ")+")")
	}
	return "{" + strings.Join(productString, ", ") + "}"
}

func interfaceSingleOutput(firstSet []interface{}) string {
	var setString []string
	for _, value := range firstSet {
		setString = append(setString, fmt.Sprint(value))
	}
	return "{" + strings.Join(setString, ", ") + "}"
}
