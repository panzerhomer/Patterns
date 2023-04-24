package creational

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	test := "<header>Header</header>" +
		"<body>Body</body>" +
		"<footer>Footer</footer>"

	product := &Product{}

	director := Director{&DocumentBuilder{product}}
	director.Construct()

	result := product.Show()

	if result != test {
		t.Errorf("Expected %s, but got %s", result, test)
	}
}
