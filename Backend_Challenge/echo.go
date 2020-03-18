package Backend_Challenge

import (
  "fmt"
  "strings"
)

func Echo(records [][]string) string {
  var response string
  for _, row := range records {
    response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
  }
  return response
}
