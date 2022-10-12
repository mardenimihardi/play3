package delivery

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"

	usecase "play3/internal/usecase"
)

type Delivery struct {
	Usecase    usecase.UsecaseInterface
	Middleware MiddlewareInferface
}

func InitDelivery(usecase usecase.Usecase) {
	d := Delivery{Usecase: usecase, Middleware: &Middleware{}}
	d.RegisterRoutes()
}

func (d *Delivery) RegisterRoutes() {
	d.Middleware.GET("/", d.HtmlIndex)
}

func (d *Delivery) HtmlIndex(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("html", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dishes []map[string]interface{}
	b, _ := json.Marshal(d.Usecase.GetAllDishes())
	json.Unmarshal(b, &dishes)

	var data = map[string]interface{}{"data": dishes}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
