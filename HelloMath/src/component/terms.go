package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"strconv"
	"strings"
)

type Terms struct {
	Bases
}

func (te *Terms) OnMount(ctx app.Context) {
	te.Bases.QuestionSetter = te
}

func (te *Terms) SetQuestion() {
	group1 := []func() Question{
		te.Te1,
		te.te3,
		te.te4,
		te.te8,
	}
	group2 := []func() Question{
		te.te7,
		te.te6,
		te.te5,
	}
	te.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})
	te.rand.Shuffle(len(group2), func(i, j int) {
		group2[i], group2[j] = group2[j], group2[i]
	})
	te.theLevel = te.seed % 10
	if te.theLevel >= 0 && te.theLevel < 5 {
		te.question1 = group1[0]()
		te.question2 = group1[1]()
		te.question3 = group1[2]()
		te.question4 = group1[3]()
		te.question5 = group2[0]()
	} else {
		te.question1 = group1[0]()
		te.question2 = group1[1]()
		te.question3 = group2[0]()
		te.question4 = group2[1]()
		te.question5 = group2[2]()
	}
	fmt.Println("terms terms terms")
}

func (te *Terms) Te1() Question {
	//The set Zn contains natural number less than
	theN := te.rand.Intn(8) + 1
	randC := te.rand.Intn(12)
	statement := ""
	theAnswer := ""
	//Random Z
	randZ := te.rand.Intn(200) - 100
	//Random N
	randN1 := te.rand.Intn(200)
	randN2 := te.rand.Intn(200)
	//random float
	randF1 := te.rand.Float64() * 100

	switch randC {
	case 0:
		statement = fmt.Sprintf("Does the set ℤ%v include the element 0? (true or false)", randZ)
		//statement = fmt.Sprintf("Does the set \\(\\mathbb{Z}_{%v}\\) include the element 0?", theN)
		theAnswer = "true"

	case 1:
		statement = fmt.Sprintf("Does the set ℤ%v include the element %v?", randN1, randN2)
		if randN1 < randN2 {
			theAnswer = "false"
		} else {
			theAnswer = "true"
		}

	case 2:
		statement = fmt.Sprintf("Is the number π(pi) in the set Q? (true or false)")
		theAnswer = "false"
	case 3:
		statement = fmt.Sprintf("Is the number π(pi) in the set R? (true or false)")
		theAnswer = "true"
	case 4:
		statement = fmt.Sprintf("Is the number %v in the set Q? (true or false)", theN)
		theAnswer = "true"
	case 5:
		statement = fmt.Sprintf("Is the square root of 2 in the set Q? (true or false)")
		theAnswer = "false"
	case 6:
		statement = fmt.Sprintf("Is the square root of 3 in the set Q? (true or false)")
		theAnswer = "false"
	case 7:
		statement = fmt.Sprintf("Is the e (Euler's number) in the set Q? (true or false)")
		theAnswer = "false"
	case 8:
		statement = fmt.Sprintf("Is the repeating decimal 0.3333... in the set Q? (true or false)")
		theAnswer = "true"

	case 9:
		statement = fmt.Sprintf("Does the set ℤ%v include the element 0? (true or false)", randZ)
		theAnswer = "true"
	case 10:
		statement = fmt.Sprintf("Does the set N include the element %v? (true or false)", randF1)
		theAnswer = "false"
	case 11:
		statement = fmt.Sprintf("Does the set ℤ%v include the element 0? (true or false)", randF1)
		theAnswer = "false"

	default:
		statement = fmt.Sprintf("Does the set ℤn include negative elements? (true or false)")
		theAnswer = "false"
	}

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    1,
	}
	return question
}

// multi-choice question
func (te *Terms) te2() Question {
	statement := ""
	Crand := te.rand.Intn(3)

	theAnswer := ""
	switch Crand {
	case 0:
		statement = "Which syntax rule is demonstrated by the expression a+b ⊢ b+a?"
		theAnswer = "commutative law of addition"
	case 1:
		statement = "Which syntax rule is demonstrated by the expression a+(b+c) ⊢ (a+b)+c?"
		theAnswer = "associative law of addition"
	case 2:
		statement = "Which syntax rule is demonstrated by the expression a×1 ⊢ a?"
		theAnswer = "multiplicative identity law"
	default:
		statement = "Which syntax rule is demonstrated by the expression a×(b×c) ⊢ (a×b)×c?"
		theAnswer = "associative law of multiplication"
	}

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 3,
		level:        1,
		parseType:    1,
	}
	return question

}

func (te *Terms) te3() Question {
	a := te.rand.Intn(30) + 1
	b := te.rand.Intn(30) + 1
	c := te.rand.Intn(30) + 1
	d := te.rand.Intn(30) + 1

	statement := fmt.Sprintf("Given the syntax rule for the Associative Law of addition: (a+b)+c ⊢ a+(b+c), perform a derivation to simplify the expression ((%v+%v)+%v)+%v?", a, b, c, d)
	theAnswer := fmt.Sprintf("%v+(%v+(%v+%v))", a, b, c, d)
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    1,
	}
	return question
}

func (te *Terms) te4() Question {
	Crand := te.rand.Intn(2)
	statement := ""
	theAnswer := ""

	if Crand == 0 {
		statement = "If terms A, B are syntactically equivalent, A and B are the same in all contexts."
		theAnswer = "false"
	} else {
		statement = "If terms A, B are syntactically equivalent, which mean A can be derived from B and B can be derived from A."
		theAnswer = "true"
	}
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
		parseType:    1,
	}
	return question
}

func (te *Terms) te5() Question {
	randN1 := te.rand.Intn(10) + 1
	randN2 := te.rand.Intn(10) + 1
	C := randN2 + 1
	CFalse := randN2 + te.rand.Intn(4) + 2
	randTf := te.rand.Intn(3)
	statement := ""
	theAnswer := ""
	if randTf == 0 {
		statement = fmt.Sprintf("If given the syntax rules: x ⊢ x + %v, x + %v ⊢ x. Determine if the terms %v and terms %v are syntactically equivalent.", randN1, randN1, randN2, CFalse)
		theAnswer = "false"
	} else {
		statement = fmt.Sprintf("If given the syntax rules: x ⊢ x + %v, x + %v ⊢ x. Determine if the terms %v and terms %v are syntactically equivalent.", randN1, randN1, randN2, C)
		theAnswer = "true"
	}
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
		parseType:    1,
	}
	return question
}

func (te *Terms) te6() Question {
	C := te.rand.Intn(10)
	expression := ""
	theAnswer := ""
	switch C {
	case 0:
		expression = "a+(b+c)"
		theAnswer = "true"
	case 1:
		expression = "(a*(b+c))"
		theAnswer = "true"
	case 2:
		expression = "a+(b*c)"
		theAnswer = "false"
	case 3:
		expression = "(a+b)*(c+d)"
		theAnswer = "false"
	default:
		expression = "(a×b)×c"
		theAnswer = "true"
	}
	statement := fmt.Sprintf("Does the expression %v contain any unnecessary brackets?", expression)
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
		parseType:    1,
	}
	return question
}

func (te *Terms) te7() Question {
	operators := []string{"+", "-", "*", "/"}
	var rpnExpression string
	var stack []string
	randomNumber := te.rand.Intn(8) + 1
	randLength := te.rand.Intn(2) + 3
	rpnExpression += strconv.Itoa(randomNumber)
	switch randLength {
	case 3:
		for i := 0; i < 3; i++ {
			randomNumber = te.rand.Intn(8) + 1
			rpnExpression += " " + strconv.Itoa(randomNumber)
			randomOperator := operators[te.rand.Intn(len(operators))]
			rpnExpression += " " + randomOperator
		}
	case 4:
		for i := 0; i < 4; i++ {
			randomNumber = te.rand.Intn(8) + 1
			rpnExpression += " " + strconv.Itoa(randomNumber)
			randomOperator := operators[te.rand.Intn(len(operators))]
			rpnExpression += " " + randomOperator
		}
	default:
		for i := 0; i < 5; i++ {
			randomNumber = te.rand.Intn(8) + 1
			rpnExpression += " " + strconv.Itoa(randomNumber)
			randomOperator := operators[te.rand.Intn(len(operators))]
			rpnExpression += " " + randomOperator
		}

	}
	theRpn := rpnExpression
	tokens := strings.Split(theRpn, " ")
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			infix := fmt.Sprintf("(%s %s %s)", operand1, token, operand2)
			stack = append(stack, infix)
		default:
			stack = append(stack, token)
		}
	}
	theInfixWithBracelet := stack[0]
	theInfix := theInfixWithBracelet[1 : len(theInfixWithBracelet)-1]
	randC := te.rand.Intn(2)
	statement := ""
	theAnswer := ""
	if randC == 0 {
		statement = fmt.Sprintf("Given the RPN expression: %v, convert it to infix notation.", theRpn)
		theAnswer = theInfix
	} else {
		statement = fmt.Sprintf("Given the infix expression: %v, convert it to RPN notation.", theInfix)
		theAnswer = theRpn

	}
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

func (te *Terms) te8() Question {
	theAnswer := ""
	randC := te.rand.Intn(5)
	theSet := ""
	var theNumber int
	switch randC {
	case 0:
		theSet = "N(Natural Numbers)"
		theNumber = te.rand.Intn(100)
		theAnswer = "true"
	case 1:
		theSet = "N(Natural Numbers)"
		theNumber = -te.rand.Intn(100) - 1
		theAnswer = "false"
	case 2:
		theSet = "Z(integers)"
		theNumber = te.rand.Intn(200) - 100
		theAnswer = "true"
	case 3:
		theSet = "Q(Rational Numbers)"
		theNumber = te.rand.Intn(200) - 100
		theAnswer = "true"
	default:
		theSet = "R(Real Numbers)"
		theNumber = te.rand.Intn(200) - 100
		theAnswer = "true"

	}
	statement := fmt.Sprintf("Does the set %v include the number %v?", theSet, theNumber)
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
		parseType:    1,
	}
	return question
}
