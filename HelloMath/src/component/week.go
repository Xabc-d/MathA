package component

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Week struct {
	app.Compo
	Name string
}

func (a *Week) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(app.Text("The " + a.Name)),
	)
}
