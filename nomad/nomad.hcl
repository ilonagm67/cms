job "cms" {
  vault {}
  namespace   = "dev"
  datacenters = ["dc1"]

  group "cms" {
    network {
      mode = "bridge"
    }

    count = 1
    task "cms" {
      shutdown_delay = "5s"
      driver = "docker"
      config {
        image = "vangdevops/cms:v0.3.3-beta"
        args = ["-memory","-verbose"]
      }
      template {
        data = <<EOH
{{with secret "kv/data/dev/cms/token"}}
TOKEN={{ .Data.data.token }}
{{end}}
{{with secret "kv/data/dev/cms/admin"}}
ADMIN={{ .Data.data.id }}
{{end}}
EOH

        destination = "${NOMAD_SECRETS_DIR}/env"
        env         = true
      }
    }
  }
}
