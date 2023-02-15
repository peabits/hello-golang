package main

import (
	"hello-golang/go-by-example/example"
)

func main() {
	// rootDir := "/home/qwj/GitHub/hello-golang"
	// dirs, err := os.ReadDir(rootDir + "/go-by-example1")
	// if err == nil {
	// 	for _, dir := range dirs {
	// 		if !dir.IsDir() && strings.HasSuffix(dir.Name(), ".go") {
	// 			bytes, _ := os.ReadFile(rootDir + "/go-by-example1/" + dir.Name())
	// 			content := string(bytes)
	// 			name := strings.Split(dir.Name(), ".")[0]
	// 			names := strings.Split(name, "-")
	// 			for i, name := range names {
	// 				names[i] = cases.Title(language.Und, cases.NoLower).String(name)
	// 			}
	// 			name = strings.Join(names, "")
	// 			content = strings.Replace(content, "package main", "package example", 1)
	// 			content = strings.Replace(content, "main()", name+"()", 1)
	// 			os.WriteFile(rootDir+"/go-by-example/example/"+dir.Name(), []byte(content), 0666)
	// 		}
	// 	}
	// }
	example.HelloWorld()
}
