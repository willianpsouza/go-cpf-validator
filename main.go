package main

import (
	"fmt"
	"strconv"
	"strings"
)

func somaValor(dados []string, total int, maximo int) int {
	sm := 0
	digit := 0
	for i, s := range dados {
		if i > maximo {
			continue
		}
		nv, err := strconv.Atoi(s)
		if err != nil {
			return -1
		} else {
			tf := nv * total
			sm += tf
		}
		total -= 1
	}

	sm = 11 - (sm % 11)
	if sm == 11 || sm == 10 {
		digit = 0
	} else {
		digit = sm
	}
	return digit
}

func validaCpf(cpf string) bool {
	str := strings.Split(cpf, "")

	if len(str) != 11 {
		return false
	}

	o_digit10 := -1
	o_digit11 := -1

	nv, err := strconv.Atoi(str[9])
	if err == nil {
		o_digit10 = nv
	}

	nv, err = strconv.Atoi(str[10])
	if err == nil {
		o_digit11 = nv
	}

	digit10 := somaValor(str, 10, 8)
	digit11 := somaValor(str, 11, 9)

	if digit10 == o_digit10 && digit11 == o_digit11 {
		return true
	}
	return false
}

func geraDigito(cpf string) string {
	str := strings.Split(cpf, "")
	digito10 := strconv.Itoa(somaValor(str, 10, 8))
	_tempcpf := cpf + digito10
	str = strings.Split(_tempcpf, "")
	digito11 := strconv.Itoa(somaValor(str, 11, 9))
	ret := cpf + digito10 + digito11
	return ret
}

func main() {
	dados := [...]string{"12345678900", "11111111111", "22222222222"}
	status := ""
	for _, s := range dados {
		p := validaCpf(s)
		if p {
			status = "OK"
		} else {
			status = "NOK"
		}
		fmt.Printf("O cpf %s retornou %s \n", s, status)

	}
	dados = [...]string{"123456789", "987654321", "147258369123456789"}
	for _, p := range dados {
		s := geraDigito(p)
		fmt.Println(s, " -- ", p)
	}
	dados2 := []string{}
	for i := 0; i < 999999; i++ {
		d := fmt.Sprintf("%09d", i)
		dados2 = append(dados2, d)
	}

	for _, p := range dados2 {
		s := geraDigito(p)
		fmt.Println(s, " -- ", p)
	}

}
