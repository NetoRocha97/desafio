# Desafio Snet

## Este projeto faz parte da inserção de um bolsista na empresa ServiceNet, no qual ele aprende qual o padrão da empresa, além de aprender algumas funcionalidades da linguagem Golang.

Para a aplicação ser rodada, tem duas formas, através da build e do compose, abaixo demosntrarei como deve ser feito.

Antes de tudo, atualize as libs, para que o projeto possa ser rodade perfeitamente.

Pode ser usado o comando **go mod download**, para instalar todos os módulos.

Instale o docker em sua máquina, e para isso, basta acessar **https://www.docker.com/products/docker-desktop/** e escolher qual SO sua máquina utiliza.

Caso o docker já esteja instalado, baixe a imagem do programa e posteriormente coloque-a para rodar

O próximo passo é acessar o banco de dados, e para acessá-lo, deve-se acessar o localhost pela porta 5432 e utilizar o user: postgres e password: postgres

Para executar o programa:

Ao usar o **go build**, você pode gerar um binário executável:

**go build**

Após criado o executável, execute-o para garantir que será compilado com sucesso

Já no docker, para fazer a build da imagem utiliza o comando:

**docker-compose build**

E para rodá-lo, usa-se:

**docker-compose up**



