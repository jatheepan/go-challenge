package Backend_Challenge

import "testing"

func TestSum(t *testing.T) {
  response := Sum(MockData())
  expectedResponse := "45"

  if response != expectedResponse {
    t.Errorf("%s expected %s", response, expectedResponse)
  }
}
