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
  auth_token  = var.argocd_token
}

