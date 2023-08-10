terraform {
  required_providers {
    mezmo = {
      source = "registry.terraform.io/mezmo/mezmo"
    }
  }
  required_version = ">= 1.1.0"
}

provider "mezmo" {
  auth_key = "my secret"
}

resource "mezmo_pipeline" "pipeline1" {
  title = "My pipeline"
}

resource "mezmo_http_source" "source1" {
  pipeline    = mezmo_pipeline.pipeline1.id
  title       = "My HTTP source"
  description = "This receives data from my webhook"
  decoding    = "json"
}
