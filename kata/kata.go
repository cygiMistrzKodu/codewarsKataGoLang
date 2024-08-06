package kata


func WordsToMarks(s string) int {

	count := 0
  for _, i := range s {
     count += int(i) - 'a' + 1;
  }
  return count
}
