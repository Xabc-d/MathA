package component

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type MathJaxConfig struct {
	app.Compo
}

func (m *MathJaxConfig) Render() app.UI {
	return app.Div().Body(
		app.Script().Src("https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-mml-chtml.js"),
		app.Script().Text(`
            window.MathJax = {
                tex: {
                    inlineMath: [['$', '$'], ['\\(', '\\)']]
                },
                svg: {
                    fontCache: 'global'
                }
            };
        `),
	)
}
