# kother

A version control and CI/CD focused Kubernetes cluster manager.

## Goals

* Provide a one click way of getting a Kubernetes cluster into your CI/CD pipeline.
* Value Simplicity over Efficiency
* One Cluster Per VPC Per Region (multiple vpcs in one region == multiple clusters)

## Notes

There is one file for variables (`cluster.yaml`).

All files in the `cloud-provider` and `kubernetes` directory will be run through a templating engine before being UnMarshalled. These variables can be consumed in this process.

There are some extra variables that will be provided to the templating engine:

* Config.DNS.etcd
* Config.DNS.apiserver
* Config.DNS.controller-manager
* Config.DNS.scheduler

TODO:

* Context.commithash
* Context.committag

## Opinion Table

| Operating System | Cloud Provider | Certificate Manager | Kubernetes | Overlay Network | Container Runtime |
| --- | --- | --- | --- | --- | --- |
| CoreOS | AWS | Vault | Open Source | Flannel (vxlan) | Docker |

## How to use

### Function 1: Generate the Skeleton

```bash
kother init --name myproject
cd myproject
...
git init
git ...
git ...
git ...
git push ...
```

### Function 2: Usage in Deployment Pipeline

```bash
kother validate
kother deploy
# OR
# for feature branching
kother deploy --branch-hash ${COMMIT_VARIABLE}
```

### Function 3: Dry Run 

```bash
kother dump
```

## Project Structure

```bash
examples
└── myproject
    ├── cloud-provider
    │   └── manifests
    │       ├── etcd.yaml
    │       ├── kube-apiserver.yaml
    │       ├── kube-controller-manager.yaml
    │       ├── kube-scheduler.yaml
    │       └── vault.yaml
    ├── cluster.yaml
    └── kubernetes
        ├── dns
        │   ├── dns-deployment.yaml
        │   └── dns-svc.yaml
        └── proxy
            └── proxy.yaml

```

### Future Work

* Use Godep, then switch to glog to logging
* Service to tell you when to update stuff. Branch, commit, and stand up environment.
* Subcommand for Generating a (large) Vagrantfile
* Provide logic to pull ENV vars from particular CI systems (Bamboo, Jenkins, Travis, etc)
* Bake each component into an AMI
* Chang Cloud Config back to Ignition ([not working at the time of this article](https://github.com/coreos/ignition/blob/95f018b7cdfb45386cb6f5c99facc0b8a888f343/internal/oem/oem.go#L174-L177))