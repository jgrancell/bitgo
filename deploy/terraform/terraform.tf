terraform {
  required_providers {
    argocd = {
      source = "oboukili/argocd"
      version = "2.1.0"
    }
  }
}

provider "argocd" {
  server_addr = var.argocd_server
  username    = var.argocd_username
  password    = var.argocd_password
}

