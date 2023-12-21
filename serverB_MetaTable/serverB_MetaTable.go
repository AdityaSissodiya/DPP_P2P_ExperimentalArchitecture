package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var dbBatteries *sql.DB

func init() {
	var err error
	dbBatteries, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dbBatteries.Exec(`CREATE TABLE IF NOT EXISTS LithiumIonBatteries (
		BatteryID INTEGER PRIMARY KEY,
		ProductID INTEGER REFERENCES Products(ProductID) ON DELETE CASCADE,
		Capacity_mAh INTEGER,
		Voltage_V DECIMAL(4, 2),
		Chemistry TEXT,
		Rechargeable BOOLEAN,
		ChargeCycleCount INTEGER,
		CellType TEXT,
		EnergyDensity_WhPerKg DECIMAL(6, 2),
		CONSTRAINT FK_ProductID FOREIGN KEY (ProductID) REFERENCES Products(ProductID) ON DELETE CASCADE
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Insert sample data
	_, err = dbBatteries.Exec(`INSERT INTO LithiumIonBatteries (ProductID, Capacity_mAh, Voltage_V, Chemistry, Rechargeable, ChargeCycleCount, CellType, EnergyDensity_WhPerKg) VALUES
		(1, 5000, 3.7, 'Li-ion', true, 500, '18650', 250),
		(2, 3000, 3.8, 'Li-polymer', true, 300, 'AAA', 200)`)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/batteries", getBatteries).Methods("GET")

	fmt.Println("Server B (Batteries) is running on :8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}

func getBatteries(w http.ResponseWriter, r *http.Request) {
	rows, err := dbBatteries.Query("SELECT BatteryID, ProductID, Capacity_mAh, Voltage_V, Chemistry, Rechargeable, ChargeCycleCount, CellType, EnergyDensity_WhPerKg FROM LithiumIonBatteries")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var batteries []Battery
	for rows.Next() {
		var battery Battery
		err := rows.Scan(&battery.BatteryID, &battery.ProductID, &battery.Capacity_mAh, &battery.Voltage_V, &battery.Chemistry, &battery.Rechargeable, &battery.ChargeCycleCount, &battery.CellType, &battery.EnergyDensity_WhPerKg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		batteries = append(batteries, battery)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(batteries)
}

// Battery represents the structure of the LithiumIonBatteries table
type Battery struct {
	BatteryID             int     `json:"batteryID"`
	ProductID             int     `json:"productID"`
	Capacity_mAh          int     `json:"capacity_mAh"`
	Voltage_V             float64 `json:"voltage_V"`
	Chemistry             string  `json:"chemistry"`
	Rechargeable          bool    `json:"rechargeable"`
	ChargeCycleCount      int     `json:"chargeCycleCount"`
	CellType              string  `json:"cellType"`
	EnergyDensity_WhPerKg float64 `json:"energyDensity_WhPerKg"`
}
