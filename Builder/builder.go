package creational

type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

// Director creates objects by using a builder
type Director struct {
	builder Builder
}

// Construct tells the builder what to do and in which order
func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

// ConcreteBuilder implements Builder interface
type DocumentBuilder struct {
	product *Product
}

func (d *DocumentBuilder) MakeHeader(str string) {
	d.product.Content += "<header>" + str + "</header>"
}

func (d *DocumentBuilder) MakeBody(str string) {
	d.product.Content += "<body>" + str + "</body>"
}

func (d *DocumentBuilder) MakeFooter(str string) {
	d.product.Content += "<footer>" + str + "</footer>"
}

type Product struct {
	Content string
}

func (p *Product) Show() string {
	return p.Content
}
