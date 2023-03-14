package exercicies

import (
	"io"
	"os"
	"strings"
)

/*
Exercício: Leitor Rot13
Um padrão comum é um io.Reader que envolve outro io.Reader, modificando o fluxo de alguma forma.

Por exemplo, a função gzip.NewReader recebe um io.Reader (um fluxo de dados compactado) e retorna um *gzip.Reader que também implementa io.Reader (um fluxo de dados descomprimidos).

Implemente um rot13Reader que implementa um io.Reader que leia a partir de um io.Reader , modificando o fluxo através da aplicação da cifra rot13 de substituição para todos os caracteres alfabéticos.

O tipo rot13Reader é fornecido para você. Torne-o em um io.Reader através da implementação de seu método Read.
*/

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)

	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}

	return
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

func Rot13Run() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
