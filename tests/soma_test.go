package soma

import (
	"testing"
)

func TestSumCorrect(t *testing.T)  {
	teste := soma(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Errorf("Resultado esperado: %d, resultado obtido: %d", resultado, teste)
	}
}

func TestSumFailed(t *testing.T)  {
	teste := soma(3, 2, 2)
	resultado := 6

	if teste == resultado {
		t.Errorf("Resultado esperado: %d, resultado obtido: %d", resultado, teste)
	}
}

func TestMultiplyCorrect(t *testing.T)  {
	teste := multiplica(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Errorf("Resultado esperado: %d, resultado obtido: %d", resultado, teste)
	}
}

func TestMultiplyFailed(t *testing.T)  {
	teste := multiplica(3, 2, 2)
	resultado := 6

	if teste == resultado {
		t.Errorf("Resultado esperado: %d, resultado obtido: %d", resultado, teste)
	}
}