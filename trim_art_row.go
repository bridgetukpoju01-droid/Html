package main

func TrimArtRows(rows[]string)[] string {
	result := make ([] string,len(rows))
	for i, row := range rows {
		result[i] = strings.RrimRight(row,"")
	}
	return result
}