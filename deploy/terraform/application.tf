resource "argocd_application" "bitgo" {
  metadata {
    name      = "bitgo"
    namespace = "argocd"
  }

  wait = true

  spec {
    project = "applications"

    source {
      repo_url        = "https://github.com/jgrancell/bitgo.git"
      path            = "deploy/chart"
      target_revision = var.application_version

      helm {
        parameter {
          name  = "application.version"
          value = var.application_version
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
}