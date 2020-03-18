package Backend_Challenge

import (
  "fmt"
  "strings"
)

func Flatten(records [][]string) string {
  var response string
  for index, row := range records {
    separator := ","
    if (index +1) % len(records) == 0 {
      separator = ""
    }
    response = fmt.Sprintf("%s%s%s", response, strings.Join(row, ","), separator)
  }
  return response
}
