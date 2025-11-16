.PHONY: swagger-format swagger

swagger:
	swag init -g ./pkg/app/routes.go -o ./pkg/docs --parseDependency --parseInternal --useStructName

swagger-format:
	swag fmt
