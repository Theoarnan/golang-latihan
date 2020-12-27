package main

import "fmt"

func main() {
	// Test 1
	fmt.Println("Hello", "world!", "how", "are", "you")

	// Test 2
	// Membuat Variable dalam golang
	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	var firstName string = "john"
	var lastName string
	lastName = "wick"
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Printf("halo john wick!\n")
	fmt.Println("halo", firstName, lastName+"!")

	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	// var firstName = "john"
	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	// lastName := "wick"

	// Tipe Data
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	var messages string = "Halo"
	fmt.Printf("messages: %s \n", messages)

	var message = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`
	fmt.Println(message)

	const awalNama string = "john"
	fmt.Print("halo ", awalNama, "!\n")
	const akhirNama = "wick"
	fmt.Print("nice to meet you ", akhirNama, "!\n")
	fmt.Println("john wick")
	fmt.Println("john", "wick")
	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")

	// Operator
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	var left = false
	var right = true
	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)
	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)
	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)

	// Kondisi Seleksi
	var point = 8
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	var points = 8840.0
	if percent := points / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var poin = 6
	switch poin {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var pointss = 6
	switch pointss {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	var pointt = 6
	switch pointt {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	var pointts = 6
	switch {
	case pointts == 8:
		fmt.Println("perfect")
	case (pointts < 8) && (pointts > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	var pointtt = 6
	switch {
	case pointtt == 8:
		fmt.Println("perfect")
	case (pointtt < 8) && (pointtt > 3):
		fmt.Println("awesome")
		fallthrough
	case pointtt < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	var pointtts = 10
	if pointtts > 7 {
		switch pointtts {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if pointtts == 5 {
			fmt.Println("not bad")
		} else if pointtts == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if pointtts == 0 {
				fmt.Println("try harder!")
			}
		}
	}

	// Perulangan
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	var o = 0
	for o < 5 {
		fmt.Println("Angka", o)
		o++
	}

	var z = 0
	for {
		fmt.Println("Angka", z)
		z++
		if z == 5 {
			break
		}
	}

	for m := 1; m <= 10; m++ {
		if m%2 == 1 {
			continue
		}
		if m > 8 {
			break
		}
		fmt.Println("Angka", m)
	}

	for n := 0; n < 5; n++ {
		for j := n; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

outerLoop:
	for u := 0; u < 5; u++ {
		for v := 0; v < 5; v++ {
			if u == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", u, "][", v, "]", "\n")
		}
	}

	// Array
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// var fruitsss [4]string
	// cara horizontal
	// fruitsss = [4]string{"apple", "grape", "banana", "melon"}
	// cara vertikal
	// fruitsss = [4]string{
	// 	"apple",
	// 	"grape",
	// 	"banana",
	// 	"melon",
	// }

	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	var fruit = [4]string{"apple", "grape", "banana", "melon"}
	for iz := 0; iz < len(fruit); iz++ {
		fmt.Printf("elemen %d : %s\n", iz, fruit[iz])
	}

	var buahs = [4]string{"apple", "grape", "banana", "melon"}
	for ic, buah := range buahs {
		fmt.Printf("elemen %d : %s\n", ic, buah)
	}

	// var buahh = [4]string{"apple", "grape", "banana", "melon"}
	// for i, buahss := range buahh {
	// 	fmt.Printf("nama buah : %s\n", buahss)
	// }

	var bu = [4]string{"apple", "grape", "banana", "melon"}
	for _, buaz := range bu {
		fmt.Printf("nama buah : %s\n", buaz)
	}

	var bh = make([]string, 2)
	bh[0] = "apple"
	bh[1] = "manggo"
	fmt.Println(bh) // [apple manggo]

	// Slice
	var f = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(f[0]) // "apple"

	// var fruitsA = []string{"apple", "grape"}     // slice
	// var fruitsB = [2]string{"banana", "melon"}   // array
	// var fruitsC = [...]string{"papaya", "grape"} // array

	var fruitsD = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruitsD[0:2]
	fmt.Println(newFruits) // ["apple", "grape"]

	var fruitsG = []string{"apple", "grape", "banana", "melon"}
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]
	fmt.Println(fruitsG)  // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruitsG)  // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]

	var zfruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(zfruits)) // 4

	var kfruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(kfruits)) // len: 4
	fmt.Println(cap(kfruits)) // cap: 4

	var abFruits = kfruits[0:3]
	fmt.Println(len(abFruits)) // len: 3
	fmt.Println(cap(abFruits)) // cap: 4

	var bcFruits = kfruits[1:4]
	fmt.Println(len(bcFruits)) // len: 3
	fmt.Println(cap(bcFruits)) // cap: 3

	var pfruits = []string{"apple", "grape", "banana"}
	var lFruits = append(pfruits, "papaya")
	fmt.Println(pfruits) // ["apple", "grape", "banana"]
	fmt.Println(lFruits) // ["apple", "grape", "banana", "papaya"]

	var ff = []string{"apple", "grape", "banana"}
	var bff = ff[0:2]
	fmt.Println(cap(bff)) // 3
	fmt.Println(len(bff)) // 2
	fmt.Println(ff)       // ["apple", "grape", "banana"]
	fmt.Println(bff)      // ["apple", "grape"]
	var cff = append(bff, "papaya")
	fmt.Println(ff)  // ["apple", "grape", "papaya"]
	fmt.Println(bff) // ["apple", "grape"]
	fmt.Println(cff) // ["apple", "grape", "papaya"]

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)
	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	dsts := []string{"potato", "potato", "potato"}
	srcs := []string{"watermelon", "pinnaple"}
	ns := copy(dst, src)
	fmt.Println(dsts) // watermelon pinnaple potato
	fmt.Println(srcs) // watermelon pinnaple
	fmt.Println(ns)   // 2

	var fr = []string{"apple", "grape", "banana"}
	var afr = fr[0:2]
	var bfr = fr[0:2:2]
	fmt.Println(fr)       // ["apple", "grape", "banana"]
	fmt.Println(len(fr))  // len: 3
	fmt.Println(cap(fr))  // cap: 3
	fmt.Println(afr)      // ["apple", "grape"]
	fmt.Println(len(afr)) // len: 2
	fmt.Println(cap(afr)) // cap: 3
	fmt.Println(bfr)      // ["apple", "grape"]
	fmt.Println(len(bfr)) // len: 2
	fmt.Println(cap(bfr)) // cap: 2

	// MAP
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40
	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	var data map[string]int
	data["one"] = 1
	// akan muncul error!
	data = map[string]int{}
	data["one"] = 1
	// tidak ada error

	// cara vertikal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}
	// cara horizontal
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}

	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)
	var chickenS = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	for key, val := range chickenS {
		fmt.Println(key, "  \t:", val)
	}

	var chickenF = map[string]int{"januari": 50, "februari": 40}
	fmt.Println(len(chickenF)) // 2
	fmt.Println(chickenF)

	delete(chicken, "januari")

	fmt.Println(len(chickenF)) // 1
	fmt.Println(chickenF)

	var chickenM = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chickenM["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}

	var chickenv = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}

	for _, chickenv := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

}
