package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"sort"
)

type Relation struct {
	Bases
}

func (r *Relation) OnMount(ctx app.Context) {
	r.Bases.QuestionSetter = r
}

func (r *Relation) SetQuestion() {

	group1 := []func() Question{
		r.createRelation,
		r.inverse,
		r.inverseAlp,
		r.inverseFix,
		r.reflexiveSingle,
		r.antisymmetricSingle,
		r.transitiveSingle,
		r.symmetricSingle,
	}
	group2 := []func() Question{
		r.composition,
		r.partialOrder,
	}
	r.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})
	r.rand.Shuffle(len(group2), func(i, j int) {
		group2[i], group2[j] = group2[j], group2[i]
	})

	r.theLevel = r.seed % 10
	if r.theLevel >= 0 && r.theLevel < 5 {
		r.question1 = group1[0]()
		r.question2 = group1[1]()
		r.question3 = group1[2]()
		r.question4 = group1[3]()
		r.question5 = group1[4]()

	} else {
		r.question1 = group1[0]()
		r.question2 = group1[1]()
		r.question3 = group1[2]()
		r.question4 = group2[0]()
		r.question5 = group2[1]()

	}

}

// 1
func (r *Relation) createRelation() Question {
	R := make([][2]int, 0)

	//set example
	theMap := make(map[int]bool)
	partStatement := ""

	ALength := r.rand.Intn(2) + 3
	ASlice := make([]int, 0, ALength)
	AExist := make(map[int]bool)
	for len(ASlice) < cap(ASlice) {
		a := r.rand.Intn(20) + 1
		if !AExist[a] {
			ASlice = append(ASlice, a)
			AExist[a] = true
		}
	}
	sort.Slice(ASlice, func(i, j int) bool {
		return ASlice[i] < ASlice[j]
	})
	ASliceOutput := setOutput(ASlice)

	BLength := r.rand.Intn(2) + 3
	BSlice := make([]int, 0, BLength)
	BExist := make(map[int]bool)
	for len(BSlice) < cap(BSlice) {
		b := r.rand.Intn(20) + 1
		if !BExist[b] {
			BSlice = append(BSlice, b)
			BExist[b] = true
		}
	}
	for _, b := range BSlice {
		theMap[b] = true
	}
	sort.Slice(BSlice, func(i, j int) bool {
		return BSlice[i] < BSlice[j]
	})

	BSliceOutput := setOutput(BSlice)

	C := r.rand.Intn(6) + 1

	switch C {
	case 1:
		partStatement = "A is double of B"
		for _, a := range ASlice {
			doubleA := 2 * a
			if theMap[doubleA] {
				R = append(R, [2]int{a, doubleA})
			}
		}
	case 2:
		partStatement = "A is triple of B"
		for _, a := range ASlice {
			tripleA := 3 * a
			if theMap[tripleA] {
				R = append(R, [2]int{a, tripleA})
			}
		}
	case 3:
		partStatement = "A is less than B"
		for _, a := range ASlice {
			for _, b := range BSlice {
				if a < b {
					R = append(R, [2]int{a, b})
				}
			}

		}
	case 4:
		partStatement = "A is greater than B"
		for _, a := range ASlice {
			for _, b := range BSlice {
				if a > b {
					R = append(R, [2]int{a, b})
				}
			}
		}
	case 5:
		partStatement = "(a+b) mod 2 = 0"
		for _, a := range ASlice {
			for _, b := range BSlice {
				if (a+b)%2 == 0 {
					R = append(R, [2]int{a, b})
				}
			}

		}
	case 6:
		partStatement = "(a-b) mod 2 = 0"
		for _, a := range ASlice {
			for _, b := range BSlice {
				if (a-b)%2 == 0 {
					R = append(R, [2]int{a, b})
				}
			}
		}
	default:
		partStatement = "A is equal to B"
		for _, a := range ASlice {
			for _, b := range BSlice {
				if a == b {
					R = append(R, [2]int{a, b})
				}
			}
		}
	}
	var theAnswer string

	if len(R) == 0 {
		theAnswer = "{} empty relation"
	} else {
		theAnswer = relationOutput(R)

	}

	statement := fmt.Sprintf("Let A = %v and B = %v. Construct a relation R from A to B where R(a, b) exists if %s.List all the elements of the relation R.", ASliceOutput, BSliceOutput, partStatement)

	return Question{
		answer:       theAnswer,
		thisRelation: R,
		statement1:   statement,
		questionType: 1,
		parseType:    3,
	}
}

//func (r *Relation) Reflexive() Question {
//	//func Reflexive(set []string, relation map[theRelation]bool) bool {
//
//	//for _, element := range set {
//	//	if !relation[theRelation{element, element}] {
//	//		return false
//	//	}
//	//}
//	statement := fmt.Sprintf("")
//
//	return Question{
//		statement1: statement,
//		//solution:
//	}
//	//return true
//}

//func (r *Relation) Antisymmetric() Question {
//	//func Antisymmetric(relation map[theRelation]bool) bool {
//	//for rela := range relation {
//	//	if rela.relation1 != rela.relation2 && relation[theRelation{rela.relation2, rela.relation1}] {
//	//		return false
//	//	}
//	//}
//	statement := fmt.Sprintf("")
//
//	return Question{
//		statement1: statement,
//		//solution:
//	}
//
//	//return true
//}

//func (r *Relation) Transitive() Question {
//	//func Transitive(relation map[theRelation]bool) bool {
//	//for rela1 := range relation {
//	//	for rela2 := range relation {
//	//		if rela1.relation2 == rela2.relation1 {
//	//			if !relation[theRelation{rela1.relation1, rela2.relation2}] {
//	//				return false
//	//			}
//	//		}
//	//	}
//	//}
//	statement := fmt.Sprintf("")
//
//	return Question{
//		statement1: statement,
//		//solution:
//	}
//	//return true
//}

//func (r *Relation) TotallyOrdered() Question {
//	//func TotallyOrdered(set []string, relation map[theRelation]bool) bool {
//	//for _, a := range set {
//	//	for _, b := range set {
//	//		if a != b {
//	//			if !(relation[theRelation{a, b}] || relation[theRelation{b, a}]) {
//	//				return false
//	//			}
//	//		}
//	//	}
//	//}
//
//	statement := fmt.Sprintf("")
//
//	return Question{
//		statement1: statement,
//		//solution:
//	}
//	//return true
//}

// 2
func (r *Relation) inverse() Question {
	numberPairs := r.rand.Intn(3) + 3
	R := make([][2]int, 0, numberPairs)
	RExist := make(map[int]bool)
	for len(R) < numberPairs {
		// Generate random integers, adjust range as needed
		x := r.rand.Intn(100) - 50 // Random integers from -50 to 49
		y := r.rand.Intn(100) - 50 // Random integers from -50 to 49

		if !RExist[x] {
			R = append(R, [2]int{x, y})
			RExist[x] = true
		}
	}
	outPutR := relationOutput(R)
	statement := fmt.Sprintf("Find the inverse of a relation R represented by %v", outPutR)

	RInverse := make([][2]int, numberPairs)
	for i, pair := range R {
		RInverse[i] = [2]int{pair[1], pair[0]}
	}

	return Question{
		answer:       relationOutput(RInverse),
		thisRelation: R,
		statement1:   statement,
		questionType: 1,
		parseType:    3,
	}
}

// 3
func (r *Relation) inverseAlp() Question {
	//Find the inverse relation of R = {(a, x), (b, y), (c, z)}
	theElement := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	theRelation := make([][2]string, 0, 3)
	r.rand.Shuffle(len(theElement), func(i, j int) {
		theElement[i], theElement[j] = theElement[j], theElement[i]
	})
	afterElement := theElement[:6]

	for i := 0; i < len(afterElement); i += 2 {
		theRelation = append(theRelation, [2]string{afterElement[i], afterElement[i+1]})
	}
	RInverse := make([][2]string, len(theRelation))
	for i, pair := range theRelation {
		RInverse[i] = [2]string{pair[1], pair[0]}
	}
	outPutR := relationAlpOutput(theRelation)
	statement := fmt.Sprintf("Find the inverse of a relation R represented by %v", outPutR)

	return Question{
		statement1:   statement,
		answer:       relationAlpOutput(RInverse),
		questionType: 1,
		parseType:    1,
	}

}

// 4
func (r *Relation) inverseFix() Question {
	C := r.rand.Intn(3)
	var phrase string
	var answer string
	switch C {
	case 1:
		phrase = "doctor d treats patient p"
		answer = "patient p is treated by doctor d"
	case 2:
		phrase = "student a owns car y"
		answer = "car y is owned by student a"
	case 0:
		phrase = "employee e is managed by manager m"
		answer = "manager m manages employee e"
	default:
		phrase = "employee e is managed by manager m"
		answer = "manager m manages employee e"
	}

	statement := fmt.Sprintf("Suppose there is a relation '%s'. What would be the inverse of this relation?", phrase)
	return Question{
		statement1:   statement,
		answer:       answer,
		questionType: 1,
		parseType:    1,
	}
}

// 5
func (r *Relation) composition() Question {
	//Given two relations R = {(1, 2), (3, 4)} and S = {(2, 5), (4, 6)}, determine the composition R ◦ S.
	numberPairs := r.rand.Intn(2) + 3
	R := make([][2]int, numberPairs)
	S := make([][2]int, numberPairs)
	J := make([][2]int, numberPairs)

	for i := 0; i < numberPairs; i++ {
		// Generate random integers, adjust range as needed
		x := r.rand.Intn(30) - 15 // Random integers from -50 to 49
		y := r.rand.Intn(30) - 15
		z := r.rand.Intn(30) - 15

		R[i] = [2]int{x, y}
		S[i] = [2]int{y, z}
		J[i] = [2]int{x, z}
	}

	outPutR := relationOutput(R)
	outPutS := relationOutput(S)
	theSolution := relationOutput(J)
	statement := fmt.Sprintf("Given two relations R = %v and S = %v, determine the composition R ◦ S", outPutR, outPutS)

	//statement := fmt.Sprintf("Given two relations R = %v and S = %v, determine the composition R ◦ S", outPutR, outPutS)
	return Question{
		answer:       theSolution,
		statement1:   statement,
		questionType: 1,
		parseType:    3,
	}
}

// 6
func (r *Relation) reflexiveSingle() Question {
	theSet := make([]int, 0, 2)
	A := r.rand.Intn(30) - 30
	B := r.rand.Intn(30) - 30
	for A == B {
		B = r.rand.Intn(30) - 30
	}
	theSet = append(theSet, A, B)
	R := [][2]int{{A, A}, {B, B}}
	outPutR := relationOutput(R)
	theSetOutput := setOutput(theSet)
	statement := fmt.Sprintf("Suppose set A = %v. What is the reflexive relation of A?", theSetOutput)
	return Question{
		statement1:   statement,
		answer:       outPutR,
		questionType: 1,
		parseType:    3,
	}
}

// 7
func (r *Relation) antisymmetricSingle() Question {
	theSet := make([]int, 0, 2)
	A := r.rand.Intn(30) - 30
	B := r.rand.Intn(30) - 30
	for A == B {
		B = r.rand.Intn(30) - 30
	}
	theSet = append(theSet, A, B)
	R := make([][2]int, 0)
	R = append(R, [2]int{A, B}, [2]int{B, A})
	C := r.rand.Intn(2)
	R1 := [][2]int{{A, B}}
	R2 := [][2]int{{B, A}}

	if C == 0 {
		R = append(R, [2]int{A, A})
	} else if C == 1 {
		R = append(R, [2]int{B, B})
	} else {
		R = append(R, [2]int{A, A}, [2]int{B, B})
	}
	//check the answer
	//if contains(R,[2]int{A, B}) || contains(R,[2]int{B, A}) {
	//	right
	//} else {
	//	wrong
	//}

	outPutR := relationOutput(R)
	theAnswer := fmt.Sprintf("Remove %v or %v", relationOutput(R1), relationOutput(R2))
	theSetOutput := setOutput(theSet)
	statement := fmt.Sprintf(
		"Suppose set A = %v and relation R = %v. To ensure R is antisymmetric, which pair should be removed? Please respond like this: Remove {(-5, -12)} or {(-12, -5)}",
		theSetOutput,
		outPutR,
	)
	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    3,
	}
}

// 8
func (r *Relation) transitiveSingle() Question {
	A := r.rand.Intn(30) - 30
	B := r.rand.Intn(30) - 30
	C := r.rand.Intn(30) - 30
	for A == B {
		B = r.rand.Intn(30) - 30
	}
	for B == C || A == C {
		C = r.rand.Intn(30) - 30
	}
	theSet := []int{A, B, C}
	theSetOutput := setOutput(theSet)
	R := [][2]int{{A, B}, {B, C}, {A, C}}
	outPutR := relationOutput(R)
	//check the answer
	//if contains(R,[2]int{A, B}) || contains(R,[2]int{B, C}) || contains(R,[2]int{A, C}) {
	//	right
	//} else {
	//	wrong
	//}

	statement := fmt.Sprintf("Provide an example of a transitive relation on the set A %v ?", theSetOutput)
	return Question{
		statement1:   statement,
		answer:       outPutR,
		questionType: 1,
		parseType:    3,
	}
}

// 9
func (r *Relation) symmetricSingle() Question {
	A := r.rand.Intn(30) - 30
	B := r.rand.Intn(30) - 30
	for A == B {
		B = r.rand.Intn(30) - 30
	}
	theSet := []int{A, B}
	theSetOutput := setOutput(theSet)
	R := [][2]int{{A, B}, {B, A}}
	outPutR := relationOutput(R)
	//check the answer
	//if contains(R,[2]int{A, B}) || contains(R,[2]int{B, A}) {
	//	right
	//} else {
	//	wrong
	//}

	statement := fmt.Sprintf("Provide an example of a symmetric relation for the set A = %v ?", theSetOutput)
	return Question{
		statement1:   statement,
		answer:       outPutR,
		questionType: 1,
		parseType:    3,
	}
}

// partial order
// single choice
// Given a set A = {1, 2, 3, 4} and relation R defined as {(1, 1), (2, 2), (3, 3), (4, 4), (1, 2), (2, 3), (1, 3), (2, 4)}, verify if R is a partial order.
func (r *Relation) partialOrder() Question {
	ToF := r.rand.Intn(2)
	FalseRe := r.rand.Intn(3)

	SetLength := r.rand.Intn(4) + 4
	theSet := make([]int, SetLength)
	for i := range theSet {
		theSet[i] = r.rand.Intn(20) + 1
	}
	x1 := r.rand.Intn(SetLength)
	x2 := r.rand.Intn(SetLength)
	x3 := r.rand.Intn(SetLength)
	R := make([][2]int, 0)
	for _, value := range theSet {
		R = append(R, [2]int{value, value})
	}
	if ToF == 1 && FalseRe == 1 {
		R = append(R[:x1], R[x1+1:]...)
	}

	//symmetric
	R = append(R, [2]int{theSet[x1], theSet[x2]})
	R = append(R, [2]int{theSet[x2], theSet[x3]})

	if ToF == 1 {
		if FalseRe == 1 {
			R = append(R, [2]int{theSet[x1], theSet[x3]})
		} else if FalseRe == 2 {
			R = append(R, [2]int{theSet[x3], theSet[x2]})
			R = append(R, [2]int{theSet[x1], theSet[x3]})
		} else {

		}

	} else {
		R = append(R, [2]int{theSet[x1], theSet[x3]})
	}

	sort.Slice(R, func(i, j int) bool {
		return R[i][0] < R[j][0]
	})
	theRelation := relationOutput(R)
	if len(theRelation) >= 2 && theRelation[0] == '{' && theRelation[len(theRelation)-1] == '}' {
		theRelation = theRelation[1 : len(theRelation)-1]
	}

	theAnser := ""
	statement := fmt.Sprintf("Given a set A = %v and relation R defined as {%v}, verify if R is a partial order", setOutput(theSet), theRelation)
	//if  ToF == 1 not partial order
	if ToF == 1 {
		theAnser = "false"
	} else {
		theAnser = "true"

	}
	return Question{
		statement1:   statement,
		answer:       theAnser,
		questionType: 2,
		parseType:    3,
	}
}

//	func (r *Relation) partialOrderDesign() Question {
//		SetLength := r.rand.Intn(1) + 3
//		theSet := make([]int, SetLength)
//		for i := range theSet {
//			theSet[i] = r.rand.Intn(20) + 1
//		}
//		//if Antisymmetric
//		//reflexive(
//		//symmetric
//		statement := fmt.Sprintf("Given a set A = %v, construct a relation that satisfies the partial order", theSet)
//		return Question{
//			statement1: statement,
//			//answer:
//		}
//	}

//func (r *Relation) equRelations() Question {
//	SetLength := r.rand.Intn(4) + 4
//	theSet := make([]int, SetLength)
//	for i := range theSet {
//		theSet[i] = r.rand.Intn(20) + 1
//	}
//	//total order
//	statement := fmt.Sprintf("Given a set A = %v and relation R defined as {}, verify if R is a an equivalence relation.", theSet)
//
//	//statement := fmt.Sprintf("Given a set A = %v and relation R defined as {%v}, verify if R is a an equivalence relation.", theSet)
//	return Question{
//		statement1: statement,
//		//answer:
//	}
//}

//func (r *Relation) equRelations2() Question {
//	SetLength := r.rand.Intn(4) + 4
//	theSet := make([]int, SetLength)
//	for i := range theSet {
//		theSet[i] = r.rand.Intn(20) + 1
//	}
//	//total order
//	statement := fmt.Sprintf("Given a set A = %v and partial relation R defined as {}, extend S to satisfy the equivalence relation", theSet)
//	return Question{
//		statement1: statement,
//		//answer:
//	}
//}

//func (r *Relation) lexicographic() Question {
//
//}

//single-choice
