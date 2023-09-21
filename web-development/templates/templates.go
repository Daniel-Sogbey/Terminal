package templates

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

type person struct {
	Name   string
	Gender string
}

type car struct {
	Model string
	Year  int
}

type design struct {
	Brand string
	Year  int
}

type items struct {
	Cars   []car
	People []person
}

var fm = template.FuncMap{
	"uc":       strings.ToUpper,
	"ft":       FirstThree,
	"fDateMDY": monthDayYear,
}

func FirstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func monthDayYear(t time.Time) string {
	return t.Format("2006/01/02")
}

func ParseTemplates() {
	// tpl := template.Must(template.ParseFiles("./templates/index.html"))

	// names := []string{"Daniel", "James", "John", "Paul"}

	// people := map[string]string{
	// 	"1": "Daniel",
	// 	"2": "John",
	// 	"3": "James",
	// 	"4": "Paul",
	// }

	// p1 := person{
	// 	"Daniel",
	// 	"Male",
	// }

	// people := []person{
	// 	{Name: "Daniel", Gender: "Male"},
	// 	{Name: "Diana", Gender: "Female"},
	// 	{Name: "Joe", Gender: "Male"},
	// }

	// newItems := items{
	// 	Cars: []car{
	// 		{Model: "Toyota,Matiz", Year: 2012},
	// 	},
	// 	People: []person{
	// 		{Name: "Daniel", Gender: "Male"},
	// 	},
	// }

	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("./templates/index.html"))

	// tpl.Execute(os.Stdout, 50)
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", time.Now())

	if err != nil {
		log.Fatalln(err)
	}

}
