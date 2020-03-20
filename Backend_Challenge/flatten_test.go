package Backend_Challenge

import "testing"

func TestFlatten(t *testing.T) {
  response := Flatten(MockData())
  expectedResponse := "1,2,3,4,5,6,7,8,9"

  if response != expectedResponse {
    t.Errorf("%s expected %s", response, expectedResponse)
  }
}
