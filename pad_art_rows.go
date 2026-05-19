package main
func PadArtRows(rowa[]string,width int)[] string {
	result := make ([]string,len (rows))
	for i, row := range rows{
		if width <= 0 || len (row)>= width {
			result[i] = row
			continue
		}
		padding := strings.Repeat("",width-len(row))
		result[i] = row = padding 
	}
	return result
}