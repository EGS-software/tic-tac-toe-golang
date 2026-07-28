package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Jogo struct {
	tabuleiro [3][3]string
}

func main() {
	// Inicializa a semente para geração de números aleatórios
	rand.Seed(time.Now().UnixNano())

	jogo := Jogo{}
	jogo.inicializarTabuleiro()

	fmt.Println("=== JOGO DA VELHA ===")

	vezDoJogador := jogo.sortearQuemComeca()

	if vezDoJogador {
		fmt.Println("O jogador X começará a partida.")
	} else {
		fmt.Println("O computador O começará a partida.")
	}

	jogo.exibirTabuleiro()

	// Loop principal da partida
	for {
		if vezDoJogador {
			jogo.jogadaDoJogador()
			jogo.exibirTabuleiro()
			if jogo.verificarVitoria("X") {
				fmt.Println("O jogador X venceu a partida.")
				break
			}
		} else {
			jogo.jogadaDoComputador()
			jogo.exibirTabuleiro()
			if jogo.verificarVitoria("O") {
				fmt.Println("O computador O venceu a partida.")
				break
			}
		}

		if jogo.tabuleiroCheio() {
			fmt.Println("A partida terminou empatada.")
			break
		}

		// Alterna a vez
		vezDoJogador = !vezDoJogador
	}
}

// 1. Inicializa o tabuleiro com espaços vazios
func (j *Jogo) inicializarTabuleiro() {
	for l := 0; l < 3; l++ {
		for c := 0; c < 3; c++ {
			j.tabuleiro[l][c] = " "
		}
	}
}

// 2. Exibe o tabuleiro formatado
func (j *Jogo) exibirTabuleiro() {
	fmt.Println("\n         Colunas")
	fmt.Println("        1       2      3")
	fmt.Println("     +---+---+---+")
	for l := 0; l < 3; l++ {
		fmt.Printf("  %d  |", l+1)
		for c := 0; c < 3; c++ {
			fmt.Printf(" %s |", j.tabuleiro[l][c])
		}
		fmt.Println("\n     +---+---+---+")
	}
	fmt.Println()
}

// 3. Sorteia quem começa (retorna true para Jogador, false para Computador)
func (j *Jogo) sortearQuemComeca() bool {
	return rand.Intn(2) == 0
}

// 4. Lida com a entrada de dados e validação do jogador humano
func (j *Jogo) jogadaDoJogador() {
	var linha, coluna int

	for {
		fmt.Print("Informe a linha: ")
		_, errL := fmt.Scan(&linha)
		fmt.Print("Informe a coluna: ")
		_, errC := fmt.Scan(&coluna)

		// Verifica se a entrada foi um número
		if errL != nil || errC != nil {
			fmt.Println("Entrada inválida. Digite apenas números.")
			// Limpa o buffer em caso de erro de leitura
			var lixo string
			fmt.Scanln(&lixo)
			continue
		}

		if !j.posicaoValida(linha, coluna) {
			fmt.Println("Posição inválida. A linha e a coluna devem estar entre 1 e 3.")
			continue
		}

		if !j.posicaoVazia(linha, coluna) {
			fmt.Println("Posição ocupada. Escolha outra posição.")
			continue
		}

		// Grava a jogada (ajustando o índice de 1-3 para 0-2)
		j.tabuleiro[linha-1][coluna-1] = "X"
		break
	}
}

// 5. Gera a jogada aleatória do computador
func (j *Jogo) jogadaDoComputador() {
	fmt.Println("Vez do computador...")
	time.Sleep(1 * time.Second) // Pequena pausa para simular pensamento

	for {
		linha := rand.Intn(3) + 1
		coluna := rand.Intn(3) + 1

		if j.posicaoVazia(linha, coluna) {
			j.tabuleiro[linha-1][coluna-1] = "O"
			fmt.Printf("O computador jogou na posição: linha %d, coluna %d.\n", linha, coluna)
			break
		}
	}
}

// 6. Verifica se as coordenadas estão dentro dos limites da matriz
func (j *Jogo) posicaoValida(linha, coluna int) bool {
	return linha >= 1 && linha <= 3 && coluna >= 1 && coluna <= 3
}

// 7. Verifica se a posição escolhida contém um espaço vazio
func (j *Jogo) posicaoVazia(linha, coluna int) bool {
	return j.tabuleiro[linha-1][coluna-1] == " "
}

// 8. Verifica todas as condições de vitória
func (j *Jogo) verificarVitoria(simbolo string) bool {
	return j.verificarLinhas(simbolo) || j.verificarColunas(simbolo) || j.verificarDiagonais(simbolo)
}

func (j *Jogo) verificarLinhas(simbolo string) bool {
	for l := 0; l < 3; l++ {
		if j.tabuleiro[l][0] == simbolo && j.tabuleiro[l][1] == simbolo && j.tabuleiro[l][2] == simbolo {
			return true
		}
	}
	return false
}

func (j *Jogo) verificarColunas(simbolo string) bool {
	for c := 0; c < 3; c++ {
		if j.tabuleiro[0][c] == simbolo && j.tabuleiro[1][c] == simbolo && j.tabuleiro[2][c] == simbolo {
			return true
		}
	}
	return false
}

func (j *Jogo) verificarDiagonais(simbolo string) bool {
	// Diagonal principal
	if j.tabuleiro[0][0] == simbolo && j.tabuleiro[1][1] == simbolo && j.tabuleiro[2][2] == simbolo {
		return true
	}
	// Diagonal secundária
	if j.tabuleiro[0][2] == simbolo && j.tabuleiro[1][1] == simbolo && j.tabuleiro[2][0] == simbolo {
		return true
	}
	return false
}

// 9. Verifica se não existem mais espaços vazios no tabuleiro
func (j *Jogo) tabuleiroCheio() bool {
	for l := 0; l < 3; l++ {
		for c := 0; c < 3; c++ {
			if j.tabuleiro[l][c] == " " {
				return false
			}
		}
	}
	return true
}
