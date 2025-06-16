package decorator

type TomatoTopping struct {
	Pizza IPizza
}

func (t TomatoTopping) GetPrince() int {
	pizzaPrice := t.Pizza.GetPrince()
	return pizzaPrice + 7
}
