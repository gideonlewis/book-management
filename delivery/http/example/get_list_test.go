package example_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"git.teqnological.asia/teq-go/teq-echo/delivery/http/example"
	"git.teqnological.asia/teq-go/teq-echo/fixture/database"
	"git.teqnological.asia/teq-go/teq-echo/payload"
	"git.teqnological.asia/teq-go/teq-echo/repository"
	"git.teqnological.asia/teq-go/teq-echo/usecase"
)

func TestGetList(t *testing.T) {
	db := database.InitDatabase()
	defer db.TruncateTables()

	repo := repository.New(db.GetClient)
	r := example.Route{UseCase: usecase.New(repo)}

	t.Run("200", func(t *testing.T) {
		t.Run("Get list", func(t *testing.T) {
			rec, c := setUpTestGetList(payload.GetListExampleRequest{})

			require.NoError(t, r.GetList(c))
			require.Equal(t, http.StatusOK, rec.Code)

			// remove data for the next test case
			db.TruncateTables()
		})
	})
}
