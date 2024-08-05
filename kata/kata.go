package kata

import "strings"

func RakeGarden(garden string) string {

	gardenElements := strings.Split(garden, " ")

  for index := range gardenElements {
    
    if (gardenElements[index] != "gravel" && gardenElements[index] != "rock" ) {
      gardenElements[index] = "gravel"
    }
    
  }

	return strings.Join(gardenElements, " ")
}
