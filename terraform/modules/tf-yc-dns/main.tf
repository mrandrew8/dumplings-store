resource "yandex_dns_zone" "dumpling-zone1" {
  name        = "dumpling-zone1"
  description = "Test public zone"

  labels = {
    label1 = "test-public"
  }

  zone    = "dumlping-store24.ru."
  public  = true
}
