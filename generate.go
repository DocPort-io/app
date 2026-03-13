//go:generate go tool oapi-codegen -config oapi-codegen.yml ./openapi-spec.yaml
//go:generate go tool sqlc generate -f sqlc.yaml
//go:generate go tool swag init -g ./pkg/app/server.go -o ./pkg/docs --parseDependency --parseInternal --useStructName

package app
