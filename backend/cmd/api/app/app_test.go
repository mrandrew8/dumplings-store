package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.praktikum-services.ru/Stasyan/momo-store/cmd/api/dependencies"
)

func TestFakeAppIntegrational(t *testing.T) {
	store, err := dependencies.NewFakeDumplingsStore()
	assert.NoError(t, err)
	app, err := NewInstance(store)
	assert.NoError(t, err)

	t.Run("create_order", func(t *testing.T) {
		for i := 1; i <= 10; i++ {
			t.Run("id"+strconv.Itoa(i), func(t *testing.T) {
				r := httptest.NewRequest("POST", "/orders", nil)
				w := httptest.NewRecorder()
				app.CreateOrderController(w, r)

				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
				fmt.Fprintln(os.Stdout, "_____")
				fmt.Fprintln(os.Stdout, w.Body.String())
				fmt.Fprintln(os.Stdout, "_____")

				expectedJSON, err := json.Marshal(map[string]interface{}{"id": i})
				assert.NoError(t, err)
				assert.JSONEq(t, string(expectedJSON), w.Body.String())
			})
		}
	})

	t.Run("list_dumplings", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/packs", nil)
		w := httptest.NewRecorder()
		app.ListDumplingsController(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		fmt.Fprintln(os.Stdout, "_____")
		fmt.Fprintln(os.Stdout, w.Body.String())
		fmt.Fprintln(os.Stdout, "_____")

		expectedJSON := "{\"results\":[{\"id\":1,\"name\":\"Пельмени\",\"price\":5,\"description\":\"С говядиной\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/32cc88a33c3243a6a8838c034878c564.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121614Z&X-Amz-Expires=1296000&X-Amz-Signature=13C599F11B76A2E359EC0514CC9839CB0383B98E2F2E71BE300633ED7D8EAA44&X-Amz-SignedHeaders=host\"},{\"id\":2,\"name\":\"Хинкали\",\"price\":3.5,\"description\":\"Со свининой\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/4bdaeab0ee1842dc888d87d4a435afdd.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121802Z&X-Amz-Expires=1296000&X-Amz-Signature=D24A4FE9B57F3A843B39C6FAFC7E24B613A3A2D0DDFBF0E46176F1CD1C1BC59F&X-Amz-SignedHeaders=host\"},{\"id\":3,\"name\":\"Манты\",\"price\":2.75,\"description\":\"С мясом молодых бычков\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/50b583271fa0409fb3d8ffc5872e99bb.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121821Z&X-Amz-Expires=1296000&X-Amz-Signature=418C45CA1A123B9093F7B9262239E6C4665727C74EE45E072509E84D462B938D&X-Amz-SignedHeaders=host\"},{\"id\":4,\"name\":\"Буузы\",\"price\":4,\"description\":\"С телятиной и луком\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/7685ad7e9e634a58a4c29120ac5a5ee1.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121839Z&X-Amz-Expires=1296000&X-Amz-Signature=D81F6FCD28EDB194171CE443539ABAC09768E9984C021DA73910AE6C0E3344EE&X-Amz-SignedHeaders=host\"},{\"id\":5,\"name\":\"Цзяоцзы\",\"price\":7.25,\"description\":\"С говядиной и свининой\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/788c073d83c14b3fa00675306dfb32b5.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121900Z&X-Amz-Expires=1296000&X-Amz-Signature=82A59C35C4E9D989321C48A19BC609659C036A977AE004790B86EC3810BDB51E&X-Amz-SignedHeaders=host\"},{\"id\":6,\"name\":\"Гедза\",\"price\":3.5,\"description\":\"С соевым мясом\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/788c073d83c14b3fa00675306dfb32b5.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121900Z&X-Amz-Expires=1296000&X-Amz-Signature=82A59C35C4E9D989321C48A19BC609659C036A977AE004790B86EC3810BDB51E&X-Amz-SignedHeaders=host\"},{\"id\":7,\"name\":\"Дим-самы\",\"price\":2.65,\"description\":\"С уткой\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/8dee5a92281746aa887d6f19cf9fdcc7%20%281%29.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121955Z&X-Amz-Expires=1382400&X-Amz-Signature=8C32FF21E0285716B52C26B12EFA640A1E7EBF1AD9500FDEF0D5B3F34B9A29C0&X-Amz-SignedHeaders=host\"},{\"id\":8,\"name\":\"Момо\",\"price\":5,\"description\":\"С бараниной\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/f64dcea998e34278a0006e0a2b104710.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T122017Z&X-Amz-Expires=1382400&X-Amz-Signature=837076AA1A8CB99B75DDF12D217EE3EE4C1CC576E5D5FC01A407EBA7F006A73E&X-Amz-SignedHeaders=host\"},{\"id\":9,\"name\":\"Вонтоны\",\"price\":4.1,\"description\":\"С креветками\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/32cc88a33c3243a6a8838c034878c564.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121614Z&X-Amz-Expires=1296000&X-Amz-Signature=13C599F11B76A2E359EC0514CC9839CB0383B98E2F2E71BE300633ED7D8EAA44&X-Amz-SignedHeaders=host\"},{\"id\":10,\"name\":\"Баоцзы\",\"price\":4.2,\"description\":\"С капустой\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/4bdaeab0ee1842dc888d87d4a435afdd.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121802Z&X-Amz-Expires=1296000&X-Amz-Signature=D24A4FE9B57F3A843B39C6FAFC7E24B613A3A2D0DDFBF0E46176F1CD1C1BC59F&X-Amz-SignedHeaders=host\"},{\"id\":11,\"name\":\"Кундюмы\",\"price\":5.45,\"description\":\"С грибами\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/50b583271fa0409fb3d8ffc5872e99bb.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121821Z&X-Amz-Expires=1296000&X-Amz-Signature=418C45CA1A123B9093F7B9262239E6C4665727C74EE45E072509E84D462B938D&X-Amz-SignedHeaders=host\"},{\"id\":12,\"name\":\"Курзе\",\"price\":3.25,\"description\":\"С крабом\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/7685ad7e9e634a58a4c29120ac5a5ee1.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121839Z&X-Amz-Expires=1296000&X-Amz-Signature=D81F6FCD28EDB194171CE443539ABAC09768E9984C021DA73910AE6C0E3344EE&X-Amz-SignedHeaders=host\"},{\"id\":13,\"name\":\"Бораки\",\"price\":4,\"description\":\"С говядиной и бараниной\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/788c073d83c14b3fa00675306dfb32b5.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121900Z&X-Amz-Expires=1296000&X-Amz-Signature=82A59C35C4E9D989321C48A19BC609659C036A977AE004790B86EC3810BDB51E&X-Amz-SignedHeaders=host\"},{\"id\":14,\"name\":\"Равиоли\",\"price\":2.9,\"description\":\"С рикоттой\",\"image\":\"https://storage.yandexcloud.net/dumpling-s3-bucket/8dee5a92281746aa887d6f19cf9fdcc7%20%281%29.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121955Z&X-Amz-Expires=1382400&X-Amz-Signature=8C32FF21E0285716B52C26B12EFA640A1E7EBF1AD9500FDEF0D5B3F34B9A29C0&X-Amz-SignedHeaders=host\"}]}\n"

		assert.NoError(t, err)
		assert.JSONEq(t, string(expectedJSON), w.Body.String())
	})

	t.Run("healthcheck", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/health", nil)
		w := httptest.NewRecorder()
		app.HealthcheckController(w, r)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
