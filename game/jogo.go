package game

import "fmt"

const (
    Empty = " "
    X     = "X"
    O     = "O"
    Size  = 3
)

// Jogo mantém o estado do tabuleiro
type Jogo struct {
    Tabuleiro [Size][Size]string
}

// Novo cria um novo jogo e inicializa o tabuleiro
func Novo() *Jogo {
    j := &Jogo{}
    j.InicializarTabuleiro()
    return j
}

// InicializarTabuleiro preenche o tabuleiro com espaços vazios
func (j *Jogo) InicializarTabuleiro() {
    for l := 0; l < Size; l++ {
        for c := 0; c < Size; c++ {
            j.Tabuleiro[l][c] = Empty
        }
    }
}

// PosicaoValida verifica se as coordenadas estão dentro do tabuleiro (1..Size)
func (j *Jogo) PosicaoValida(linha, coluna int) bool {
    return linha >= 1 && linha <= Size && coluna >= 1 && coluna <= Size
}

// PosicaoVazia verifica se a posição (1-based) está vazia
func (j *Jogo) PosicaoVazia(linha, coluna int) bool {
    return j.Tabuleiro[linha-1][coluna-1] == Empty
}

// FazerJogada tenta marcar a posição com o símbolo fornecido, retornando erro em caso de inválido/ocupado
func (j *Jogo) FazerJogada(linha, coluna int, simbolo string) error {
    if !j.PosicaoValida(linha, coluna) {
        return fmt.Errorf("posição inválida: %d,%d", linha, coluna)
    }
    if !j.PosicaoVazia(linha, coluna) {
        return fmt.Errorf("posição ocupada: %d,%d", linha, coluna)
    }
    j.Tabuleiro[linha-1][coluna-1] = simbolo
    return nil
}

// VerificarVitoria combina as checagens de linhas, colunas e diagonais
func (j *Jogo) VerificarVitoria(simbolo string) bool {
    return j.verificarLinhas(simbolo) || j.verificarColunas(simbolo) || j.verificarDiagonais(simbolo)
}

func (j *Jogo) verificarLinhas(simbolo string) bool {
    for l := 0; l < Size; l++ {
        if j.Tabuleiro[l][0] == simbolo && j.Tabuleiro[l][1] == simbolo && j.Tabuleiro[l][2] == simbolo {
            return true
        }
    }
    return false
}

func (j *Jogo) verificarColunas(simbolo string) bool {
    for c := 0; c < Size; c++ {
        if j.Tabuleiro[0][c] == simbolo && j.Tabuleiro[1][c] == simbolo && j.Tabuleiro[2][c] == simbolo {
            return true
        }
    }
    return false
}

func (j *Jogo) verificarDiagonais(simbolo string) bool {
    // Diagonal principal
    if j.Tabuleiro[0][0] == simbolo && j.Tabuleiro[1][1] == simbolo && j.Tabuleiro[2][2] == simbolo {
        return true
    }
    // Diagonal secundária
    if j.Tabuleiro[0][2] == simbolo && j.Tabuleiro[1][1] == simbolo && j.Tabuleiro[2][0] == simbolo {
        return true
    }
    return false
}

// TabuleiroCheio retorna true se não houver mais posições vazias
func (j *Jogo) TabuleiroCheio() bool {
    for l := 0; l < Size; l++ {
        for c := 0; c < Size; c++ {
            if j.Tabuleiro[l][c] == Empty {
                return false
            }
        }
    }
    return true
}

