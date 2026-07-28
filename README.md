Sumário
---------
- Descrição
- Estrutura do repositório
- Requisitos
- Como compilar e executar
- Como jogar
- Design / Modularização
- Possíveis melhorias
- Contribuindo
- Licença

Descrição
---------
Este repositório contém uma implementação simples do Jogo da Velha em Go. O código original foi refatorado para ficar modular:

- `game/`  : lógica do jogo e regras (estado do tabuleiro, validações, verificação de vitória).
- `ai/`    : decisões do computador (quem começa, escolha de jogada — atualmente aleatória).
- `ui/`    : interface de linha de comando (exibe tabuleiro, lê jogadas do jogador, anuncia jogadas do computador).
- `main.go`: orquestra o jogo usando os pacotes acima.

Estrutura do repositório
------------------------
Exemplo de arquivos principais:

- `go.mod`
- `LICENSE`
- `main.go`
- `game/jogo.go`
- `ai/ai.go`
- `ui/cli.go`
- `README.md`

Requisitos
----------
- Go 1.26.x (o `go.mod` indica `go 1.26.2`).

Como compilar e executar
------------------------
No diretório do projeto, você pode executar diretamente com `go run` ou compilar um binário:

```bash
cd /Users/jao/Documents/Projetos/tic-tac-tue-golang
go run .
```

ou

```bash
cd /Users/jao/Documents/Projetos/tic-tac-tue-golang
go build -o ttt .
./ttt
```

Como jogar
----------
O jogo é jogado no terminal.

- Ao começar, será informado quem inicia (jogador X ou computador O).
- Se for sua vez (X), o programa pedirá primeiro a linha (1-3) e depois a coluna (1-3).
- Digite números inteiros entre 1 e 3. Se a entrada for inválida ou a posição estiver ocupada, será solicitado novamente.

Design / Modularização
-----------------------
- O pacote `game` contém toda a lógica pura (sem prints), facilitando testes unitários.
- O pacote `ui` lida com interação com o usuário (entrada/saída) e pequenas esperas para melhorar a experiência.
- O pacote `ai` contém políticas para o computador; atualmente a jogada é aleatória.

Possíveis melhorias
-------------------
- Implementar IA inteligente (Minimax) em `ai/`.
- Cobertura de testes unitários para `game/` (detectar vitórias, empates, validações).
- Suporte a modos diferentes (PvP, PvC com níveis de dificuldade).
- Interfaces para permitir frontends alternativos (web, GUI) reaproveitando `game/`.


Contribuindo
------------
Pull requests são bem-vindos. Para mudanças de comportamento do jogo, prefira:

1. Adicionar testes em `game/` que descrevam o comportamento esperado.
2. Implementar a mudança em um branch separado e abrir um PR com descrição e exemplos.

Contribuidores
--------------
Pessoas que contribuíram/criaram este projeto:

- @jvbenetti
- @Mauro-Roncata

Licença
-------
O repositório contém um arquivo `LICENSE`. Verifique-o para detalhes sobre uso e contribuições.

Contato
-------
Se quiser que eu adicione testes automáticos ou implemente IA (minimax), responda aqui e eu posso implementar como próximo passo.

