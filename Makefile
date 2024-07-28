.PHONY: gql_generate
gql_generate:
	go run github.com/99designs/gqlgen generate

.PHONY: run
run:
	go run main.go