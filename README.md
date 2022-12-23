# Example GraphQL server with Go

## Run server:

```bash
git clone git@github.com:will-t-harris/graphql-go.git
cd graphql-go
go run .
```

## Example query:

```bash
curl -X POST -H "Content-Type: application/json" --data '{ "query": "{ beastList {id name } }" }' http://localhost:8080/graphql
```
