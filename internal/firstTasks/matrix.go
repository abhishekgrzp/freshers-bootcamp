package firsttasks

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Row      int
	Column   int
	Elements [][]int
}

func newMatrix(rows int, column int) *Matrix {
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, column)
	}
	return &Matrix{Row: rows, Column: column, Elements: data}
}

func (m *Matrix) GetRows() int {
	return m.Row
}

func (m *Matrix) GetColumns() int {
	return m.Column
}

func (m *Matrix) PrintMatrix() {
	jsonData, err := json.Marshal(m.Elements)
	if err != nil {
		fmt.Println("error=>", err)
		return
	}
	fmt.Printf("Matrix: %v\n", string(jsonData))
}

func (m *Matrix) AddValue(value int, i, j int) {
	m.Elements[i][j] = value
}
