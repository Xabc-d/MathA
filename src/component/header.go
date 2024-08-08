package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Header struct {
	app.Compo
}

func (h *Header) Render() app.UI {
	return app.Div().Class("container-fluid").Style("background-color", "#F0F8FF").Style("height", "100vh").Body(
		app.Div().Class("row bg-primary").Style("min-height", "50px").Body(),
		app.Div().Class("row").Style("background-color", "#add8e6").Body(
			app.H1().Style("width", "90%").Style("margin", "0 auto").Body(
				app.Text("HelloMath"),
			),
		),
		&MathJaxConfig{},
		//app.P().Text("\\(\\frac{a}{b}\\)"),
		app.Div().Style("height", "30vh").Body(
			app.H6().
				Style("margin-top", "10vh").
				Style("font-size", "20px").
				Style("text-align", "center").
				Text("These questions are designed for COMS10014 Maths A"),
			app.H6().
				Style("margin-top", "10vh").
				Style("font-size", "18px").
				Style("text-align", "center").
				Text("Please note that these questions are primarily aimed to help you achieve a mark of 50"),
			app.H6().
				Style("margin-top", "2vh").
				Style("font-size", "18px").
				Style("text-align", "center").
				Text("for higher expectations please refer to more specialized sources"),
		),

		app.Div().Class("row").
			Style("justify-content", "center").
			Body(
				app.A().
					Href("week01").
					Class("clickable-square").
					Style("width", "20vh").
					Style("height", "20vh").
					Style("position", "relative").
					Style("background-color", "#add8e6").
					Style("display", "flex").
					Style("justify-content", "center").
					Style("align-items", "center").
					Style("margin", "3vh").
					Body(
						app.Img().
							Src("/web/func.svg").
							Style("width", "15vh").
							Style("height", "15vh").
							Style("position", "absolute").
							Style("top", "0").
							Alt("Embedded clickable image"),
						app.H5().
							Style("color", "black").
							Style("position", "absolute").
							Style("bottom", "0").
							Text("Terms"),
					),
				app.A().
					Href("week02").
					Class("clickable-square").
					Style("width", "20vh").
					Style("height", "20vh").
					Style("position", "relative").
					Style("background-color", "#add8e6").
					Style("display", "flex").
					Style("justify-content", "center").
					Style("align-items", "center").
					Style("margin", "3vh").
					Body(
						app.Img().
							Src("/web/func.svg").
							Style("width", "15vh").
							Style("height", "15vh").
							Style("position", "absolute").
							Style("top", "0").
							Alt("Embedded clickable image"),
						app.H5().
							Style("color", "black").
							Style("position", "absolute").
							Style("bottom", "0").
							Text("Logic"),
					),
				//app.A().
				//	Href("week03").
				//	Class("clickable-square").
				//	Style("width", "20vh").
				//	Style("height", "20vh").
				//	Style("position", "relative").
				//	Style("background-color", "#add8e6").
				//	Style("display", "flex").
				//	Style("justify-content", "center").
				//	Style("align-items", "center").
				//	Style("margin", "3vh").
				//	Body(
				//		app.Img().
				//			Src("/web/func.svg").
				//			Style("width", "15vh").
				//			Style("height", "15vh").
				//			Style("position", "absolute").
				//			Style("top", "0").
				//			Alt("Embedded clickable image"),
				//		app.H5().
				//			Style("color", "black").
				//			Style("position", "absolute").
				//			Style("bottom", "0").
				//			Text("Logic More"),
				//	),
				//app.A().
				//	Href("week04").
				//	Class("clickable-square").
				//	Style("width", "20vh").
				//	Style("height", "20vh").
				//	Style("position", "relative").
				//	Style("background-color", "#add8e6").
				//	Style("display", "flex").
				//	Style("justify-content", "center").
				//	Style("align-items", "center").
				//	Style("margin", "3vh").
				//	Body(
				//		app.Img().
				//			Src("/web/func.svg").
				//			Style("width", "15vh").
				//			Style("height", "15vh").
				//			Style("position", "absolute").
				//			Style("top", "0").
				//			Alt("Embedded clickable image"),
				//		app.H5().
				//			Style("color", "black").
				//			Style("position", "absolute").
				//			Style("bottom", "0").
				//			Text("Proofs"),
				//	),
				//app.A().
				//	Href("week05").
				//	Class("clickable-square").
				//	Style("width", "20vh").
				//	Style("height", "20vh").
				//	Style("position", "relative").
				//	Style("background-color", "#add8e6").
				//	Style("display", "flex").
				//	Style("justify-content", "center").
				//	Style("align-items", "center").
				//	Style("margin", "3vh").
				//	Body(
				//		app.Img().
				//			Src("/web/func.svg").
				//			Style("width", "15vh").
				//			Style("height", "15vh").
				//			Style("position", "absolute").
				//			Style("top", "0").
				//			Alt("Embedded clickable image"),
				//		app.H5().
				//			Style("color", "black").
				//			Style("position", "absolute").
				//			Style("bottom", "0").
				//			Text("Proofs More"),
				//	),
				//app.A().
				//	Href("week06").
				//	Class("clickable-square").
				//	Style("width", "20vh").
				//	Style("height", "20vh").
				//	Style("position", "relative").
				//	Style("background-color", "#add8e6").
				//	Style("display", "flex").
				//	Style("justify-content", "center").
				//	Style("align-items", "center").
				//	Style("margin", "3vh").
				//	Body(
				//		app.Img().
				//			Src("/web/func.svg").
				//			Style("width", "15vh").
				//			Style("height", "15vh").
				//			Style("position", "absolute").
				//			Style("top", "0").
				//			Alt("Embedded clickable image"),
				//		app.H5().
				//			Style("color", "black").
				//			Style("position", "absolute").
				//			Style("bottom", "0").
				//			Text("Induction"),
				//	),
				//app.A().
				//	Href("week07").
				//	Class("clickable-square").
				//	Style("width", "20vh").
				//	Style("height", "20vh").
				//	Style("position", "relative").
				//	Style("background-color", "#add8e6").
				//	Style("display", "flex").
				//	Style("justify-content", "center").
				//	Style("align-items", "center").
				//	Style("margin", "3vh").
				//	Body(
				//		app.Img().
				//			Src("/web/func.svg").
				//			Style("width", "15vh").
				//			Style("height", "15vh").
				//			Style("position", "absolute").
				//			Style("top", "0").
				//			Alt("Embedded clickable image"),
				//		app.H5().
				//			Style("color", "black").
				//			Style("position", "absolute").
				//			Style("bottom", "0").
				//			Text("Predicates"),
				//	),
				app.A().
					Href("week08").
					Class("clickable-square").
					Style("width", "20vh").
					Style("height", "20vh").
					Style("position", "relative").
					Style("background-color", "#add8e6").
					Style("display", "flex").
					Style("justify-content", "center").
					Style("align-items", "center").
					Style("margin", "3vh").
					Body(
						app.Img().
							Src("/web/func.svg").
							Style("width", "15vh").
							Style("height", "15vh").
							Style("position", "absolute").
							Style("top", "0").
							Alt("Embedded clickable image"),
						app.H5().
							Style("color", "black").
							Style("position", "absolute").
							Style("bottom", "0").
							Text("Sets"),
					),
				app.A().
					Href("week09").
					Class("clickable-square").
					Style("width", "20vh").
					Style("height", "20vh").
					Style("position", "relative").
					Style("background-color", "#add8e6").
					Style("display", "flex").
					Style("justify-content", "center").
					Style("align-items", "center").
					Style("margin", "3vh").
					Body(
						app.Img().
							Src("/web/func.svg").
							Style("width", "15vh").
							Style("height", "15vh").
							Style("position", "absolute").
							Style("top", "0").
							Alt("Embedded clickable image"),
						app.H5().
							Style("color", "black").
							Style("position", "absolute").
							Style("bottom", "0").
							Text("Functions"),
					),
				app.A().
					Href("week10").
					Class("clickable-square").
					Style("width", "20vh").
					Style("height", "20vh").
					Style("position", "relative").
					Style("background-color", "#add8e6").
					Style("display", "flex").
					Style("justify-content", "center").
					Style("align-items", "center").
					Style("margin", "3vh").
					Body(
						app.Img().
							Src("/web/func.svg").
							Style("width", "15vh").
							Style("height", "15vh").
							Style("position", "absolute").
							Style("top", "0").
							Alt("Embedded clickable image"),
						app.H5().
							Style("color", "black").
							Style("position", "absolute").
							Style("bottom", "0").
							Text("Relations"),
					),
				app.A().
					Href("all").
					Class("clickable-square").
					Style("width", "20vh").
					Style("height", "20vh").
					Style("position", "relative").
					Style("background-color", "#add8e6").
					Style("display", "flex").
					Style("justify-content", "center").
					Style("align-items", "center").
					Style("margin", "3vh").
					Body(
						app.Img().
							Src("/web/func.svg").
							Style("width", "15vh").
							Style("height", "15vh").
							Style("position", "absolute").
							Style("top", "0").
							Alt("Embedded clickable image"),
						app.H5().
							Style("color", "black").
							Style("position", "absolute").
							Style("bottom", "0").
							Text("All"),
					),
			),
	)
}
