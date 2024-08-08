package component

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Functions struct {
	Bases
}

func (fu *Functions) OnMount(ctx app.Context) {
	fu.Bases.QuestionSetter = fu
}

func (fu *Functions) SetQuestion() {
	group1 := []func() Question{
		fu.func1,
		fu.func2,
		fu.func3,
		fu.func4,
		fu.func5,
	}

	fu.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})

	fu.theLevel = fu.seed % 10
	fu.question1 = group1[0]()
	fu.question2 = group1[1]()
	fu.question3 = group1[2]()
	fu.question4 = group1[3]()
	fu.question5 = group1[4]()
	fmt.Println("functions functions functions")

}

func (fu *Functions) GenerateSeed(ctx app.Context, e app.Event) {
	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
	fmt.Println("input value:", inputValue)
	seed, err := strconv.ParseInt(inputValue, 10, 64)
	fu.seed = seed
	if err != nil {
		fmt.Println("No input or error converting string to integer:", err)
		seconds := time.Now().UnixNano() / 1000000
		fu.seed = (seconds % 9000) + 1000
		//seed = time.Now().UnixNano()/1000000 - 1713220000000
		//one day 86400 seconds
		fu.currentValue = strconv.FormatInt(fu.seed, 10) + " (No input or error input use random seed)"

	} else {
		fu.currentValue = strconv.FormatInt(fu.seed, 10)
	}

	fu.answerClicked = false
	source := rand.NewSource(fu.seed)
	fu.rand = rand.New(source)
	fu.SetQuestion()
	fu.seedClicked = true
	fu.Update()
}

type Element interface{}

func (fu *Functions) func1() Question {
	randFunc := fu.rand.Intn(5)
	theFunc := ""
	C := fu.rand.Intn(20) + 1
	D := fu.rand.Intn(20) + 1
	theAnswer := ""
	do := ""
	switch randFunc {
	case 0:
		theFunc = fmt.Sprintf("f(x) = %dx + %d", C, D)
		theAnswer = "bijective"
		do = "R→R"
	case 1:
		theFunc = fmt.Sprintf("f(x) = %dx^2 + %d", C, D)
		theAnswer = "none"
		do = "R→R"
	case 2:
		theFunc = fmt.Sprintf("f(x) = e^x + %d", C)
		theAnswer = "injective"
		do = "R→R"
	case 3:
		theFunc = fmt.Sprintf("f(x) = sin(x)")
		do = "[0,π]→[−1,1]"
		theAnswer = "bijective"
	default:
		theFunc = fmt.Sprintf("f(x) = sin(x)")
		do = "[0,π]→R"
		theAnswer = "injective"
	}

	statement := fmt.Sprintf("Determine whether the function f:%v, defined by %v, is injective, surjective, bijective or none.", do, theFunc)

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    1,
	}
	return question
}

func (fu *Functions) func2() Question {
	//inverse
	theAnswer := ""
	theFunc := ""
	randC := fu.rand.Intn(6) + 1
	randD := fu.rand.Intn(10) + 1
	randE := fu.rand.Intn(10) + 1
	switch randC {
	case 1:
		theFunc = fmt.Sprintf("(2x-3)/(x+1)")
		theAnswer = "(3x+3)/(2-x)"
	case 2:
		theFunc = fmt.Sprintf("x/(1-x)")
		theAnswer = "x/(1+x)"

	case 3:
		theFunc = fmt.Sprintf("x/(1+x)")
		theAnswer = "x/(1-x)"
	case 4:
		theFunc = fmt.Sprintf("(2x-3)/(x+1)")
		theAnswer = "(3x+3)/(2-x)"
	default:
		theFunc = fmt.Sprintf(" %vx + %v", randD, randE)
		theAnswer = fmt.Sprintf("(x-%v)/%v", randE, randD)
	}

	statement := fmt.Sprintf("Given the function f(x) = %v, find the inverse of the function.", theFunc)

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    1,
	}
	return question
}

func (fu *Functions) func3() Question {
	//image
	randC := fu.rand.Intn(3)
	theFunc := ""
	theAnswer := ""
	theX := fu.rand.Intn(20) + 1
	randD := fu.rand.Intn(20) + 1
	randE := fu.rand.Intn(20) + 1
	//y = kx + c
	LinearFunc := fmt.Sprintf("%vx + %v", randD, randE)
	LinearImage := randD*theX + randE
	//y = x^2 + b+ c
	QuadraticFunc := fmt.Sprintf("x^2 + %vx + %v", randD, randE)
	QuadraticImage := theX*theX + randD*theX + randE
	//y = x^3
	CubicFunc := fmt.Sprintf("x^3")
	CubicImage := theX * theX * theX

	switch randC {
	case 0:
		theFunc = LinearFunc
		theAnswer = fmt.Sprintf("%v", LinearImage)

	case 1:
		theFunc = QuadraticFunc
		theAnswer = fmt.Sprintf("%v", QuadraticImage)
	default:
		theFunc = CubicFunc
		theAnswer = fmt.Sprintf("%v", CubicImage)

	}

	statement := fmt.Sprintf("Given the function f(x) = %v, find the image of x= %v", theFunc, theX)
	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    1,
	}

}

func (fu *Functions) func4() Question {
	randC := fu.rand.Intn(3)
	theFunc := ""
	theAnswer := ""
	theX := fu.rand.Intn(20) + 1
	randD := fu.rand.Intn(20) + 1
	randE := fu.rand.Intn(20) + 1
	//y = kx + c
	LinearFunc := fmt.Sprintf("%vx + %v", randD, randE)
	LinearImage := randD*theX + randE
	X2 := -randD - theX
	//y = x^2 + b+ c
	QuadraticFunc := fmt.Sprintf("x^2 + %vx + %v", randD, randE)
	QuadraticImage := theX*theX + randD*theX + randE

	//y = x^3
	CubicFunc := fmt.Sprintf("x^3")
	CubicImage := theX * theX * theX
	var image int

	switch randC {
	case 0:
		theFunc = LinearFunc
		image = LinearImage
		theAnswer = fmt.Sprintf("%v", theX)

	case 1:
		theFunc = QuadraticFunc
		image = QuadraticImage
		if X2 == theX {
			theAnswer = fmt.Sprintf("%v", theX)
		} else {
			theAnswer = fmt.Sprintf("%v %v", theX, X2)
		}

		//theAnswer = fmt.Sprintf("%v", QuadraticImage)
	default:
		theFunc = CubicFunc
		image = CubicImage
		theAnswer = fmt.Sprintf("%v", theX)

	}

	statement := fmt.Sprintf("Given the function f(x) = %v, find the Preimage of f(x)= %v", theFunc, image)
	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    1,
	}

}

func (fu *Functions) func5() Question {
	//composition
	fun1 := ""
	fun2 := ""
	theAnswer := ""
	randC := fu.rand.Intn(3)
	randD := fu.rand.Intn(20) + 1
	randE := fu.rand.Intn(20) + 1
	randF := fu.rand.Intn(20) + 1
	randG := fu.rand.Intn(20) + 1

	LinearFunc1 := fmt.Sprintf("%vx + %v", randD, randE)
	LinearFunc2 := fmt.Sprintf("%vx + %v", randF, randG)
	CubicFunc := fmt.Sprintf("x^3")
	switch randC {
	case 0:
		fun1 = LinearFunc1
		fun2 = LinearFunc2
		theAnswer = fmt.Sprintf("%vx + %v", randF*randD, randD*randG+randE)
	case 1:
		fun1 = LinearFunc1
		fun2 = CubicFunc
		theAnswer = fmt.Sprintf("%vx^3 + %v", randD, randE)
	default:
		fun1 = fmt.Sprintf("1/(x+%v)", randD)
		fun2 = fmt.Sprintf("%vx", randE)
		theAnswer = fmt.Sprintf("1/(%vx+%v)", randE, randD)

	}

	statement := fmt.Sprintf("If f(x) = %v and g(x) = %v, find f(g(x))", fun1, fun2)
	return Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
		parseType:    1,
	}
}

func (fu *Functions) randomSet() []int {
	//randomSet
	setLength := fu.rand.Intn(3) + 3
	elements := make(map[int]bool)
	set := make([]int, 0, setLength)

	for len(set) < cap(set) {
		element := fu.rand.Intn(20) + 1
		if !elements[element] {
			set = append(set, element)
			elements[element] = true
		}

	}
	return set
}
