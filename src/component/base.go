package component

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type QuestionSetter interface {
	SetQuestion()
}

type Bases struct {
	app.Compo
	rand       *rand.Rand
	inputValue string
	seed       int64
	Q1input    string
	Q2input    string
	Q3input    string
	Q4input    string
	Q5input    string
	Q6input    string
	Q1option   string
	Q2option   string
	Q3option   string
	Q4option   string
	Q5option   string

	question1      Question
	question2      Question
	question3      Question
	question4      Question
	question5      Question
	question6      Question
	theLevel       int64
	seedClicked    bool
	answerClicked  bool
	currentValue   string
	QuestionSetter QuestionSetter
}

func (b *Bases) Render() app.UI {
	return app.Div().Class("container-fluid").Style("background-color", "#F0F8FF").Style("min-height", "100vh").Body(
		app.Div().Class("row bg-primary").Style("min-height", "50px").Body(),
		app.Div().Class("row").Style("background-color", "#add8e6").Body(
			app.H1().Style("width", "90%").Style("margin", "0 auto").Body(
				app.Text("HelloMath"),
			),
		),
		&MathJaxConfig{},
		//app.Text("\\(\\frac{a}{b}\\)"),

		app.Div().Class("row").Body(
			app.Div().Class("col-2").Body(&Sidebar{}),
			app.Div().Class("col-10").Body(
				app.Div().Style("height", "8vh").Body(
					app.H6().
						Style("margin-top", "4vh").
						//Style("text-align", "center").
						Text("Instructions HelloMath!"),
				),
				app.Div().Body(
					app.Input().
						ID("seedInput").
						Type("text").
						Value(b.inputValue).
						OnChange(b.OnInputChange),
					app.Button().
						Text("Generate Seed").
						OnClick(b.GenerateSeed),
					app.Button().
						Text("up").
						OnClick(b.upLevel),

					app.Button().
						Text("down").
						OnClick(b.downLevel),
				),

				app.H6().Text("Seed Value: "+b.currentValue),
				app.Div().Style("height", "3vh").Body(),
				app.If(b.seedClicked,
					//question 1
					///////////
					app.If(b.question1.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(b.question1.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q1input").
									Type("text").
									Value(b.Q1input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(b.answerClicked,
								app.If(b.question1.parseType == 1,
									app.If(parse1(b.Q1input) == parse1(b.question1.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
							),
							app.If(b.question1.parseType == 2,
								app.If(parse2(b.Q1input) == parse2(b.question1.answer),
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
							app.If(b.question1.parseType == 3,
								app.If(parse3(b.Q1input) == parse3(b.question1.answer),
									app.Div().Class("col").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.If(b.answerClicked,
							app.P().Text("Answer: "+b.question1.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(b.question1.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q1 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(b.question1.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Div().Class("form-check").Style("margin-left", "50px").Body(
									app.Input().
										Type("radio").
										Name("options1").
										Class("form-check-input").
										Value("Option A").
										ID("Q1option-a"),
									app.Label().Class("form-check-label").Text("True"),
								),
							),
							app.If(b.answerClicked,
								app.If(b.Q1option == b.question1.answer,
									app.Div().Class("col-auto").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.Div().Class("form-check").Style("margin-left", "50px").Body(
							app.Input().
								Type("radio").
								Name("options1").
								Class("form-check-input").
								Value("Option B").
								ID("Q1option-b"),
							app.Label().Class("form-check-label").For("option-b").Text("False"),
						),
						app.If(b.answerClicked,
							app.P().Text("Answer: "+b.question1.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					////question 2
					app.If(b.question2.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(b.question2.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q2input").
									Type("text").
									Value(b.Q2input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(b.answerClicked,
								app.If(b.question2.parseType == 1,
									app.If(parse1(b.Q2input) == parse1(b.question2.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
								app.If(b.question2.parseType == 2,
									app.If(parse2(b.Q2input) == parse2(b.question2.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
								app.If(b.question2.parseType == 3,
									app.If(parse3(b.Q2input) == parse3(b.question2.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
							),
						),
						app.If(b.answerClicked,
							app.P().Text("Answer: "+b.question2.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(b.question2.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q2 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(b.question2.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Div().Class("form-check").Style("margin-left", "50px").Body(
									app.Input().
										Type("radio").
										Name("options2").
										Class("form-check-input").
										Value("Option A").
										ID("Q2option-a"),
									app.Label().Class("form-check-label").Text("True"),
								),
							),
							app.If(b.answerClicked,
								app.If(b.Q2option == b.question2.answer,
									app.Div().Class("col-auto").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.Div().Class("form-check").Style("margin-left", "50px").Body(
							app.Input().
								Type("radio").
								Name("options2").
								Class("form-check-input").
								Value("Option B").
								ID("Q2option-b"),
							app.Label().Class("form-check-label").For("option-b").Text("False"),
						),
						app.If(b.answerClicked,
							app.P().Text("Answer: "+b.question2.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////3
					app.If(b.question3.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(b.question3.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q3input").
									Type("text").
									Value(b.Q3input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(b.answerClicked,
								app.If(b.question3.parseType == 1,
									app.If(parse1(b.Q3input) == parse1(b.question3.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
								app.If(b.question3.parseType == 2,
									app.If(parse2(b.Q3input) == parse2(b.question3.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
								app.If(b.question3.parseType == 3,
									app.If(parse3(b.Q3input) == parse3(b.question3.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
							),
							//app.If(b.answerClicked,
							//	app.If(b.Q3input == b.question3.answer,
							//		//app.If(se.question1.answer == se.Q1input,
							//		app.Div().Class("col").Body(
							//			app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
							//		),
							//	),
							//),
						),
						app.If(b.answerClicked,
							app.P().Text("Answer: "+b.question3.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(b.question3.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q3 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(b.question3.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Div().Class("form-check").Style("margin-left", "50px").Body(
									app.Input().
										Type("radio").
										Name("options3").
										Class("form-check-input").
										Value("Option A").
										ID("Q3option-a"),
									app.Label().Class("form-check-label").Text("True"),
								),
							),
							app.If(b.answerClicked,
								app.If(b.Q3option == b.question3.answer,
									app.Div().Class("col-auto").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.Div().Class("form-check").Style("margin-left", "50px").Body(
							app.Input().
								Type("radio").
								Name("options3").
								Class("form-check-input").
								Value("Option B").
								ID("Q3option-b"),
							app.Label().Class("form-check-label").For("option-b").Text("False"),
						),
						app.If(b.answerClicked,
							app.P().Text("Answer: "+b.question3.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					//////4
					app.If(b.question4.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(b.question4.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q4input").
									Type("text").
									Value(b.Q4input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(b.answerClicked,
								app.If(b.question4.parseType == 1,
									app.If(parse1(b.Q4input) == parse1(b.question4.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
								app.If(b.question4.parseType == 2,
									app.If(parse2(b.Q4input) == parse2(b.question4.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
								app.If(b.question4.parseType == 3,
									app.If(parse3(b.Q4input) == parse3(b.question4.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
							),
							//app.If(b.answerClicked,
							//	app.If(b.Q4input == b.question4.answer,
							//		app.Div().Class("col").Body(
							//			app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
							//		),
							//	),
							//),
						),
						app.If(b.answerClicked,
							app.P().Text("Answer: "+b.question4.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(b.question4.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q4 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(b.question4.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Div().Class("form-check").Style("margin-left", "50px").Body(
									app.Input().
										Type("radio").
										Name("options4").
										Class("form-check-input").
										Value("Option A").
										ID("Q4option-a"),
									app.Label().Class("form-check-label").Text("True"),
								),
							),
							app.If(b.answerClicked,
								app.If(b.Q4option == b.question4.answer,
									app.Div().Class("col-auto").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.Div().Class("form-check").Style("margin-left", "50px").Body(
							app.Input().
								Type("radio").
								Name("options4").
								Class("form-check-input").
								Value("Option B").
								ID("Q4option-b"),
							app.Label().Class("form-check-label").For("option-b").Text("False"),
						),
						app.If(b.answerClicked,
							app.P().Text("Answer: "+b.question4.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),
					///////5
					app.If(b.question5.questionType == 1,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(b.question5.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Input().
									ID("Q5input").
									Type("text").
									Value(b.Q5input).
									Style("width", "400px").
									Style("height", "30px").
									Style("margin-left", "50px").
									Class("form-control mb-3").
									Placeholder("").
									AutoFocus(true),
							),
							app.If(b.answerClicked,
								app.If(b.question5.parseType == 1,
									app.If(parse1(b.Q5input) == parse1(b.question5.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
								app.If(b.question5.parseType == 2,
									app.If(parse2(b.Q5input) == parse2(b.question5.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
								app.If(b.question5.parseType == 3,
									app.If(parse3(b.Q5input) == parse3(b.question5.answer),
										app.Div().Class("col").Body(
											app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
										),
									),
								),
							),
							//app.If(b.answerClicked,
							//	app.If(b.Q5input == b.question5.answer,
							//		//app.If(se.question1.answer == se.Q1input,
							//		app.Div().Class("col").Body(
							//			app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
							//		),
							//	),
							//),
						),
						app.If(b.answerClicked,
							app.P().Text("Answer: "+b.question5.answer).Style("font-size", "18px").Style("margin-left", "50px"),
						),
					),
					app.If(b.question5.questionType == 2,
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.H6().Text("Q5 :"),
							),
							app.Div().Class("col").Body(
								app.P().Text(b.question5.statement1),
							),
						),
						app.Div().Class("row").Body(
							app.Div().Class("col-auto").Body(
								app.Div().Class("form-check").Style("margin-left", "50px").Body(
									app.Input().
										Type("radio").
										Name("options5").
										Class("form-check-input").
										Value("Option A").
										ID("Q5option-a"),
									app.Label().Class("form-check-label").Text("True"),
								),
							),
							app.If(b.answerClicked,
								app.If(b.Q5option == b.question5.answer,
									app.Div().Class("col-auto").Body(
										app.Img().Src("/web/greencheck.svg").Class("img-fluid").Alt("check").Style("width", "30px").Style("height", "30px"),
									),
								),
							),
						),
						app.Div().Class("form-check").Style("margin-left", "50px").Body(
							app.Input().
								Type("radio").
								Name("options5").
								Class("form-check-input").
								Value("Option B").
								ID("Q5option-b"),
							app.Label().Class("form-check-label").For("option-b").Text("False"),
						),
						app.If(b.answerClicked,
							app.P().Text("Answer: "+b.question5.answer).Style("font-size", "17px").Style("margin-left", "50px"),
						),
					),

					app.Button().
						Text("Finish").
						Style("margin-left", "50px").
						OnClick(b.finishClicked),
				),
			),
		),
	)
}

func (b *Bases) OnInputChange(ctx app.Context, e app.Event) {
	b.inputValue = ctx.JSSrc().Get("value").String()
	b.Update()
}

func (b *Bases) SetQuestion() {
	group1 := []func() Question{
		b.Te1,
		b.te3,
		b.te4,
		b.te8,
	}
	group2 := []func() Question{
		b.te7,
		b.te6,
		b.te5,
	}
	b.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})
	b.rand.Shuffle(len(group2), func(i, j int) {
		group2[i], group2[j] = group2[j], group2[i]
	})
	b.theLevel = b.seed % 10
	if b.theLevel >= 0 && b.theLevel < 5 {
		b.question1 = group1[0]()
		b.question2 = group1[1]()
		b.question3 = group1[2]()
		b.question4 = group1[3]()
		b.question5 = group2[0]()
	} else {
		b.question1 = group1[0]()
		b.question2 = group1[1]()
		b.question3 = group2[0]()
		b.question4 = group2[1]()
		b.question5 = group2[2]()
	}
	fmt.Println("bbbbbb")
}

func (b *Bases) GenerateSeed(ctx app.Context, e app.Event) {
	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
	fmt.Println("input value:", inputValue)
	seed, err := strconv.ParseInt(inputValue, 10, 64)
	b.seed = seed
	if err != nil {
		fmt.Println("No input or error converting string to integer:", err)
		seconds := time.Now().UnixNano() / 1000000
		b.seed = (seconds % 9000) + 1000
		//seed = time.Now().UnixNano()/1000000 - 1713220000000
		//one day 86400 seconds
		b.currentValue = strconv.FormatInt(b.seed, 10) + " (No input or error input use random seed)"

	} else {
		b.currentValue = strconv.FormatInt(b.seed, 10)
	}

	b.answerClicked = false
	source := rand.NewSource(b.seed)
	b.rand = rand.New(source)
	fmt.Printf("base: %v\n", b.QuestionSetter != nil)
	if b.QuestionSetter != nil {
		b.QuestionSetter.SetQuestion()
	} else {
		fmt.Println("set base")
		b.SetQuestion()
	}
	b.seedClicked = true
	b.renderMathJax(ctx)
	b.clearMathJax(ctx)
	b.Update()

}

func (b *Bases) upLevel(ctx app.Context, e app.Event) {
	if b.seed%10 == 9 {
		b.seed = b.seed + 10
	} else {
		b.seed++
	}

	b.currentValue = strconv.FormatInt(b.seed, 10)
	b.inputValue = strconv.FormatInt(b.seed, 10)

	b.answerClicked = false
	source := rand.NewSource(b.seed)
	b.rand = rand.New(source)

	if b.QuestionSetter != nil {
		b.QuestionSetter.SetQuestion()
	} else {
		b.SetQuestion()
	}

	b.seedClicked = true
	b.Update()
}
func (b *Bases) downLevel(ctx app.Context, e app.Event) {
	if b.seed%10 == 0 {
		b.seed = b.seed - 10
	} else {
		b.seed--
	}

	b.currentValue = strconv.FormatInt(b.seed, 10)
	b.inputValue = strconv.FormatInt(b.seed, 10)

	b.answerClicked = false
	source := rand.NewSource(b.seed)
	b.rand = rand.New(source)

	if b.QuestionSetter != nil {
		b.QuestionSetter.SetQuestion()
	} else {
		b.SetQuestion()
	}

	b.seedClicked = true
	b.Update()
}

func (b *Bases) finishClicked(ctx app.Context, e app.Event) {
	b.answerClicked = true
	//q1
	if b.question1.questionType == 1 {
		b.Q1input = app.Window().GetElementByID("Q1input").Get("value").String()
	}
	if b.question1.questionType == 2 {
		if app.Window().GetElementByID("Q1option-a").Get("checked").Bool() {
			b.Q1option = "true"
		} else if app.Window().GetElementByID("Q1option-b").Get("checked").Bool() {
			b.Q1option = "false"
		}
	}
	//q2
	if b.question2.questionType == 1 {
		b.Q2input = app.Window().GetElementByID("Q2input").Get("value").String()
	}
	if b.question2.questionType == 2 {
		if app.Window().GetElementByID("Q2option-a").Get("checked").Bool() {
			b.Q2option = "true"
		} else if app.Window().GetElementByID("Q2option-b").Get("checked").Bool() {
			b.Q2option = "false"
		}
	}
	//q3
	if b.question3.questionType == 1 {
		b.Q3input = app.Window().GetElementByID("Q3input").Get("value").String()
	}
	if b.question3.questionType == 2 {
		if app.Window().GetElementByID("Q3option-a").Get("checked").Bool() {
			b.Q3option = "true"
		} else if app.Window().GetElementByID("Q3option-b").Get("checked").Bool() {
			b.Q3option = "false"
		}
	}
	//q4
	if b.question4.questionType == 1 {
		b.Q4input = app.Window().GetElementByID("Q4input").Get("value").String()
	}
	if b.question4.questionType == 2 {
		if app.Window().GetElementByID("Q4option-a").Get("checked").Bool() {
			b.Q4option = "true"
		} else if app.Window().GetElementByID("Q4option-b").Get("checked").Bool() {
			b.Q4option = "false"
		}
	}
	//q5
	if b.question5.questionType == 1 {
		b.Q5input = app.Window().GetElementByID("Q5input").Get("value").String()
	}
	if b.question5.questionType == 2 {
		if app.Window().GetElementByID("Q5option-a").Get("checked").Bool() {
			b.Q5option = "true"
		} else if app.Window().GetElementByID("Q5option-b").Get("checked").Bool() {
			b.Q5option = "false"
		}
	}
	b.Update()
}

func (b *Bases) Te1() Question {
	//The set Zn contains natural number less than
	theN := b.rand.Intn(8) + 1
	randC := b.rand.Intn(9)
	statement := ""
	theAnswer := ""

	//Random Z
	randZ := b.rand.Intn(200) - 100
	//Random N
	randN1 := b.rand.Intn(200)
	randN2 := b.rand.Intn(200)
	//random float
	randF1 := b.rand.Float64() * 100

	switch randC {
	case 0:
		statement = fmt.Sprintf("Does the set ℤ%v include the element 0?", randZ)
		theAnswer = "true"
	case 1:
		if randN1 < randN2 {
			theAnswer = "false"
		} else {
			theAnswer = "true"
		}
		statement = fmt.Sprintf("Does the set ℤ%v include the element %v?", randN1, randN2)

	case 9:
		statement = fmt.Sprintf("Does the set ℤ%v include the element 0?", randZ)
		theAnswer = "true"
	case 10:
		statement = fmt.Sprintf("Does the set N include the element %v?", randF1)
		theAnswer = "false"
	case 11:
		statement = fmt.Sprintf("Does the set ℤ%v include the element 0?", randF1)
		theAnswer = "false"

	//fix
	case 2:
		statement = fmt.Sprintf("Is the number π(pi) in the set Q?")
		theAnswer = "false"
	case 3:
		statement = fmt.Sprintf("Is the number π(pi) in the set R?")
		theAnswer = "true"
	case 4:
		statement = fmt.Sprintf("Is the number %v in the set Q?", theN)
		theAnswer = "true"
	case 5:
		statement = fmt.Sprintf("Is the square root of 2 in the set Q??")
		theAnswer = "false"
	case 6:
		statement = fmt.Sprintf("Is the square root of 3 in the set Q??")
		theAnswer = "false"
	case 7:
		statement = fmt.Sprintf("Is the e (Euler's number) in the set Q?")
		theAnswer = "false"
	case 8:
		statement = fmt.Sprintf("Is the repeating decimal 0.3333... in the set Q?")
		theAnswer = "true"
	default:
		statement = fmt.Sprintf("Does the set ℤn include negative elements?")
		theAnswer = "false"
	}

	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

// multi-choice question
func (b *Bases) te2() Question {
	statement := ""
	Crand := b.rand.Intn(3)

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
	}
	return question
}

func (ba *Bases) te3() Question {
	a := ba.rand.Intn(30) + 1
	b := ba.rand.Intn(30) + 1
	c := ba.rand.Intn(30) + 1
	d := ba.rand.Intn(30) + 1

	statement := fmt.Sprintf("Given the syntax rule for the Associative Law of addition: (a+b)+c ⊢ a+(b+c), perform a derivation to simplify the expression ((%v+%v)+%v)+%v?", a, b, c, d)
	theAnswer := fmt.Sprintf("%v+(%v+(%v+%v))", a, b, c, d)
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}

func (b *Bases) te4() Question {
	Crand := b.rand.Intn(2)
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
	}
	return question
}

func (b *Bases) te5() Question {
	randN1 := b.rand.Intn(10) + 1
	randN2 := b.rand.Intn(10) + 1
	C := randN2 + 1
	CFalse := randN2 + b.rand.Intn(4) + 2
	randTf := b.rand.Intn(3)
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
	}
	return question
}

func (b *Bases) te6() Question {
	C := b.rand.Intn(10)
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
	}
	return question

}

func (b *Bases) te7() Question {
	operators := []string{"+", "-", "*", "/"}
	var rpnExpression string
	var stack []string
	randomNumber := b.rand.Intn(8) + 1
	randLength := b.rand.Intn(2) + 3
	rpnExpression += strconv.Itoa(randomNumber)
	switch randLength {
	case 3:
		for i := 0; i < 3; i++ {
			randomNumber = b.rand.Intn(8) + 1
			rpnExpression += " " + strconv.Itoa(randomNumber)
			randomOperator := operators[b.rand.Intn(len(operators))]
			rpnExpression += " " + randomOperator
		}
	case 4:
		for i := 0; i < 4; i++ {
			randomNumber = b.rand.Intn(8) + 1
			rpnExpression += " " + strconv.Itoa(randomNumber)
			randomOperator := operators[b.rand.Intn(len(operators))]
			rpnExpression += " " + randomOperator
		}
	default:
		for i := 0; i < 5; i++ {
			randomNumber = b.rand.Intn(8) + 1
			rpnExpression += " " + strconv.Itoa(randomNumber)
			randomOperator := operators[b.rand.Intn(len(operators))]
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
	randC := b.rand.Intn(2)
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

func (b *Bases) te8() Question {
	theAnswer := ""
	randC := b.rand.Intn(5)
	theSet := ""
	var theNumber int
	switch randC {
	case 0:
		theSet = "N(Natural Numbers)"
		theNumber = b.rand.Intn(100)
		theAnswer = "true"
	case 1:
		theSet = "N(Natural Numbers)"
		theNumber = -b.rand.Intn(100) - 1
		theAnswer = "false"
	case 2:
		theSet = "Z(integers)"
		theNumber = b.rand.Intn(200) - 100
		theAnswer = "true"
	case 3:
		theSet = "Q(Rational Numbers)"
		theNumber = b.rand.Intn(200) - 100
		theAnswer = "true"
	default:
		theSet = "R(Real Numbers)"
		theNumber = b.rand.Intn(200) - 100
		theAnswer = "true"

	}
	statement := fmt.Sprintf("Does the set %v include the number %v?", theSet, theNumber)
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 2,
	}
	return question
}

func (b *Bases) renderMathJax(ctx app.Context) {
	ctx.Async(func() {
		app.Window().Get("MathJax").Call("typeset")
	})
}
func (b *Bases) clearMathJax(ctx app.Context) {
	ctx.Async(func() {
		app.Window().Get("MathJax").Call("typesetClear")
	})
}

// pa
func parse1(preInput string) string {
	afInput := strings.ReplaceAll(preInput, " ", "")
	return afInput
}

// {1, 2, 3} questiontype 2
func parse2(preInput string) string {
	removeSpace := strings.ReplaceAll(preInput, " ", "")
	removeBr := strings.Trim(removeSpace, "{}")
	theN := strings.Split(removeBr, ",")
	var numSet []int
	for _, strNum := range theN {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return ""
		}
		numSet = append(numSet, num)
	}
	sort.Ints(numSet)
	var setString []string
	for _, value := range numSet {
		setString = append(setString, strconv.Itoa(value))
	}
	afInput := "{" + strings.Join(setString, ", ") + "}"
	return afInput
}

func parse3(preInput string) string {
	var relation [][2]int
	removeSpace := strings.ReplaceAll(preInput, " ", "")
	removeBr := strings.Trim(removeSpace, "{}")
	thePair := strings.Split(removeBr, "),(")
	for _, pair := range thePair {
		pair = strings.Trim(pair, "()")
		pairStr := strings.Split(pair, ",")
		if len(pairStr) < 2 {
			continue
		}
		x, err1 := strconv.Atoi(pairStr[0])
		y, err2 := strconv.Atoi(pairStr[1])
		if err1 != nil || err2 != nil {
			continue
		}
		relation = append(relation, [2]int{x, y})
	}
	sort.Slice(relation, func(i, j int) bool {
		if relation[i][0] == relation[j][0] {
			return relation[i][1] < relation[j][1]
		}
		return relation[i][0] < relation[j][0]
	})
	afInput := relationOutput(relation)
	return afInput
}

//tiankong 1
//xuanze 2
