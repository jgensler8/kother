package vagrantfile

var vagrantfile_template string =`{{ $root := . }}
Vagrant.configure("2") do |config|
  config.vm.box = "coreos-stable"
  config.vm.box_check_update = false
  config.vm.network "public_network", type: "dhcp"
  config.landrush.enabled = true
  config.landrush.tld = '{{ .Config.DNS.RootDomain }}'

  {{ loadBalancerBlock $root.Etcd.Component $root }}
  {{ componenetBlock $root.Etcd.Component $root }}

  {{ loadBalancerBlock $root.Vault.Component $root }}
  {{ componenetBlock $root.Vault.Component $root }}

  {{ loadBalancerBlock $root.APIServer.Component $root }}
  {{ componenetBlock $root.APIServer.Component $root }}

  {{ loadBalancerBlock $root.ControllerManager.Component $root }}
  {{ componenetBlock $root.ControllerManager.Component $root }}

  {{ loadBalancerBlock $root.Scheduler.Component $root }}
  {{ componenetBlock $root.Scheduler.Component $root }}

end
`

var vagrantfile_loadbalancer_template = `
{{ $root := . }}{{ $pod := getPod $root.Component }}
config.vm.define "{{ $pod.Name }}-lb" do |lb|
    lb.vm.hostname = "{{ $pod.Name}}-lb.{{ $root.Spec.Config.DNS.RootDomain }}"
    lb.vm.provision "shell", inline: "echo '{{ getNGINXConfig $root.Component $root.Spec | b64enc }}' | base64 -d -w 0 > /tmp/nginx.conf"
    lb.vm.provision "shell", inline: "docker run -d {{ getPortString $root.Component }} --net=host --privileged --name nginx --restart always -v /tmp/nginx.conf:/etc/nginx/nginx.conf nginx"
    lb.vm.provider "virtualbox" do |v|
	v.memory = 256
	v.cpus = 1
    end
end`

var vagrantfile_component_template = `{{ $root := . }}{{ $pod := getPod $root.Component }}
(1..{{ $root.Component.Replicas }}).each do |i|
	config.vm.define "{{ $pod.Name }}-%d" % i do |target|
	    target.vm.hostname = "{{ $pod.Name }}-%d.{{ $root.Spec.Config.DNS.RootDomain }}" % i
	    target.vm.provision "shell", inline: "echo '{{ (getUserData $root.Component $root.Spec) | b64enc }}' | base64 -d -w 0 > /var/lib/coreos-vagrant/vagrantfile-user-data"
	    target.vm.provider "virtualbox" do |v|
		v.memory = 512
		v.cpus = 1
	    end
	end
end`