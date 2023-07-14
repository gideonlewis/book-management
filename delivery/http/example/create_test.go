package example_test

import (
	"net/http"
	"testing"

	"git.teqnological.asia/teq-go/teq-pkg/teq"
	"github.com/stretchr/testify/require"

	"git.teqnological.asia/teq-go/teq-echo/delivery/http/example"
	"git.teqnological.asia/teq-go/teq-echo/fixture/database"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/repository"
	"git.teqnological.asia/teq-go/teq-echo/usecase"
)

func TestCreate(t *testing.T) {
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := example.Route{UseCase: usecase.New(repo)}

	t.Run("200", func(t *testing.T) {
		t.Run("Create", func(t *testing.T) {
			rec, c := setUpTestCreate(payload.CreateExampleRequest{Name: teq.String("test")})

			require.NoError(t, r.Create(c))
			require.Equal(t, http.StatusOK, rec.Code)

			// remove data for the next test case
			db.TruncateTables()
		})
	})

	t.Run("400", func(t *testing.T) {
		t.Run("Invalid name", func(t *testing.T) {
			rec, c := setUpTestCreate(payload.CreateExampleRequest{Name: teq.String("  ")})

			require.NoError(t, r.Create(c))
			require.Equal(t, http.StatusBadRequest, rec.Code)

			// remove data for the next test case
			db.TruncateTables()
		})
	})
}
