package tests

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/Askaell/homework/pkg/server"
)

const port = "80"
const addres = "http://127.0.0.1/test"

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func Test_serverShutdown(t *testing.T) {
	server := new(server.Server)

	go func() {
		if err := server.Run(port, http.HandlerFunc(testHandler)); err != nil && err != http.ErrServerClosed {
			t.Errorf("error occured while running http server: %s", err)
		}
	}()

	for {
		resp, err := http.Get(addres)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if err := server.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
			t.Errorf("error occured on server shutting down: %s", err)
		}
		break
	}

}
