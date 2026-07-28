package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/EGS-software/tic-tac-tue-golang.git/ai"
	"github.com/EGS-software/tic-tac-tue-golang.git/game"
	"github.com/EGS-software/tic-tac-tue-golang.git/ui"
)

func main() {
	// Inicializa a semente para geração de números aleatórios
	rand.Seed(time.Now().UnixNano())

	jogo := game.Novo()

	fmt.Println("=== JOGO DA VELHA ===")

	vezDoJogador := ai.SortearQuemComeca()

	if vezDoJogador {
		fmt.Println("O jogador X começará a partida.")
	} else {
		fmt.Println("O computador O começará a partida.")
	}

	ui.ExibirTabuleiro(jogo)

	// Loop principal da partida
	for {
		if vezDoJogador {
			ui.JogadaDoJogador(jogo)
			ui.ExibirTabuleiro(jogo)
			if jogo.VerificarVitoria(game.X) {
				fmt.Println("O jogador X venceu a partida.")
				break
			}
		} else {
			linha, coluna := ai.JogadaAleatoria(jogo)
			ui.JogadaComputadorUI(jogo, linha, coluna)
			ui.ExibirTabuleiro(jogo)
			if jogo.VerificarVitoria(game.O) {
				fmt.Println("O computador O venceu a partida.")
				break
			}
		}

		if jogo.TabuleiroCheio() {
			fmt.Println("A partida terminou empatada.")
			break
		}

		// Alterna a vez
		vezDoJogador = !vezDoJogador
	}
}

