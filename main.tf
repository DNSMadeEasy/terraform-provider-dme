provider "dme" {
  api_key    = "1"
  secret_key = "2"
  sandbox = true
}

# resource "dme_template_record" "record" {
#   template_id = "25095"
#   name = "darsh14"
#   type = "A"
#   value = "1.2.3.6"
#   ttl = "86405"
# }

data "dme_template_record" "first" {
  template_id = "25095"
  name = "darsh"
  type = "A"
}

output "demo" {
  value = "${data.dme_template_record.first}"
}

