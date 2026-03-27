package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	//定性查找：指返回有（true）和无（false）的查找
	//bytes包和strings提供了一组名字相同的定性查找API，包括Contains系列、HasPrefix和HasSuffix。

	// Contains函数
	fmt.Println(strings.Contains("Golang", "Go"))
	fmt.Println(strings.Contains("Golang", "go"))
	fmt.Println(strings.Contains("Golang", "l"))
	fmt.Println(strings.Contains("Golang", ""))
	fmt.Println(strings.Contains("", ""))

	fmt.Println("=============")

	fmt.Println(bytes.Contains([]byte("Golang"), []byte("Go")))
	fmt.Println(bytes.Contains([]byte("Golang"), []byte("go")))
	fmt.Println(bytes.Contains([]byte("Golang"), []byte("l")))
	fmt.Println(bytes.Contains([]byte("Golang"), []byte("")))
	fmt.Println(bytes.Contains([]byte("Golang"), nil))
	fmt.Println(bytes.Contains([]byte("Golang"), []byte{}))
	fmt.Println(bytes.Contains(nil, nil))

	//ContainsAny函数：如果两个集合存在不为空的交集，就返回true
	fmt.Println(strings.ContainsAny("Golang", "java"))   //true
	fmt.Println(strings.ContainsAny("Golang", "python")) //true
	fmt.Println(strings.ContainsAny("Golang", "c"))      //false
	fmt.Println(strings.ContainsAny("Golang", ""))       //false
	fmt.Println(strings.ContainsAny("", ""))             //false

	fmt.Println("=============")

	fmt.Println(bytes.ContainsAny([]byte("Golang"), "java")) //true
	fmt.Println(bytes.ContainsAny([]byte("Golang"), "c"))    //false
	fmt.Println(bytes.ContainsAny([]byte("Golang"), ""))     //false
	fmt.Println(bytes.ContainsAny(nil, ""))                  //false

	fmt.Println("================")

	//ContainseRune函数：判断某一个Unicode字符（以码点形式即rune类型值传入）是否包含在第一个参数代表的字符串或字节切片中。
	fmt.Println(strings.ContainsRune("Golang", 97))
	fmt.Println(strings.ContainsRune("Golang", rune('中')))
	fmt.Println(bytes.ContainsRune([]byte("Golang"), 97))
	fmt.Println(bytes.ContainsRune([]byte("Golang"), rune('中')))

	fmt.Println("================")
	//HasPrefix和HasSuffix函数
	//注意：空字符串（""）是任何字符串的前缀和后缀，空字节切片（​[​]byte{}）和nil切片也是任何字节切片的前缀和后缀
	fmt.Println(strings.HasPrefix("Golang", "Go"))     //true
	fmt.Println(strings.HasPrefix("Golang", "Golang")) //true
	fmt.Println(strings.HasPrefix("Golang", "lang"))   //false
	fmt.Println(strings.HasPrefix("Golang", ""))       //true
	fmt.Println(strings.HasPrefix("", ""))             //true
	fmt.Println(strings.HasSuffix("Golang", "Go"))     //false
	fmt.Println(strings.HasSuffix("Golang", "Golang")) //true
	fmt.Println(strings.HasSuffix("Golang", "lang"))   //true
	fmt.Println(strings.HasSuffix("Golang", ""))       //true
	fmt.Println(strings.HasSuffix("", ""))             //true

	fmt.Println(bytes.HasPrefix([]byte("Golang"), []byte("Go")))
	fmt.Println(bytes.HasPrefix([]byte("Golang"), []byte("Golang")))
	fmt.Println(bytes.HasPrefix([]byte("Golang"), []byte("lang")))
	fmt.Println(bytes.HasPrefix([]byte("Golang"), []byte{}))
	fmt.Println(bytes.HasPrefix([]byte("Golang"), nil))
	fmt.Println(bytes.HasPrefix(nil, nil))

	fmt.Println(bytes.HasSuffix([]byte("Golang"), []byte("Go")))
	fmt.Println(bytes.HasSuffix([]byte("Golang"), []byte("Golang")))
	fmt.Println(bytes.HasSuffix([]byte("Golang"), []byte("lang")))
	fmt.Println(bytes.HasSuffix([]byte("Golang"), []byte{}))
	fmt.Println(bytes.HasSuffix([]byte("Golang"), nil))
	fmt.Println(bytes.HasSuffix(nil, nil))

	fmt.Println("==========================")
	//定位查找
	//定位相关查找函数会给出第二个参数代表的字符串/字节切片在第一个参数中第一次出现的位置（下标）​，如果没有找到，则返回-1。
	// 另外定位查找还有方向性，从左到右为正向定位查找（Index系列）​，反之为反向定位查找（LastIndex系列）​。
	fmt.Println(strings.Index("Learn Golang, Go!", "Go"))
	fmt.Println(strings.Index("Learn Golang, Go!", ""))
	fmt.Println(strings.Index("Learn Golang, Go!", "Java"))
	fmt.Println(strings.IndexAny("Learn Golang, Go!", "Java"))
	fmt.Println(strings.IndexRune("Learn Golang, Go!", rune('a')))

	fmt.Println(bytes.Index([]byte("Learn Golang, Go!"), []byte("Go")))
	fmt.Println(bytes.Index([]byte("Learn Golang, Go!"), nil))
	fmt.Println(bytes.Index([]byte("Learn Golang, Go!"), []byte("Java")))
	fmt.Println(bytes.IndexAny([]byte("Learn Golang, Go!"), "Java"))
	fmt.Println(bytes.IndexRune([]byte("Learn Golang, Go!"), rune('a')))

	//反向定位查找（string）
	fmt.Println(strings.LastIndex("Learn Golang, Go!", "Go"))
	fmt.Println(strings.LastIndex("Learn Golang, Go!", ""))
	fmt.Println(strings.LastIndex("Learn Golang, Go!", "Java"))
	fmt.Println(strings.LastIndexAny("Learn Golang, Go!", "Java"))

	fmt.Println(bytes.LastIndex([]byte("Learn Golang, Go!"), []byte("Go")))
	fmt.Println(bytes.LastIndex([]byte("Learn Golang, Go!"), nil))
	fmt.Println(bytes.LastIndex([]byte("Learn Golang, Go!"), []byte("Java")))
	fmt.Println(bytes.LastIndexAny([]byte("Learn Golang, Go!"), "Java"))

	fmt.Println("===============")
	//Go标准库在strings包中提供了两种进行字符串替换的方法：Replace函数和Replacer类型。bytes包中则只提供了Replace函数用于字节切片的替换。
	fmt.Println(strings.Replace("I love java, java, java!!", "java", "go", 1))
	fmt.Println(strings.Replace("I love java, java, java!!", "java", "go", 2))
	fmt.Println(strings.Replace("I love java, java, java!!", "java", "go", -1))
	fmt.Println(strings.Replace("math", "", "go", -1))
	fmt.Println(strings.ReplaceAll("I love java, java, java!!", "java", "go"))
	replacer := strings.NewReplacer("java", "go", "python", "go")
	fmt.Println(replacer.Replace("I love java, python, go!!"))

	// 替换byte
	fmt.Printf("%s\n", bytes.Replace([]byte("I love java, java, java!!"), []byte("java"), []byte("go"), 1))
	fmt.Printf("%s\n", bytes.Replace([]byte("I love java, java, java!!"), []byte("java"), []byte("go"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("I love java, java, java!!"), []byte("java"), []byte("go"), -1))
	fmt.Printf("%s\n", bytes.Replace([]byte("math"), nil, []byte("go"), -1))
	fmt.Printf("%s\n", bytes.ReplaceAll([]byte("I love java, java, java!!"), []byte("java"), []byte("go")))
}
