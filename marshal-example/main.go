package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type Result interface {
	GetType() int
	GetJson() ([]byte, error)
}

type A struct {
	I int `json:"I"`
}

func (a A) GetType() int {
	return 1
}

func (a A) GetJson() ([]byte, error) {
	s := fmt.Sprint(a.I)
	return []byte(s), nil
}

type ResultSet struct {
	Results []Result `json:"results"`
}

func (rs *ResultSet) UnmarshalJSON(b []byte) error {
	var (
		err error
	)
	// naive approach
	prefix := []byte("{\"results\":")
	if !bytes.HasPrefix(b, prefix) {
		return errors.New("Unmarshal error")
	}
	bb := bytes.TrimSuffix(b[len(prefix):], []byte("}"))
	// Next part is hard-typed.
	// If need to work with dynamic types then provide special marshalJSON() with type definitions
	// such as {"results": {"type":"A"}, [{"I":0}...]}
	// and also provide your own type extraction mechanic
	rsResult := make([]A, 0)
	if err = json.Unmarshal(bb, &rsResult); err != nil {
		return err
	}
	for i := range rsResult {
		rs.Results = append(rs.Results, rsResult[i])
	}
	return nil
}

func main() {
	rs := ResultSet{}
	for i := 0; i < 10; i++ {
		rs.Results = append(rs.Results, A{i})
	}

	data, err := json.Marshal(rs)
	if err != nil {
		panic("1")
	}

	fmt.Println("json:", string(data))

	r := ResultSet{}
	err = json.Unmarshal(data, &r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("struct: %+v\n", r)
}
