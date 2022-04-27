//go:build pact
// +build pact

package pact

import (
	"assignment/pact/mockRepository"
	"assignment/server"
	"assignment/service"
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
)

func TestProvider(t *testing.T) {
	port, _ := utils.GetFreePort()
	newServer := server.NewServer()

	go newServer.StartServer(port)

	pact := dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "todo-backend",
		Consumer:                 "todo-frontend",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", port),
		BrokerURL:                  "https://sak-assignment.pactflow.io",
		ProviderVersion:            "1.0.0",
		ConsumerVersionSelectors:   []types.ConsumerVersionSelector{types.ConsumerVersionSelector{Latest: true}},
		BrokerToken:                "oN-yv0nYlrYmxkcdpO-slQ",
		PublishVerificationResults: true,
		StateHandlers: types.StateHandlers{
			"fetch items successfully": func() error {
				todoRepository := mockRepository.NewMockTodoRepository()
				service.NewTodoService(todoRepository)
				return nil
			},
			"create todo list item": func() error {
				return nil
			},
		},
	}

	verifyResponses, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(len(verifyResponses), "pact tests run")
}
