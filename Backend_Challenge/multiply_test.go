package Backend_Challenge

import "testing"

func TestMultiply(t *testing.T) {
  response := Multiply(MockData())
  expectedResponse := "362880"

  if response != expectedResponse {
    t.Errorf("%s expected %s", response, expectedResponse)
  }
}
