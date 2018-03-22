package cart

// Cart model
type Cart struct {
	Name  string
	Total int
	Items []*Item
}

// Item model
type Item struct {
	Name  string
	Price int
}

// New will create new cart
func New(n string) *Cart {
	c := new(Cart)

	c.Name = n

	return c
}

// NewItem will create new cart item
func NewItem(n string, p int) *Item {
	i := new(Item)

	i.Name = n
	i.Price = p

	return i
}

// Add new item to cart
func (c *Cart) Add(it *Item) {
	c.Items = append(c.Items, it)
}

// Remove item from cart
func (c *Cart) Remove(it *Item) {
	for i, el := range c.Items {
		if it == el {
			copy(c.Items[i:], c.Items[i+1:])
			c.Items[len(c.Items)-1] = nil
			c.Items = c.Items[:len(c.Items)-1]
		}
	}
}
