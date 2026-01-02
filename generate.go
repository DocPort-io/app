//go:generate sqlc generate -f sqlc.yaml
//go:generate swag init -g ./pkg/app/routes.go -o ./pkg/docs --parseDependency --parseInternal --useStructName

package app
