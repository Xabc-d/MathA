package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"math/rand"
	"strconv"
	"time"
)

type template struct {
	app.Compo
	rand       *rand.Rand
	inputValue string
	question   Question
}

func (t *template) Render() app.UI {
	return app.Div().Class("container-fluid").Body(
		app.Div().Body(
			app.Input().
				ID("seedInput").
				Type("text").
				Value(t.inputValue).
				OnChange(t.OnInputChange),
			app.Button().
				Text("Generate Seed").
				OnClick(t.GenerateSeed),
		),
		app.H5().Text("Input Seed Value: "+t.inputValue),
		app.H1().Text("question : "+t.question.statement1),
	)
}
func (t *template) OnInputChange(ctx app.Context, e app.Event) {
	t.inputValue = ctx.JSSrc().Get("value").String()
	t.Update()
}

func (t *template) GenerateSeed(ctx app.Context, e app.Event) {
	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
	fmt.Println("input value:", inputValue)
	seed, err := strconv.ParseInt(inputValue, 10, 64)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		seed = time.Now().UnixNano()
	}
	source := rand.NewSource(seed)
	t.rand = rand.New(source)
	t.question = t.createRelation()
	t.Update()
}

type thisQuestion struct {
	thisRelation [][2]int
	statement1   string
	statement2   string
	solution     string
}

func (t *template) createRelation() Question {

	partStatement := ""
	C := t.rand.Intn(2) + 1
	switch C {
	case 1:
		partStatement = "is double"
	case 2:
		partStatement = "is triple"
	case 3:
		partStatement = "is less than"
	default:
	}
	statement := fmt.Sprintf(" %s ", partStatement)
	return Question{
		statement1: statement,
	}
}
