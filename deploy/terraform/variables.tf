variable "application_version"  {
  description = "The git tag or reference to deploy through Helm"
  type        = string
}

variable "argocd_server" {
  description = "The server that the ArgoCD provider will connect to."
  type        = string
  sensitive   = true
}

variable "argocd_token" {
  description = "The token that the ArgoCD provider will authenticate with."
  type        = string
  sensitive   = true
}