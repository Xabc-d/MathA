package component

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Hello struct {
	app.Compo
	Sidebar     Sidebar
	Name        string
	CurrentWeek string
}

func (h *Hello) OnNav(ctx app.Context, e app.Event) {
	h.CurrentWeek = e.Value.String()
	h.Update()
}

func (h *Hello) Render() app.UI {
	//var theWeek app.UI
	//if h.CurrentWeek == "week01" {
	//	theWeek = &function{}
	//} else {
	//	//theWeek = &relation{}
	//}

	return app.Main().Body(
		&Header{},
		app.Div().Class("row").Body(
		//app.Div().Class("col-2").Body(
		//	&Sidebar{}),
		//app.Div().Class("col-10").Body(
		//	theWeek),
		//&relation{}),
		//&relation{}),
		//&function{}),
		//&predicates{}),
		),
		//&Header{},
		//&h.Sidebar,
		//&Sidebar{},
		//&Calculator{},
		//&Week{Name: h.Sidebar.SelectedItem},
	)
}
