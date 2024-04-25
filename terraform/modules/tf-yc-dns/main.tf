resource "yandex_dns_zone" "dumpling-zone" {
  name        = "dumpling-zone"
  description = "Test public zone"
  zone    = "student12.dumpling-store.com."
  public  = true
}

resource "yandex_dns_recordset" "rs1" {
  zone_id = yandex_dns_zone.dumpling-zone.id
  name    = "student12.dumpling-store.com."
  type    = "A"
  ttl     = 200
  data    = ["10.1.0.1"]
}