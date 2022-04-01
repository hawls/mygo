package main
import "fmt"
func main(){
	var countryCapitalMap map[string]string = make(map[string]string)
	
	countryCapitalMap [ "France" ] = "巴黎"
    countryCapitalMap [ "Italy" ] = "罗马"
    countryCapitalMap [ "Japan" ] = "东京"
    countryCapitalMap [ "India " ] = "新德里"

	for k,v := range countryCapitalMap {
        fmt.Println(k, "首都是", v)
    }
	// delete(countryCapitalMap, "France")

	// capital, ok := countryCapitalMap [ "France" ]
	// if ok {
	// 	fmt.Println(countryCapitalMap["France"], capital)
	// }else{
	// 	fmt.Println(countryCapitalMap["France"]+"值不存在")
	// }
}