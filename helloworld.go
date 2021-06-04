package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	category := [4]string{"fashion", "electronics", "sport", "food"}
	products := [5]Product{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(products); i++ {
		products[i] = Product{
			fmt.Sprintf("%s %d", "Product", i),
			category[rand.Intn(len(category))],
			randomInt(100, 200),
		}
	}
	sort.Sort(ByPrice(products[:]))
	for j := len(products) - 1; j >= 45; j-- {
		fmt.Println(products[j])
	}
	myMap := make(map[string]int)
	for _, product := range products {
		myMap[product.category]++
	}
	fmt.Println(myMap)
	// fmt.Println(products[4].String())

	fmt.Println(products)

	fmt.Println("Find By Name:")
	var ans1 Product
	var count int = 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the name: ")
	scanner.Scan()
	name := scanner.Text()
	for _, product := range products {
		if strings.EqualFold(product.name, name) {
			ans1 = product
		} else {
			count++
		}
	}
	fmt.Println(ans1)
	if count == len(products) {
		fmt.Println("Không có sản phẩm nào")
	}

	var count1 int = 0
	fmt.Println("Find by Category:")
	fmt.Print("Enter the category:")
	scanner.Scan()
	category1 := scanner.Text()
	for _, product := range products {
		if strings.EqualFold(product.category, category1) {
			fmt.Println(product)
			count1++
		}
	}
	if count1 == 0 {
		fmt.Println("Không tồn tại danh mục này")
	}

	var count2 int = 0
	fmt.Println("Find in the price range:")
	fmt.Print("Enter min price:")
	scanner.Scan()
	priceMin := scanner.Text()
	fmt.Print("Enter max price:")
	scanner.Scan()
	priceMax := scanner.Text()
	minToNum, _ := strconv.Atoi(priceMin)
	maxToNum, _ := strconv.Atoi(priceMax)
	for _, product := range products {
		if product.price >= minToNum && product.price <= maxToNum {
			fmt.Println(product)
			count2++
		}
	}
	if count2 == 0 {
		fmt.Println("Không tồn tại sản phẩm trong khoảng giá này")
	}
}

// func findProductByName() []string {
// 	reader := bufio.NewReader(os.Stdin)
//     fmt.Print("Enter the name: ")
//    	name, _ := reader.ReadString('\n')
//     fmt.Print("You live in " + city)

// }

type Product struct {
	name     string
	category string
	price    int
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func NewPerson(name string, category string, price int) *Product {
	p := new(Product)
	p.name = name
	p.category = category
	p.price = price
	return p
}

type ByPrice []Product

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Less(i, j int) bool { return a[i].price < a[j].price }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (p *Product) FullName() string { //Pointer receiver
	return p.name + " " + p.category
}

// func (p Product) String() string { //Value receiver
// 	return fmt.Sprintf("%v is %v years old", p.FullName(), p.price)
// }
