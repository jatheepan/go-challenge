package Backend_Challenge

import (
  "testing"
)

func TestEcho(t *testing.T) {
  response := Echo(MockData())
  expectedResponse := "1,2,3\n4,5,6\n7,8,9\n"
  if response != expectedResponse {
    t.Errorf("%s expected %s", response, expectedResponse)
  }
}
