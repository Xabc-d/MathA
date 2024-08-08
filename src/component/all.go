package component

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type All struct {
	Bases
}

func (all *All) SetQuestion() {
	source := rand.NewSource(all.seed)
	all.rand = rand.New(source)
	terms := &Terms{}
	logic := &Logic{}
	function := &Functions{}
	relation := &Relation{}
	set := &Set{}
	all.question2 = terms.Te1()
	all.question1 = all.al()
	group1 := []func() Question{
		terms.Te1,
		terms.te3,
		terms.te4,
		terms.te8,
		set.set1,
		set.set2,
		set.set4,
		set.set6,
		set.set7,
		relation.createRelation,
		relation.inverse,
		relation.inverseAlp,
		relation.inverseFix,
		relation.reflexiveSingle,
		relation.antisymmetricSingle,
		relation.transitiveSingle,
		relation.symmetricSingle,
		function.func1,
		function.func2,
		function.func3,
		function.func4,
		logic.lo1,
		logic.lo4,
		logic.lo2,
	}
	group2 := []func() Question{
		terms.te7,
		terms.te6,
		set.set3,
		relation.composition,
		relation.partialOrder,
		logic.lo3,
		logic.lo5,
		function.func5,

		terms.te5,
	}
	all.rand.Shuffle(len(group1), func(i, j int) {
		group1[i], group1[j] = group1[j], group1[i]
	})
	all.rand.Shuffle(len(group2), func(i, j int) {
		group2[i], group2[j] = group2[j], group2[i]
	})
	all.theLevel = all.seed % 10

}

func (all *All) GenerateSeed(ctx app.Context, e app.Event) {
	inputValue := app.Window().GetElementByID("seedInput").Get("value").String()
	fmt.Println("input value:", inputValue)
	seed, err := strconv.ParseInt(inputValue, 10, 64)
	all.seed = seed
	if err != nil {
		fmt.Println("No input or error converting string to integer:", err)
		seconds := time.Now().UnixNano() / 1000000
		all.seed = (seconds % 9000) + 1000
		//seed = time.Now().UnixNano()/1000000 - 1713220000000
		//one day 86400 seconds
		all.currentValue = strconv.FormatInt(all.seed, 10) + " (No input or error input use random seed)"

	} else {
		all.currentValue = strconv.FormatInt(all.seed, 10)
	}
	all.answerClicked = false

	all.SetQuestion()
	all.seedClicked = true
	all.Update()
}

func (all *All) al() Question {
	a := all.rand.Intn(30) + 1
	b := all.rand.Intn(30) + 1
	c := all.rand.Intn(30) + 1
	d := all.rand.Intn(30) + 1

	statement := fmt.Sprintf("Given the syntax rule for the Associative Law of addition: (a+b)+c ⊢ a+(b+c), perform a derivation to simplify the expression ((%v+%v)+%v)+%v?", a, b, c, d)
	theAnswer := fmt.Sprintf("%v+(%v+(%v+%v))", a, b, c, d)
	question := Question{
		statement1:   statement,
		answer:       theAnswer,
		questionType: 1,
	}
	return question
}
