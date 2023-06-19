package resources

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/sunil494/cq-source-pricelabs/client"
	"github.com/sunil494/cq-source-pricelabs/internal/pricelabs"
)

func TestListingsTable(t *testing.T) {
	var comic pricelabs.PriceLabs
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
