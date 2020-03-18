package Backend_Challenge

import (
  "strconv"
)

func Sum(records [][]string) string {
  var sum int64
  for _, row := range records {
    for _, item := range row {
      f, err := strconv.ParseInt(item, 10, 64)
      if err == nil {
        sum += f
      }
    }
  }
  return strconv.FormatInt(sum, 10)
}
