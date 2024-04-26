resource "yandex_dns_zone" "dumpling-zone1" {
  name        = "my-public-zone"
  description = "Test public zone"

  labels = {
    label1 = "test-public"
  }

  zone    = "test.example-public2.com."
  public  = true
}

resource "yandex_dns_recordset" "rs1" {
  zone_id = yandex_dns_zone.dumpling-zone1.id
  name    = "test.example-public.com."
  type    = "A"
}