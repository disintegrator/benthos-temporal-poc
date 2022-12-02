package main

import (
	"context"

	"github.com/benthosdev/benthos/v4/public/service"

	_ "money-transfer-project-template-go/app/benthos/temporal"

	_ "github.com/benthosdev/benthos/v4/public/components/all"
)

func main() {
	service.RunCLI(context.Background())
}
