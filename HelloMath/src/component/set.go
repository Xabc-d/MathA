package component

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Set struct {
	Bases
}

func (se *Set) OnMount(ctx app.Context) {
	se.Bases.QuestionSetter = se
}

func (se *Set) SetQuestion() {
	group1 := []func() Question{
		se.set1,
		se.set2,
		se.set3,
		se.set4,
		se.set6,
		se.set7,
	}
	se.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	se.theLevel = se.seed % 10
	se.question1 = group1[0]()
	se.question2 = group1[1]()
	se.question3 = group1[2]()
	se.question4 = group1[3]()
	se.question5 = group1[4]()
	fmt.Println("set set set")
}

// Cartesian product A x B
func (se *Set) set1() Question {
	ALength := se.rand.Intn(2) + 3
	ASlice := make([]int, 0, ALength)
	AExist := make(map[int]bool)
	for len(ASlice) < cap(ASlice) {
		a := se.rand.Intn(20) + 1
		if !AExist[a] {
			ASlice = append(ASlice, a)
			AExist[a] = true
		}
	}
	sort.Slice(ASlice, func(i, j int) bool {
		return ASlice[i] < ASlice[j]
	})
	BLength := se.rand.Intn(2) + 3
	BSlice := make([]int, 0, BLength)
	BExist := make(map[int]bool)
	for len(BSlice) < cap(BSlice) {
		b := se.rand.Intn(20) + 1
		if !BExist[b] {
			BSlice = append(BSlice, b)
			BExist[b] = true
		}
	}
	sort.Slice(BSlice, func(i, j int) bool {
		return BSlice[i] < BSlice[j]
	})
	interfaceSliceA := make([]interface{}, len(ASlice))
	interfaceSliceB := make([]interface{}, len(BSlice))
	for i, v := range ASlice {
		interfaceSliceA[i] = v
	}
	for i, v := range BSlice {
		interfaceSliceB[i] = v
	}
	interfaceSliceX := []interface{}{"x", "y", "z"}
	interfaceSliceD := []interface{}{"x", "y", "z"}
	Sets := [][]interface{}{interfaceSliceA, interfaceSliceB, interfaceSliceX, interfaceSliceD}
	twoSet := se.rand.Perm(len(Sets))
	firstSet := Sets[twoSet[0]]
	secondSet := Sets[twoSet[1]]
	var product [][]interface{}
	for _, a := range firstSet {
		for _, b := range secondSet {
			product = append(product, []interface{}{a, b})
		}
	}
	firstSetOutput := interfaceSingleOutput(firstSet)
	secondSetOutput := interfaceSingleOutput(secondSet)
	//randC := se.rand.Intn(6)
	statement := ""
	theAnswer := ""
	statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of Cartesian product A x B.", firstSetOutput, secondSetOutput)
	theAnswer = interfaceTwoOutput(product)
	//randC := se.rand.Intn(6)
	//switch randC {
	//case 0:
	//	statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of Cartesian product A x B.", firstSetOutput, secondSetOutput)
	//	theAnswer = interfaceTwoOutput(product)
	//case 1:
	//	statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of Cartesian product B x A.", firstSetOutput, secondSetOutput)
	//	theAnswer = interfaceTwoOutput(product)
	//case 2:
	//	statement = fmt.Sprintf("Given the sets A = %v A = %v. List all elements of Cartesian product A x A.", firstSetOutput, firstSetOutput)
	//	theAnswer = interfaceTwoOutput(product)
	//case 3:
	//	statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of Cartesian product A x B x B.", secondSetOutput, secondSetOutput)
	//	theAnswer = interfaceTwoOutput(product)
	//default:
	//	statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of Cartesian product A x B x A.", firstSetOutput, secondSetOutput)
	//	theAnswer = interfaceTwoOutput(product)
	//}

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    1,
	}
	return question
}

func (se *Set) set2() Question {

	C := se.rand.Intn(15)

	statement := fmt.Sprintf("If set A contains exactly %v distinct elements, How many elements does the power set of A contain?", C)
	theAnswer := strconv.Itoa(1 << uint(C))

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    1,
	}
	return question
}
func (se *Set) set3() Question {
	ALength := se.rand.Intn(4)
	ADou := se.rand.Intn(3)
	ASlice := make([]int, 0, ALength)
	for len(ASlice) < cap(ASlice) {
		if len(ASlice) == ADou {
			a := se.rand.Intn(20) + 1
			ASlice = append(ASlice, a)
			ASlice = append(ASlice, a)
		} else {
			a := se.rand.Intn(20) + 1
			ASlice = append(ASlice, a)
		}

	}
	ASliceDE := make(map[int]bool)
	for _, value := range ASlice {
		ASliceDE[value] = true
	}
	DE := len(ASliceDE)
	theAnswer := fmt.Sprintf("%v", DE)

	ASliceOutput := setOutput(ASlice)
	statement := fmt.Sprintf("What is the cardinality of this set A = %v?", ASliceOutput)

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    1,
	}
	return question
}
func (se *Set) set4() Question {
	ALength := se.rand.Intn(2) + 3
	ASlice := make([]int, 0, ALength)
	AExist := make(map[int]bool)
	for len(ASlice) < cap(ASlice) {
		a := se.rand.Intn(30)
		if !AExist[a] {
			ASlice = append(ASlice, a)
			AExist[a] = true
		}
	}
	sort.Slice(ASlice, func(i, j int) bool {
		return ASlice[i] < ASlice[j]
	})
	BLength := se.rand.Intn(2) + 3
	BSlice := make([]int, 0, BLength)
	BExist := make(map[int]bool)
	for len(BSlice) < cap(BSlice) {
		b := se.rand.Intn(30)
		if !BExist[b] {
			BSlice = append(BSlice, b)
			BExist[b] = true
		}
	}
	sort.Slice(BSlice, func(i, j int) bool {
		return BSlice[i] < BSlice[j]
	})

	C := se.rand.Intn(3)
	newSet := make([]int, 0)
	statement := ""
	theAnswer := ""
	if C == 1 {
		//union
		newSetMapDu := make(map[int]bool)
		for _, a := range ASlice {
			newSetMapDu[a] = true
		}
		for _, b := range BSlice {
			newSetMapDu[b] = true
		}
		newSet = make([]int, 0, len(newSetMapDu))
		for key := range newSetMapDu {
			newSet = append(newSet, key)
		}
		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of the union A U B.", setOutput(ASlice), setOutput(BSlice))

	} else if C == 2 {
		//intersection
		for _, a := range ASlice {
			for _, b := range BSlice {
				if a == b {
					newSet = append(newSet, a)
				}
			}
		}
		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of the intersection A ∩ B.", setOutput(ASlice), setOutput(BSlice))
	} else {
		//difference
		newSetMapDu := make(map[int]bool)
		for _, b := range BSlice {
			newSetMapDu[b] = true
		}
		for _, a := range ASlice {
			if !newSetMapDu[a] {
				newSet = append(newSet, a)
			}
		}

		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = %v B = %v. List all elements of the difference A \\ B.", setOutput(ASlice), setOutput(BSlice))
	}

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    2,
	}
	return question
}

// choice question
func (se *Set) randomSet() []int {
	//randomSet
	setLength := se.rand.Intn(3) + 3
	elements := make(map[int]bool)
	set := make([]int, 0, setLength)

	for len(set) < cap(set) {
		element := se.rand.Intn(20) + 1
		if !elements[element] {
			set = append(set, element)
			elements[element] = true
		}

	}
	return set
}

func (se *Set) set6() Question {
	ALength := se.rand.Intn(2) + 3
	ASlice := make([]int, 0, ALength)
	//AExist := make(map[int]bool)
	a := 0
	for len(ASlice) < cap(ASlice) {
		ASlice = append(ASlice, a)
		a++
		//if !AExist[a] {
		//	ASlice = append(ASlice, a)
		//	a++
		//	AExist[a] = true
		//}
	}
	//sort.Slice(ASlice, func(i, j int) bool {
	//	return ASlice[i] < ASlice[j]
	//})
	BLength := se.rand.Intn(2) + 3
	BSlice := make([]int, 0, BLength)
	BExist := make(map[int]bool)
	for len(BSlice) < cap(BSlice) {
		b := se.rand.Intn(30)
		if !BExist[b] {
			BSlice = append(BSlice, b)
			BExist[b] = true
		}
	}
	sort.Slice(BSlice, func(i, j int) bool {
		return BSlice[i] < BSlice[j]
	})

	C := se.rand.Intn(3)
	newSet := make([]int, 0)
	statement := ""
	theAnswer := ""
	//Arestriction :=
	if C == 1 {
		//union
		newSetMapDu := make(map[int]bool)
		for _, a := range ASlice {
			newSetMapDu[a] = true
		}
		for _, b := range BSlice {
			newSetMapDu[b] = true
		}
		newSet = make([]int, 0, len(newSetMapDu))
		for key := range newSetMapDu {
			newSet = append(newSet, key)
		}
		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = {x ∈ N | x<%v} B = %v. List all elements of the union A U B.", ALength, setOutput(BSlice))

	} else if C == 2 {
		//intersection
		for _, a := range ASlice {
			for _, b := range BSlice {
				if a == b {
					newSet = append(newSet, a)
				}
			}
		}
		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = {x ∈ N | x<%v} B = %v. List all elements of the intersection A ∩ B.", ALength, setOutput(BSlice))
	} else {
		//difference
		newSetMapDu := make(map[int]bool)
		for _, b := range BSlice {
			newSetMapDu[b] = true
		}
		for _, a := range ASlice {
			if !newSetMapDu[a] {
				newSet = append(newSet, a)
			}
		}

		sort.Slice(newSet, func(i, j int) bool {
			return newSet[i] < newSet[j]
		})
		theAnswer = setOutput(newSet)
		statement = fmt.Sprintf("Given the sets A = {x ∈ N | x<%v} B = %v. List all elements of the difference A \\ B.", ALength, setOutput(BSlice))
	}

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    2,
	}
	return question
}

func (se *Set) set7() Question {
	setLength := se.rand.Intn(3) + 3
	elements := make(map[int]bool)
	set := make([]int, 0, setLength)
	for len(set) < cap(set) {
		element := se.rand.Intn(20) + 1
		if !elements[element] {
			set = append(set, element)
			elements[element] = true
		}
	}
	statement := fmt.Sprintf("Given the set A = %v. How many elements does the power set of A contain?", setOutput(set))
	theAnswer := fmt.Sprintf("%v", 1<<uint(len(set)))

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    1,
	}
	return question
}

//choice question
