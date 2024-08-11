package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", handleRequest)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		desiredProfit, err := strconv.ParseFloat(r.FormValue("desiredProfit"), 64)
		if err != nil {
			http.Error(w, "Invalid desired profit", http.StatusBadRequest)
			return
		}

		monthlyCost, err := strconv.ParseFloat(r.FormValue("monthlyCost"), 64)
		if err != nil {
			http.Error(w, "Invalid monthly cost", http.StatusBadRequest)
			return
		}

		targetKm, err := strconv.ParseFloat(r.FormValue("targetKm"), 64)
		if err != nil {
			http.Error(w, "Invalid target km", http.StatusBadRequest)
			return
		}

		kmPerLiter, err := strconv.ParseFloat(r.FormValue("kmPerLiter"), 64)
		if err != nil {
			http.Error(w, "Invalid km per liter", http.StatusBadRequest)
			return
		}

		fuelCostPerLiter, err := strconv.ParseFloat(r.FormValue("fuelCostPerLiter"), 64)
		if err != nil {
			http.Error(w, "Invalid fuel cost per liter", http.StatusBadRequest)
			return
		}

		costPerKm := (monthlyCost / targetKm) + (fuelCostPerLiter / kmPerLiter)
		minRatePerKm := (desiredProfit + monthlyCost) / targetKm
		profitPerKm := minRatePerKm - costPerKm

		data := struct {
			DesiredMonthlyProfit float64
			MonthlyCost          float64
			KmPerLiter           float64
			FuelCostPerLiter     float64
			TargetKm             float64
			CostPerKm            float64
			MinRatePerKm         float64
			ProfitPerKm          float64
		}{
			DesiredMonthlyProfit: desiredProfit,
			MonthlyCost:          monthlyCost,
			KmPerLiter:           kmPerLiter,
			FuelCostPerLiter:     fuelCostPerLiter,
			TargetKm:             targetKm,
			CostPerKm:            costPerKm,
			MinRatePerKm:         minRatePerKm,
			ProfitPerKm:          profitPerKm,
		}

		tmpl.Execute(w, data)
		return
	}

	tmpl.Execute(w, nil)
}
