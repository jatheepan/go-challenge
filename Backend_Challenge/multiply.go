package Backend_Challenge

import (
  "strconv"
)

func Multiply(records [][]string) string {
  var multiply int64 = 1
  for _, row := range records {
    for _, item := range row {
      f, err := strconv.ParseInt(item, 10, 64)
      if err == nil {
        multiply *= f
      }
    }
  }
  return strconv.FormatInt(multiply, 10)
}
