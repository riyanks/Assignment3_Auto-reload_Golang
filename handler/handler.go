package handler

import (
	"html/template"
	"net/http"
	"path"

	"github.com/husfuu/go-auto-reload/helper"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	var FilePath = path.Join("view", "index.html")

	var tmpl, err = template.ParseFiles(FilePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	disaster := helper.GetDisaster()
	Waternumber := disaster.Status.Water
	Windnumber := disaster.Status.Wind

	waterDesc := helper.WaterStatus(Waternumber)
	windDesc := helper.WindStatus(Windnumber)

	var statusDesc = map[string]interface{}{
		"waterDesc": waterDesc,
		"windDesc":  windDesc,
	}

	err = tmpl.Execute(w, statusDesc)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
