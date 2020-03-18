package Backend_Challenge

import (
  "fmt"
  "math"
)

func Invert(records [][]string) string {
  recordsCount := math.Pow(float64(len(records)), 2)
  var response string
  for i := float64(0); i < recordsCount; {
    separator := ","
    colIndex := int(math.Floor(i / float64(len(records))))
    rowIndex := int(i) % len(records)
    if rowIndex == len(records) -1 {
      separator = "\n"
    }
    item := records[rowIndex][colIndex]
    response = fmt.Sprintf("%s%s%s", response, item, separator)
    i += 1
  }

  return response
}
