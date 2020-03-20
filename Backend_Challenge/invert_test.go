package Backend_Challenge

import "testing"

func TestInvert(t *testing.T) {
  response := Invert(MockData())
  expectedResponse := "1,4,7\n2,5,8\n3,6,9\n"

  if response != expectedResponse {
    t.Errorf("%s expected %s", response, expectedResponse)
  }
}
