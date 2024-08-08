package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Sidebar struct {
	app.Compo
	SelectedItem string
}

func (s *Sidebar) onClick(ctx app.Context, e app.Event) {
	e.PreventDefault()
	v := ctx.JSSrc().Get("text").String()
	fmt.Println(v)
	s.SelectedItem = v
	s.Update()
}

func (s *Sidebar) Render() app.UI {
	return app.Div().Class("container-fluid").Body(
		app.Div().Class("row ").Body(

			app.Div().Class("text-bg-light sidebar").Style("height", "100vh").Style("border-right", "1px solid rgba(0, 0, 0, .1)").Body(
				app.Ul().Class("nav flex-column").Style("white-space", "nowrap").Style("overflow-x", "auto").Style("overflow-y", "auto").Body(

					app.Li().Class("nav-item").Style("max-width", "100%").Style("margin", "0 auto").Body(
						app.A().OnClick(s.onClick).Class("nav-link active").Href("week01").Body(
							app.Text("Terms"),
						),
					),
					app.Li().Class("nav-item").Style("max-width", "100%").Style("margin", "0 auto").Body(
						app.A().OnClick(s.onClick).Class("nav-link active").Href("week02").Body(
							app.Text("Logic"),
						),
					),

					//app.A().
					//	Href("week0122").
					//	Class("clickable-square").
					//	Style("width", "8vw").
					//	Style("height", "4vh").
					//	Style("position", "relative").
					//	Style("background-color", "#add8e6").
					//	Style("display", "flex").
					//	Style("justify-content", "center").
					//	Style("align-items", "center").
					//	Style("margin", "3vh").
					//	Body(
					//		app.Div().Class("row").Body(
					//			app.Div().Class("col").Style("position", "relative").Body(
					//				app.Img().
					//					Src("/web/func.svg").
					//					Style("width", "2vw").
					//					Style("height", "2vh").
					//					Style("position", "absolute").
					//					//Style("top", "0").
					//					Style("left", "0").
					//					Alt("Embedded clickable image"),
					//			),
					//			app.Div().Class("col").Style("position", "relative").Body(
					//				app.H5().
					//					Style("font-size", "2vw").
					//					Style("color", "black").
					//					Style("position", "absolute").
					//					Style("bottom", "0").
					//					Style("overflow", "hidden").
					//					Style("text-overflow", "ellipsis").
					//					Text("Function"),
					//			),
					//		),
					//	),
					//
					//
					//app.Li().Class("nav-item").Style("flex-grow", "1").Style("margin", "0 auto").Body(
					//	app.A().OnClick(s.onClick).Class("nav-link active").Href("week0122").Body(
					//		app.Text("week01"),
					//	),
					//),
					//
					app.Li().Class("nav-item").Style("max-width", "100%").Style("margin", "0 auto").Body(
						app.A().OnClick(s.onClick).Class("nav-link active").Href("week08").Body(
							app.Text("Sets"),
						),
					),
					app.Li().Class("nav-item").Style("max-width", "100%").Style("margin", "0 auto").Body(
						app.A().OnClick(s.onClick).Class("nav-link active").Href("week09").Body(
							app.Text("Function"),
						),
					),
					app.Li().Class("nav-item").Style("max-width", "100%").Style("margin", "0 auto").Body(
						app.A().OnClick(s.onClick).Class("nav-link active").Href("week10").Body(
							app.Text("Relations"),
						),
					),
					app.Li().Class("nav-item").Style("max-width", "100%").Style("margin", "0 auto").Body(
						app.A().OnClick(s.onClick).Class("nav-link active").Href("all").Body(
							app.Text("All"),
						),
					),
				),
			),
		),
	)
}
