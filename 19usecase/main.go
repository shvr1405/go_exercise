package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("hello")

	ProductFamilies := make(map[string][]string)
	ProductFamilies["sr"] = []string{"7750", "7950", "7450"}
	fmt.Println("Initialization is done")

	//printing all the product family and node types
	for key, val := range ProductFamilies {
		fmt.Printf("The node types for product family is %v : %v\n", key, val)
	}
	fmt.Println("--------------------------")
	//adding new family
	ProductFamilies["sas"] = []string{"7710", "9999"}
	ProductFamilies["omni"] = []string{"7705", "7705-hm"}

	for key, val := range ProductFamilies {
		fmt.Printf("The node types for product family before performing operations are  %v : %v\n", key, val)
	}

	fmt.Println("--------------------------")

	//checking if a (7250) node is present or not. appending if not present
	for key, val := range ProductFamilies {
		if key == "sr" && !slices.Contains(ProductFamilies[key], "7250") {
			ProductFamilies[key] = append(val, "7250")
		}

		//deleting "9999" as depreciated node
		if slices.Contains(val, "9999") {
			for i, _ := range val {
				if val[i] == "9999" {
					ProductFamilies[key] = append(val[:i], val[i+1:]...)
				}
			}
		}
	}

	//printing final map
	for key, val := range ProductFamilies {
		fmt.Printf("The node types for product family after performing operations is %v : %v\n", key, val)
	}

}
