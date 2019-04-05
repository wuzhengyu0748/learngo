package basic

import "fmt"

func main() {
	m := map[string]string {
		"name" : "rainy",
		"course" : "golang",
		"site" : "github",
		"quality" : "notbad",
	}

	m2 := make(map[string]string)

	var m3 map[string]int

	fmt.Println(m, m2, m3)

	fmt.Println("Traversting map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	if normal, ok := m["normal"]; ok {
		fmt.Println(normal)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
