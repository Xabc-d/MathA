package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Logic struct {
	Bases
}

func (lo *Logic) SetQuestion() {
	group1 := []func() Question{
		lo.lo1,
		lo.lo2,
		lo.lo3,
		lo.lo4,
		lo.lo5,
	}

	lo.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	lo.theLevel = lo.seed % 10
	lo.question1 = group1[0]()
	lo.question2 = group1[1]()
	lo.question3 = group1[2]()
	lo.question4 = group1[3]()
	lo.question5 = group1[4]()
	fmt.Println("logic logic logic")
}

func (lo *Logic) Render() app.UI {
	return lo.Bases.Render()
}

func (lo *Logic) OnMount(ctx app.Context) {
	lo.Bases.QuestionSetter = lo
}

func (lo *Logic) lo1() Question {
	statement := ""
	answer := ""
	aS := ""
	bS := ""
	syb := ""
	//randC := lo.rand.Intn(7)
	randA := lo.rand.Intn(2)
	if randA == 0 {
		aS = "false"
	} else {
		aS = "true"
	}
	randB := lo.rand.Intn(2)
	if randB == 0 {
		bS = "false"
	} else {
		bS = "true"
	}
	randSyb := lo.rand.Intn(2)
	if randSyb == 0 {
		syb = "→"
		if randA-randB == 1 {
			answer = "false"
		} else {
			answer = "true"
		}
	} else {
		syb = "↔"
		if randA == randB {
			answer = "true"
		} else {
			answer = "false"
		}
	}
	statement = fmt.Sprintf("If a = %v, b = %v, what is a%vb ?", aS, bS, syb)

	question := Question{
		statement1:   statement,
		answer:       answer,
		questionType: 2,
		parseType:    1,
	}
	return question
}

func (lo *Logic) lo2() Question {
	statement := ""
	answer := ""
	var aS string
	var bS string
	var syb string

	randA := lo.rand.Intn(2)
	if randA == 0 {
		aS = "false"
	} else {
		aS = "true"
	}
	randB := lo.rand.Intn(2)
	if randB == 0 {
		bS = "false"
	} else {
		bS = "true"
	}

	randSyb := lo.rand.Intn(3)
	if randSyb == 0 {
		syb = "∧"
		if randA+randB == 2 {
			answer = "true"
		} else {
			answer = "false"
		}

	} else if randSyb == 1 {
		syb = "∨"
		if randA+randB == 0 {
			answer = "false"
		} else {
			answer = "true"
		}

	} else {
		syb = "⊕"
		if randA == randB {
			answer = "false"
		} else {
			answer = "true"
		}

	}
	statement = fmt.Sprintf("If a = %v, b = %v, what is a%vb ?", aS, bS, syb)

	return Question{
		statement1:   statement,
		answer:       answer,
		questionType: 2,
		parseType:    1,
	}
}

func trueTableCheck(a int, b int, c string) int {
	var result int
	switch c {
	case "→":
		if a == 1 && b == 1 {
			result = 1
		}
		if a == 1 && b == 0 {
			result = 0
		}
		if a == 0 && b == 1 {
			result = 1
		} else {
			result = 1
		}
	case "↔":
		if a == 1 && b == 1 {
			result = 1
		}
		if a == 1 && b == 0 {
			result = 0
		}
		if a == 0 && b == 1 {
			result = 0
		} else {
			result = 1
		}
	case "∧":
		if a == 1 && b == 1 {
			result = 1
		}
		if a == 1 && b == 0 {
			result = 0
		}
		if a == 0 && b == 1 {
			result = 0
		} else {
			result = 0
		}
	case "∨":
		if a == 1 && b == 1 {
			result = 1
		}
		if a == 1 && b == 0 {
			result = 1
		}
		if a == 0 && b == 1 {
			result = 1
		} else {
			result = 0
		}
	case "⊕":
		if a == 1 && b == 1 {
			result = 0
		}
		if a == 1 && b == 0 {
			result = 1
		}
		if a == 0 && b == 1 {
			result = 1
		} else {
			result = 0
		}
	default:
		result = 0
	}
	return result
}

func negation(a int) int {
	if a == 1 {
		return 0
	}
	return 1
}

func (lo *Logic) lo3() Question {
	//var propositions = []string{"p", "q", "r", "s"}
	var operators = []string{"→", "↔", "∧", "∨", "⊕"}
	// Generate two random logical expressions
	randProposition1 := "p"
	randProposition2 := "q"
	randProposition3 := "r"
	randProposition4 := "s"
	randProposition5 := "z"
	randProposition6 := "y"
	randOperator1 := operators[lo.rand.Intn(len(operators))]
	randOperator2 := operators[lo.rand.Intn(len(operators))]
	randOperator3 := operators[lo.rand.Intn(len(operators))]
	randOperator4 := operators[lo.rand.Intn(len(operators))]
	randOperator5 := operators[lo.rand.Intn(len(operators))]
	//randOperator6 := operators[lo.rand.Intn(len(operators))]

	a := lo.rand.Intn(2)
	var aTruth string
	if a == 0 {
		aTruth = "false"
	} else {
		aTruth = "true"
	}
	b := lo.rand.Intn(2)
	var bTruth string
	if b == 0 {
		bTruth = "false"
	} else {
		bTruth = "true"
	}
	c := lo.rand.Intn(2)
	var cTruth string
	if c == 0 {
		cTruth = "false"
	} else {
		cTruth = "true"
	}
	d := lo.rand.Intn(2)
	var dTruth string
	if d == 0 {
		dTruth = "false"
	} else {
		dTruth = "true"
	}
	e := lo.rand.Intn(2)
	var eTruth string
	if e == 0 {
		eTruth = "false"
	} else {
		eTruth = "true"
	}
	var fTruth string
	f := lo.rand.Intn(2)
	if f == 0 {
		fTruth = "false"
	} else {
		fTruth = "true"
	}
	expression1 := ""
	expression2 := ""
	var result1 int

	randC := lo.rand.Intn(7)
	switch randC {
	case 0:
		//(p→q)∧(r→s)
		expression1 = "(" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "(" + randProposition3 + randOperator3 + randProposition4 + ")"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth
		result1 = trueTableCheck(trueTableCheck(a, b, randOperator1), trueTableCheck(c, d, randOperator3), randOperator2)
	case 1:
		//((p↔q)∧(r↔s))∧(p∨q)
		expression1 = "((" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "(" + randProposition3 + randOperator3 + randProposition4 + "))" + randOperator4 + "(" + randProposition5 + randOperator5 + randProposition6 + ")"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth + ", " + randProposition5 + " = " + eTruth + ", " + randProposition6 + " = " + fTruth
		result1 = trueTableCheck(trueTableCheck(trueTableCheck(a, b, randOperator1), trueTableCheck(c, d, randOperator3), randOperator2), trueTableCheck(e, f, randOperator5), randOperator4)
	case 2:
		//(p∧q)∨((r∧s)∧(p∨q))
		expression1 = "(" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "((" + randProposition3 + randOperator3 + randProposition4 + ")" + randOperator4 + "(" + randProposition5 + randOperator5 + randProposition6 + "))"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth + ", " + randProposition5 + " = " + eTruth + ", " + randProposition6 + " = " + fTruth
		result1 = trueTableCheck(trueTableCheck(a, b, randOperator1), trueTableCheck(trueTableCheck(c, d, randOperator3), trueTableCheck(e, f, randOperator5), randOperator4), randOperator2)
	case 3:
		//((p∧q)∨(r∧s))∧(p∨q)
		expression1 = "¬(" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "(" + randProposition3 + randOperator3 + randProposition4 + ")"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth
		result1 = trueTableCheck(negation(trueTableCheck(a, b, randOperator1)), trueTableCheck(c, d, randOperator3), randOperator2)
	case 4:
		expression1 = "(¬(" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "(" + randProposition3 + randOperator3 + randProposition4 + "))" + randOperator4 + "(" + randProposition5 + randOperator5 + randProposition6 + ")"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth + ", " + randProposition5 + " = " + eTruth + ", " + randProposition6 + " = " + fTruth
		result1 = trueTableCheck(trueTableCheck(negation(trueTableCheck(a, b, randOperator1)), trueTableCheck(c, d, randOperator3), randOperator2), trueTableCheck(e, f, randOperator5), randOperator4)
	default:
		expression1 = "¬((" + randProposition1 + randOperator1 + randProposition2 + ")" + randOperator2 + "(" + randProposition3 + randOperator3 + randProposition4 + "))" + randOperator4 + "(" + randProposition5 + randOperator5 + randProposition6 + ")"
		expression2 = randProposition1 + " = " + aTruth + ", " + randProposition2 + " = " + bTruth + ", " + randProposition3 + " = " + cTruth + ", " + randProposition4 + " = " + dTruth + ", " + randProposition5 + " = " + eTruth + ", " + randProposition6 + " = " + fTruth
		result1 = trueTableCheck(negation(trueTableCheck(trueTableCheck(a, b, randOperator1), trueTableCheck(c, d, randOperator3), randOperator2)), trueTableCheck(e, f, randOperator5), randOperator4)

	}
	theAnswer := ""
	if result1 == 1 {
		theAnswer = "true"
	} else {
		theAnswer = "false"
	}

	statement := fmt.Sprintf("If %v, what is the value of %v", expression2, expression1)

	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
	}
}

func (lo *Logic) lo4() Question {
	//randC := lo.rand.Intn(5)
	//sentence := ""
	theAnswer := ""
	s1 := ""
	s2 := ""
	s3 := ""
	s4 := ""

	//group00 := []string{""}
	//group01 := []string{""}
	//group11 := []string{""}
	//group10 := []string{""}
	rands1 := lo.rand.Intn(2)
	if rands1 == 0 {
		s1 = "not rain"
	} else {
		s1 = "rain"
	}
	rands2 := lo.rand.Intn(2)
	if rands2 == 0 {
		s2 = "not be windy"
	} else {
		s2 = "be windy"
	}

	rands3 := lo.rand.Intn(2)
	if rands3 == 0 {
		s3 = "not rain"
	} else {
		s3 = "rain"
	}
	rands4 := lo.rand.Intn(2)
	if rands4 == 0 {
		s4 = "not be windy"
	} else {
		s4 = "be windy"
	}
	if (s1 == "rain" || s2 == "be windy") && (s3 == "rain" && s4 == "be windy") {
		theAnswer = "false"
	} else {
		theAnswer = "true"
	}

	statement := fmt.Sprintf("Are the statements, 'it will %v or it will %v' and 'it will %v and it will %v' logically equivalent? ", s1, s2, s3, s4)

	//switch randC {
	//case 0:
	//	sentence = "'it will not rain or it will not be windy' and 'it will not rain and it will not be windy'"
	//	theAnswer = "false"
	//
	//case 1:
	//	sentence = "'it will rain or it will be windy' and 'it will rain and it will be windy'"
	//	theAnswer = "false"
	//case 2:
	//	sentence = "'it will rain or it will be windy' and 'it will rain and it will be windy'"
	//	theAnswer = "false"
	//case 3:
	//	sentence = "'it will not rain or it will be windy' and 'it will not rain and it will be windy'"
	//	theAnswer = "false"
	//default:
	//	sentence = "'it will rain or it will be windy' and 'it will be windy or it will rain'"
	//	theAnswer = "true"
	//
	//}

	//statement := fmt.Sprintf("\\(\\frac{a}{b}\\) Are the statements, %v logically equivalent?", sentence)
	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
		parseType:    1,
	}
}

func (lo *Logic) lo5() Question {
	randC := lo.rand.Intn(9)
	sentence := ""
	theAnswer := ""
	switch randC {
	case 0:
		sentence = "P ∧ P to true"
		theAnswer = "false"
	case 1:
		sentence = "P ∨ P to true"
		theAnswer = "false"
	case 2:
		sentence = "P∧(Q∨R) to P∧Q∨P∧R"
		theAnswer = "true"
	case 3:
		sentence = "¬(¬(P∧Q)) to P∨Q"
		theAnswer = "false"
	case 4:
		sentence = "(P∨Q)∨R to P∨Q∨R"
		theAnswer = "true"
	case 5:
		sentence = "P→Q to ¬P∨Q"
		theAnswer = "true"
	case 6:
		sentence = "P∨(P∧Q) to P"
		theAnswer = "true"
	case 7:
		sentence = "(P∧Q)∨(P∧¬Q) to P"
		theAnswer = "true"
	default:
		sentence = "P∧P to P"
		theAnswer = "true"
	}
	//var operators = []string{"→", "↔", "∧", "∨", "⊕" ¬}
	statement := fmt.Sprintf("Determine if the simplification of the following expression is correct: %v", sentence)

	//statement := fmt.Sprintf("\\(\\frac{a}{b}\\) Determine if the simplification of the following expression is correct: %v", sentence)
	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
		parseType:    1,
	}
}
