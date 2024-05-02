package dependencies

import (
	"gitlab.praktikum-services.ru/Stasyan/momo-store/internal/store/dumplings"
	"gitlab.praktikum-services.ru/Stasyan/momo-store/internal/store/dumplings/fake"
)

// NewFakeDumplingsStore returns new fake store for app
func NewFakeDumplingsStore() (dumplings.Store, error) {
	packs := []dumplings.Product{
		{
			ID:          1,
			Name:        "Пельмени",
			Description: "С говядиной",
			Price:       5.00,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/32cc88a33c3243a6a8838c034878c564.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121614Z&X-Amz-Expires=1296000&X-Amz-Signature=13C599F11B76A2E359EC0514CC9839CB0383B98E2F2E71BE300633ED7D8EAA44&X-Amz-SignedHeaders=host",
		},
		{
			ID:          2,
			Name:        "Хинкали",
			Description: "Со свининой",
			Price:       3.50,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/4bdaeab0ee1842dc888d87d4a435afdd.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121802Z&X-Amz-Expires=1296000&X-Amz-Signature=D24A4FE9B57F3A843B39C6FAFC7E24B613A3A2D0DDFBF0E46176F1CD1C1BC59F&X-Amz-SignedHeaders=host",
		},
		{
			ID:          3,
			Name:        "Манты",
			Description: "С мясом молодых бычков",
			Price:       2.75,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/50b583271fa0409fb3d8ffc5872e99bb.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121821Z&X-Amz-Expires=1296000&X-Amz-Signature=418C45CA1A123B9093F7B9262239E6C4665727C74EE45E072509E84D462B938D&X-Amz-SignedHeaders=host",
		},
		{
			ID:          4,
			Name:        "Буузы",
			Description: "С телятиной и луком",
			Price:       4.00,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/7685ad7e9e634a58a4c29120ac5a5ee1.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121839Z&X-Amz-Expires=1296000&X-Amz-Signature=D81F6FCD28EDB194171CE443539ABAC09768E9984C021DA73910AE6C0E3344EE&X-Amz-SignedHeaders=host",
		},
		{
			ID:          5,
			Name:        "Цзяоцзы",
			Description: "С говядиной и свининой",
			Price:       7.25,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/788c073d83c14b3fa00675306dfb32b5.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121900Z&X-Amz-Expires=1296000&X-Amz-Signature=82A59C35C4E9D989321C48A19BC609659C036A977AE004790B86EC3810BDB51E&X-Amz-SignedHeaders=host",
		},
		{
			ID:          6,
			Name:        "Гедза",
			Description: "С соевым мясом",
			Price:       3.50,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/788c073d83c14b3fa00675306dfb32b5.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121900Z&X-Amz-Expires=1296000&X-Amz-Signature=82A59C35C4E9D989321C48A19BC609659C036A977AE004790B86EC3810BDB51E&X-Amz-SignedHeaders=host",
		},
		{
			ID:          7,
			Name:        "Дим-самы",
			Description: "С уткой",
			Price:       2.65,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/8dee5a92281746aa887d6f19cf9fdcc7%20%281%29.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121955Z&X-Amz-Expires=1382400&X-Amz-Signature=8C32FF21E0285716B52C26B12EFA640A1E7EBF1AD9500FDEF0D5B3F34B9A29C0&X-Amz-SignedHeaders=host",
		},
		{
			ID:          8,
			Name:        "Момо",
			Description: "С бараниной",
			Price:       5.00,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/f64dcea998e34278a0006e0a2b104710.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T122017Z&X-Amz-Expires=1382400&X-Amz-Signature=837076AA1A8CB99B75DDF12D217EE3EE4C1CC576E5D5FC01A407EBA7F006A73E&X-Amz-SignedHeaders=host",
		},
		{
			ID:          9,
			Name:        "Вонтоны",
			Description: "С креветками",
			Price:       4.10,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/32cc88a33c3243a6a8838c034878c564.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121614Z&X-Amz-Expires=1296000&X-Amz-Signature=13C599F11B76A2E359EC0514CC9839CB0383B98E2F2E71BE300633ED7D8EAA44&X-Amz-SignedHeaders=host",
		},
		{
			ID:          10,
			Name:        "Баоцзы",
			Description: "С капустой",
			Price:       4.20,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/4bdaeab0ee1842dc888d87d4a435afdd.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121802Z&X-Amz-Expires=1296000&X-Amz-Signature=D24A4FE9B57F3A843B39C6FAFC7E24B613A3A2D0DDFBF0E46176F1CD1C1BC59F&X-Amz-SignedHeaders=host",
		},
		{
			ID:          11,
			Name:        "Кундюмы",
			Description: "С грибами",
			Price:       5.45,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/50b583271fa0409fb3d8ffc5872e99bb.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121821Z&X-Amz-Expires=1296000&X-Amz-Signature=418C45CA1A123B9093F7B9262239E6C4665727C74EE45E072509E84D462B938D&X-Amz-SignedHeaders=host",
		},
		{
			ID:          12,
			Name:        "Курзе",
			Description: "С крабом",
			Price:       3.25,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/7685ad7e9e634a58a4c29120ac5a5ee1.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121839Z&X-Amz-Expires=1296000&X-Amz-Signature=D81F6FCD28EDB194171CE443539ABAC09768E9984C021DA73910AE6C0E3344EE&X-Amz-SignedHeaders=host",
		},
		{
			ID:          13,
			Name:        "Бораки",
			Description: "С говядиной и бараниной",
			Price:       4.00,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/788c073d83c14b3fa00675306dfb32b5.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121900Z&X-Amz-Expires=1296000&X-Amz-Signature=82A59C35C4E9D989321C48A19BC609659C036A977AE004790B86EC3810BDB51E&X-Amz-SignedHeaders=host",
		},
		{
			ID:          14,
			Name:        "Равиоли",
			Description: "С рикоттой",
			Price:       2.90,
			Image:       "https://storage.yandexcloud.net/dumpling-s3-bucket/8dee5a92281746aa887d6f19cf9fdcc7%20%281%29.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=YCAJErJIHCBcoMxi-JK9k1dpR%2F20240502%2Fru-central1%2Fs3%2Faws4_request&X-Amz-Date=20240502T121955Z&X-Amz-Expires=1382400&X-Amz-Signature=8C32FF21E0285716B52C26B12EFA640A1E7EBF1AD9500FDEF0D5B3F34B9A29C0&X-Amz-SignedHeaders=host",
		},
	}

	store := fake.NewStore()
	store.SetAvailablePacks(packs...)

	return store, nil
}
