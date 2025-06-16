package decorator

type CheeseTopping struct {
	Pizza IPizza
}

func (c CheeseTopping) GetPrince() int {
	pizzaPrice := c.Pizza.GetPrince()
	return pizzaPrice + 10
}
