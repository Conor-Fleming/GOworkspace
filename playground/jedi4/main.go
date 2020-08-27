package main

import "fmt"

func main() {
	//exercise 1
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	for _, v := range arr {
		fmt.Print(v, "\n")
	}
	fmt.Printf("%T\n", arr)

	//exercise 2
	slice1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, v := range slice1 {
		fmt.Print(v, "\n")
	}
	fmt.Printf("%T\n", slice1)

	//exercise 3
	newSlce1 := slice1[0:5]
	fmt.Println(newSlce1)
	newSlce2 := slice1[5:]
	fmt.Println(newSlce2)
	newSlce3 := slice1[2:7]
	fmt.Println(newSlce3)
	newSlce4 := slice1[1:6]
	fmt.Println(newSlce4)

	//exercise4
	slice1 = append(slice1, 52)
	slice1 = append(slice1, 53, 54, 55)
	fmt.Println(slice1)
	y := []int{56, 57, 58, 59, 60}
	slice1 = append(slice1, y...)
	fmt.Println(slice1)

	//exercise5
	ex5Slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	ex5Slice = append(ex5Slice[:3], ex5Slice[6:]...)
	fmt.Println(ex5Slice)

	//exercise6
	states := []string{`Alabama`, `Alaska`, `Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(len(states))
	fmt.Println(cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

	//exercise7
	dubSlice := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	for i, v := range dubSlice {
		fmt.Println(v)
		for _, k := range dubSlice[i] {
			fmt.Print(k, "\n")
		}
	}

	//exercise8
	var m = map[string][]string{
		"fleming_conor":  {`bikes`, `beer`, `cars`},
		"kyne_ciara":     {`buying things`, `drawing`, `hamilton`},
		"fleming_cormac": {`clout`, `chew`},
	}
	/*for _, v := range m {
		for k, w := range v {
			fmt.Print(k, " ", w, "\n")
		}
	}*/

	//exercise 9
	m["fleming_karen"] = []string{"her kids,", "her car", "ireland"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "fleming_cormac")
	fmt.Println(m)
}
