package structs

type Product struct {
	Name  string
	Price float64
}

func PrintStructs() {

	println("***  begin structs ***")
	product := Product{
		Name:  "Laptop",
		Price: 2500.00,
	}

	println(product.Name)

	var productList []Product

	productList = append(productList, product)

	println(len(productList))

	println("***  end structs ***")
}
