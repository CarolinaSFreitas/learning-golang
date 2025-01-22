# Fundamentos I - Golang

1. ``go env``: Encontrar o GOPATH
O comando ``go env`` exibe informações sobre o ambiente Go configurado na máquina, incluindo variáveis importantes como GOPATH e GOROOT.

2. Pacotes em Go
Um pacote é um grupo de arquivos que trabalham juntos, e eles devem estar no mesmo diretório.
Todos os arquivos dentro do diretório fazem parte do mesmo pacote e são compilados juntos.
O pacote principal para a criação de executáveis em Go é sempre o package main.

3. Estrutura de um Executável
Para criar um executável em Go, você precisa:
- Definir o pacote principal: package main.
- Importar bibliotecas necessárias: Usar import para pacotes internos ou externos.
- Criar a função principal: A função func main() é obrigatória como ponto de entrada.

4. ``go mod init``: Criar um Módulo
O comando ``go mod init <modulo>`` inicializa um projeto Go como módulo.
Cria o arquivo go.mod, que é similar ao package.json em projetos Node.js. Ele gerencia as dependências e configurações do projeto.

5. ``go build``: Compilar o Projeto
O comando ``go build`` compila o código em um único arquivo binário executável.
Ele compila todos os arquivos do diretório atual.

6. Exportação de Variáveis e Funções
Em Go, a visibilidade de variáveis, funções e tipos é controlada pela capitalização da primeira letra do identificador:

- Letra inicial maiúscula: O identificador é exportado (público) e pode ser acessado por outros pacotes.
- Letra inicial minúscula: O identificador é não exportado (privado) e só pode ser usado dentro do pacote onde foi declarado.

7. Módulos - atualização

O motivo de o módulo não atualizar automaticamente no Go está relacionado à necessidade de recompilar o código após mudanças. O Go gera um binário estático com o estado atual do código-fonte no momento da execução do comando ``go build``.

**Se você alterar qualquer parte do código (em um arquivo principal ou em pacotes auxiliares), precisará rodar o comando go build novamente para que essas alterações sejam refletidas no binário gerado.**

- ``go install``: O comando go install é usado para compilar e instalar um programa ou pacote Go no seu ambiente, de forma que ele fique facilmente acessível de qualquer lugar no sistema. Movendo assim o binário gerado para o diretório especificado no GOPATH/bin ou no diretório configurado na variável GOBIN.

8. Instalação de Pacotes: ``go get <link do pacote>``
    - ex:  ``go get github.com/badoux/checkmail``

9. Chamar importação

Quando você importa um pacote em Go, você usa o que vem por último no caminho do pacote, ou seja, o nome que aparece após a última barra (/). Por exemplo:

Se você importar o pacote "github.com/badoux/checkmail", o nome do pacote que você vai usar é checkmail.
Portanto, para chamar a função ValidateFormat desse pacote, você utiliza o nome do pacote seguido de um ponto, assim:

``checkmail.ValidateFormat("carolina@hotmail.com")``

Importação: ``import "github.com/badoux/checkmail"``
Uso: ``checkmail.ValidateFormat("carolina@hotmail.com")``

10. Limpar Dependências Não Usadas
O comando ``go mod tidy`` é usado para limpar e organizar as dependências do projeto Go. 
Ele remove as dependências não utilizadas no código e adiciona as faltantes, garantindo que os arquivos go.mod e go.sum estejam atualizados e consistentes.

11. No Go, o código **não compila** se você declarar uma variável e não usá-la. O Go exige que todas as variáveis declaradas sejam utilizadas no código.

12. Tipos de Números Inteiros: 
- Existem int8, int16, int32, int64, com limites de valores dependendo do número de bits. int varia conforme a arquitetura do computador.

13. Números Inteiros Não-Sinalizados: 
``uint`` (unsigned int) não aceita valores negativos.

14. Aliases:
- ``byte`` é um alias para uint8.
- ``rune`` é um alias para int32.

15. Números Reais: 
float32 e float64. O tipo ``float`` é inferido quando não declarado.

16. Strings: 
- ``char`` não existe em Go; caracteres são tratados como int32 (equivalente na tabela ASCII). 
- Strings podem ser atribuídas diretamente ou com declaração implícita.

17. Valor Zero de Tipos: 
- Strings têm valor zero como string vazia, números inteiros têm zero como valor, e booleanos têm false como valor zero.

18. Erros: 
O tipo ``error`` é usado para representar erros. Para criar um erro, usa-se a função ``errors.New()``.

