package nginx

var nginx_template string = `
{{ $root := . }}
{{ $pod := getPod $root.Component }}
events {
    worker_connections  1024;  ## Default: 1024
}
stream { {{ range $c := $pod.Spec.Containers }}{{ range $port := $c.Ports }}
    {{ $upstream := (print $c.Name "-" $port.ContainerPort "-upstream")  }}
    upstream {{ $upstream }} { {{ range $i, $e := (until $root.Component.Replicas) }}
        server {{ $c.Name }}-{{ $i }}.{{ $root.Spec.Config.DNS.RootDomain }}:{{ $port.ContainerPort }};{{ end }}
    }
    server {
        listen     {{ $port.ContainerPort }};
        proxy_pass {{ $upstream }};
    }{{ end}}{{ end }}
}
`