package resources

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/hermanschaaf/cq-source-xkcd/client"
	"github.com/hermanschaaf/cq-source-xkcd/internal/xkcd"
	"github.com/sunil494/cq-source-test/internal/xkcd"
)

func TestListingsTable(t *testing.T) {
	var comic xkcd.PriceLabs
	if err := faker.FakeObject(&comic); err != nil {
		t.Fatal(err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		d, _ := json.Marshal(comic)
		_, _ = w.Write(d)
	}))
	defer ts.Close()

	client.TestHelper(t, ListingsTable(), ts)
}
