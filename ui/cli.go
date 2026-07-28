package ui

import (
    "fmt"
    "time"

    "github.com/EGS-software/tic-tac-tue-golang.git/game"
)

// ExibirTabuleiro mostra o tabuleiro formatado na saída padrão
func ExibirTabuleiro(j *game.Jogo) {
    fmt.Println("\n         Colunas")
    fmt.Println("        1       2      3")
    fmt.Println("     +---+---+---+")
    for l := 0; l < game.Size; l++ {
        fmt.Printf("  %d  |", l+1)
        for c := 0; c < game.Size; c++ {
            fmt.Printf(" %s |", j.Tabuleiro[l][c])
        }
        fmt.Println("\n     +---+---+---+")
    }
    fmt.Println()
}

// JogadaDoJogador faz a leitura da jogada via stdin e tenta aplicá-la ao jogo
func JogadaDoJogador(j *game.Jogo) {
    var linha, coluna int

    for {
        fmt.Print("Informe a linha: ")
        _, errL := fmt.Scan(&linha)
        fmt.Print("Informe a coluna: ")
        _, errC := fmt.Scan(&coluna)

        if errL != nil || errC != nil {
            fmt.Println("Entrada inválida. Digite apenas números.")
            var lixo string
            _, _ = fmt.Scanln(&lixo)
            continue
        }

        if !j.PosicaoValida(linha, coluna) {
            fmt.Println("Posição inválida. A linha e a coluna devem estar entre 1 e 3.")
            continue
        }

        if !j.PosicaoVazia(linha, coluna) {
            fmt.Println("Posição ocupada. Escolha outra posição.")
            continue
        }

        // Aplicar jogada
        _ = j.FazerJogada(linha, coluna, game.X)
        break
    }
}

// JogadaComputadorUI aplica o movimento do computador e mostra mensagens/pausa para o usuário
func JogadaComputadorUI(j *game.Jogo, linha, coluna int) {
    fmt.Println("Vez do computador...")
    time.Sleep(1 * time.Second)
    _ = j.FazerJogada(linha, coluna, game.O)
    fmt.Printf("O computador jogou na posição: linha %d, coluna %d.\n", linha, coluna)
}

