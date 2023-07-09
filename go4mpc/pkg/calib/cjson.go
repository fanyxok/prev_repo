package calib

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

/*
*
Cost Value is in MicroSecond,
1 Second = 10**6 MicroSecond
*
*/

type Protocols struct {
	Cost map[string]float64 `json:"Cost"`
}
type OpCostJson struct {
	OpName string               `json:"OpName"`
	OpCost map[string]Protocols `json:"OpCost"`
}
type CostJson struct {
	Name string             `json:"Name"`
	Cost map[string]float64 `json:"Cost"`
}

func LoadOpCostJson(filename string) map[string]map[string]Protocols {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()
	byteResult, _ := io.ReadAll(file)
	costs := []OpCostJson{}
	json.Unmarshal(byteResult, &costs)
	costm := make(map[string]map[string]Protocols)
	for _, cost := range costs {
		costm[cost.OpName] = cost.OpCost
	}
	log.Println(costm)
	return costm
}

func LoadCostJson(filename string) map[string]map[string]float64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()
	byteResult, _ := io.ReadAll(file)
	costs := []CostJson{}
	json.Unmarshal(byteResult, &costs)
	costm := make(map[string]map[string]float64)
	for _, cost := range costs {
		costm[cost.Name] = cost.Cost
	}
	return costm
}

func WriteCostJson(filename string, costm map[string]map[string]float64) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	costs := []CostJson{}
	for name, cost := range costm {
		costs = append(costs, CostJson{Name: name, Cost: cost})
	}
	json.NewEncoder(file).Encode(costs)
}

func WriteOpCostJson(filename string, costm map[string]map[string]Protocols) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	costs := []OpCostJson{}
	for name, cost := range costm {
		costs = append(costs, OpCostJson{OpName: name, OpCost: cost})
	}
	json.NewEncoder(file).Encode(costs)
}
