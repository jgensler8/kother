# kother

A version control and CI/CD focused Kubernetes cluster manager.

## Goals

* Provide a one click way of getting a Kubernetes cluster into your CI/CD pipeline.
* Value Simplicity over Efficiency

## Notes

A File is one of two types:
  * Variable File
  * Template File

The only variable files are:

* cluster.yaml
* cloud-provider.yaml
* domains.yaml

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

```

### Future Work

* Service to tell you when to update stuff. Branch, commit, and stand up environment.
* Bake each component into an AMI