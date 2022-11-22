package error_handling

import "fmt"

type borderExeption struct {
	parameter int
	message   string
}

func (b *borderExeption) Error() string {
	return fmt.Sprint("%d --- %s", b.parameter, b.message)
}

func TahminEt(tahmin int) (string, error) {
	if tahmin > 100 || tahmin < 1 {
		return "", &borderExeption{tahmin, "Sınır dışı"}
	}

	return "Bildiniz", nil
}
