<p align="center">
  <img width="460" height="300" src="https://www.mundodocker.com.br/wp-content/uploads/2015/06/docker_facebook_share.png">
</p>

# Criação e Execução da imagem
#### GoLang
> Para executar o teste execute os comandos abaixo

1. Criar a imagem ```docker build -t luizfilipy/app .```
2. Executar imagem ```docker run -p 3001:3000 -d luizfilipy/app```

# Comandos úteis
* Ver os processos em execução ```docker ps```
* Parar a execução de uma imagem ```docker stop {CONTAINER ID}```
* Apagar todas as imagens ```docker system prune -a```

[ :whale: Veja a lista completa :whale:](https://docs.docker.com/engine/reference/commandline/docker/)