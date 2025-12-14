//go:generate sqlc generate -f sqlc.yaml
//go:generate swag init -g ./pkg/app/routes.go -o ./pkg/docs --parseDependency --parseInternal --useStructName

package app

import (
	// Import for text processing, needed by Swagger file generation
	_ "golang.org/x/text"
)
