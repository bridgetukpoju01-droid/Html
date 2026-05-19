package main
func StackTwo(top[]string,bottom[]string) [] string {
	var result [] string {}
	result = append(result,top...)
	result = append(result,bottom...)
	return result
}


func StackAll(blocks[][] string) [] string {
	result := [] string{}
	for _,block := append blocks {
		result = append(result,block...)
	}
	return result
}