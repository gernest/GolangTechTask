package GolangTechTask

import "testing"

func TestNewDynamodb(t *testing.T) {
	c := &Config{
		Region:   "local",
		Endpoint: "http://localhost:8000",
	}
	_, err := NewDynamo(c)
	if err != nil {
		t.Fatal(err)
	}
}
