package lm_testing

import (
	"encoding/json"
	"fmt"
	"io"
	"jrobic/lawn-mower/catalog-service/domain"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubCatalogRepository struct {
	Mowers map[string]*domain.Mower
}

func (r *StubCatalogRepository) Find(id string) (*domain.Mower, error) {
	return r.Mowers[id], nil
}

func (r *StubCatalogRepository) Add(input domain.AddMowerDTO) (*domain.Mower, error) {
	id := fmt.Sprint(len(r.Mowers) + 1)

	mower := &domain.Mower{
		Id:   id,
		Name: input.Name,
	}

	r.Mowers[id] = mower

	return mower, nil
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func AssertMowerEquals(t testing.TB, got, want domain.Mower) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func GetMowerFromResponse(t testing.TB, body io.Reader) (mower domain.Mower) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&mower)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into Mower, '%v'", body, err)
	}

	return
}
