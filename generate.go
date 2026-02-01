//go:generate sqlc generate -f sqlc.yaml
//go:generate swag init -g ./pkg/app/server.go -o ./pkg/docs --parseDependency --parseInternal --useStructName

package app
