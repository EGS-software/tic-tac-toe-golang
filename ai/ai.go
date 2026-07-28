package ai

import (
    "math/rand"

    "github.com/EGS-software/tic-tac-tue-golang.git/game"
)

// SortearQuemComeca retorna true se o jogador humano (X) começar; false para computador
func SortearQuemComeca() bool {
    return rand.Intn(2) == 0
}

// JogadaAleatoria escolhe uma posição vazia aleatória no tabuleiro
func JogadaAleatoria(j *game.Jogo) (linha, coluna int) {
    for {
        linha = rand.Intn(game.Size) + 1
        coluna = rand.Intn(game.Size) + 1
        if j.PosicaoVazia(linha, coluna) {
            return linha, coluna
        }
    }
}

