# DNSMadeEasy Provider


- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">


Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) Latest Version

- [Go](https://golang.org/doc/install) go1.13.8

## Building The Provider ##
Clone this repository to: `$GOPATH/src/github.com/DNSMadeEasy/terraform-provider-dme`.

```sh
$ mkdir -p $GOPATH/src/github.com/DNSMadeEasy; cd $GOPATH/src/github.com/DNSMadeEasy
$ git clone https://github.com/DNSMadeEasy/terraform-provider-dme.git
```

Enter the provider directory and run make build to build the provider binary.

```sh
$ cd $GOPATH/src/github.com/DNSMadeEasy/terraform-provider-dme
$ make build

```

Using The Provider
------------------
If you are building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory, run `terraform init` to initialize it.

ex.
```hcl
#configure provider with your DNSMadeEasy credentials.
provider "dme" {
  # DNSMadeEasy Api key
  apikey = "apikey"
  # DNSMadeEasy secret key
  secretkey = "secretkey"
  insecure = true
  proxy_url = "https://proxy_server:proxy_port"
}

resource "dme_domain" "example" {
  name            = "example.com"
  gtd_enabled     = "false"
  soa_id          = "${dme_custom_soa_record.example.id}"
  template_id     = "${dme_template.example.id}"
  vanity_id       = "${dme_vanity_nameserver_record.example.id}"
  transfer_acl_id = "${dme_transfer_acl.example.id}"
  folder_id       = "${dme_folder.example.id}"
}

```


```
terraform plan -parallelism=1
terraform apply -parallelism=1
```  

Developing The Provider
-----------------------
If you want to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine. You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider with sanity checks present in scripts directory and put the provider binary in `$GOPATH/bin` directory.

