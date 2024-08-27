# CleanArch Challenger - FullCycle

## Passo a Passo para Rodar a Aplicação

1. **Inicie o Banco de Dados e RabbitMQ**
   - Execute o comando abaixo para levantar o ambiente:
     ```bash
     docker compose up -d
     ```

2. **Execute as Migrations Necessárias**
   - Utilize o comando a seguir para rodar as migrations:
     ```bash
     make migrate
     ```
   - No arquivo `Makefile`, há alguns comandos de migrations já criados, e na pasta `sql` você encontrará os comandos `up` e `down`.

3. **Rode a Aplicação**
   - Dentro do diretório `cmd/ordersystem`, execute o comando:
     ```bash
     go run main.go wire_gen.go
     ```

## Portas das Aplicações

- **WEB**: `8000`
- **gRPC**: `50051`
- **GraphQL**: `8080`
