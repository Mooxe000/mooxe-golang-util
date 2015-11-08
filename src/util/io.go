package util

import (
  "os"
  "io/ioutil"
)

func FileStrFromPath(filePath string) (fileStr string) {
  file, _ := os.Open(filePath)
	defer file.Close()
  reader, _ := ioutil.ReadAll(file)
  return string(reader)
}
