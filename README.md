<h1 align="center"> Video Streaming Project </h1>

<p align="center">
<img src="http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=GREEN&style=for-the-badge"/>
</p>

## Descrição do Projeto

Streaming não é um conceito novo, mas vem crescendo e se tornando cada vez mais presente na vida das pessoas.Foi pensando nisso que decidi criar o meu próprio serviço de streaming, para estudar e entender as etapas desde a solicitação feita pelo cliente, processamento dos dados e entrega para o usuário final, e colocando em prática uma arquitetura que fosse possível escalar e manter uma alta performance.

## ✔️ Técnicas e tecnologias utilizadas

- ``Kotlin``
- ``Java``
- ``Golang``
- ``Golang``
- ``Apache Kafka``
- ``Elasticsearch``
- ``MySql``
- ``WebSocket``

## Arquitetura
![system-design-v2(2) drawio](https://user-images.githubusercontent.com/43520784/224572318-e25ac206-ee9d-4e9e-8064-5e802c57028d.png)

Event-driven: Arquitetura orientada a eventos, a comunicação entre os componentes(microserviços) se dá através do envio de eventos(publishers) para um barramento de eventos(Broker), onde os componentes inscritos(consumers) consomem esses eventos.

## Componentes: 
- stream-writter: escrita de arquivos
- stream-reader: leitura dos arquivos
- Logcenter: centralizados de logs
- admin-manager: gerenciamento dos arquivos(escrita, leitura, alteração e remoção dos videos)
- content-searcher: busca de vídeos no indexador
- delivery-websocket-server: mantem conexão com o cliente, consome os dados de vídeos e repassa para o cliente.
