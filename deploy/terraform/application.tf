resource "argocd_application" "bitgo" {
  metadata {
    name      = "bitgo"
    namespace = "argocd"
  }

  wait = true

  spec {
    source {
      repo_url        = "https://github.com/jgrancell/bitgo"
      path            = "deploy/chart"
      target_revision = var.version

      helm {
        parameter {
          name  = "application.version"
          value = var.version
        }
      }
    }
  }

  destination {
    server    = "https://kubernetes.default.svc"
    namespace = "app-gc-websites"
  }

  sync_policy {
    automated = {
      prune     = true
      self_heal = true
    }
  }
}