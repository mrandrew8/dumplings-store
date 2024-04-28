resource "yandex_dns_zone" "dumpling-zone1" {
  name        = "dumpling-zone1"
  description = "Test public zone"

  labels = {
    label1 = "test-public"
  }

  zone    = "dumlping-store24.ru."
  public  = true
}

resource "yandex_dns_recordset" "rs1" {
  zone_id = yandex_dns_zone.dumpling-zone1.id
  name    = "dumlping-store24.ru."
  type    = "A"
  ttl     = 200
  data    = [var.data]
}