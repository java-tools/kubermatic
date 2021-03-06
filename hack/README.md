# Scripts list and their description

This file is generated by running `hack/meta.sh`.

## changelog-gen.sh

Create a changelog since last release, commit and create a new release tag

    Usage:
    changelog-gen.sh -r v2.x.x - create changelog, commit and tag new release, using closed PRs release-note

## coverage.sh

Generate test coverage statistics for Go packages.

Works around the fact that `go test -coverprofile` currently does not work  
with multiple packages, see https://code.google.com/p/go/issues/detail?id=6909

    Usage: cover.sh [--html]

    --html      Additionally create HTML report and open it in browser

## gen-api-client.sh

Generates the KKP API Swagger spec and client. The generated client is then
used in the api-e2e tests and published into https://github.com/kubermatic/go-kubermatic

## lib.sh

Contains commonly used functions for the other scripts.

## meta.sh

This README.md generator

That generates README.md with all scripts from this directory described.

## publish-s3-exporter.sh

Releases a new quay.io/kubermatic/s3-exporter Docker image; must be
run manually whenever the s3-exporter is updated.

## release-docker-images.sh

Builds and pushes all KKP Docker images:

* quay.io/kubermatic/kubermatic[-ee]
* quay.io/kubermatic/addons
* quay.io/kubermatic/openshift-addons
* quay.io/kubermatic/nodeport-proxy
* quay.io/kubermatic/kubeletdnat-controller
* quay.io/kubermatic/user-ssh-keys-agent
* quay.io/kubermatic/etcd-launcher

The images are tagged with all arguments given to the script, i.e
`./release-docker-images.sh foo bar` will tag `kubermatic:foo` and
`kubermatic:bar`.

Before running this script, all binaries in `cmd/` must have been
built already by running `make build`.

## retag-images.sh

This script takes YAML on stdin and tries to find all references
to Docker images. The found images will then be pulled, retagged
with `$TARGET_REGISTRY` as their registry and pushed to said target
registry.

The purpose of this is to mirror all images used in a KKP setup
to prewarm a local Docker registry, for example in offline setups.

## run-api.sh

TBD

## run-dashboard-and-api.sh

TBD

## run-machine-controller.sh

TBD

## run-master-controller-manager.sh

TBD

## run-operator.sh

TBD

## run-seed-controller-manager.sh

TBD

## run-user-cluster-controller-manager.sh

TBD

## update-cert-manager-crds.sh

TBD

## update-codegen.sh

TBD

## update-docs.sh

TBD

## update-fixtures.sh

TBD

## update-grafana-dashboards.sh

TBD

## update-kubermatic-chart.sh

TBD

## update-openshift-version-codegen.sh

This script can be used to update the generated image names for Openshift.  
The desired versions msut be configured first in  
codegen/openshift_versions/main.go and a const for each version must be  
added to pkg/controller/openshift/resources/const.go

Also, executing this script requires access to the ocp quay repo.

## update-prometheus-rules.sh

TBD

## update-swagger.sh

TBD

## update-velero-crds.sh

TBD

## verify-api-client.sh

TBD

## verify-boilerplate.sh

TBD

## verify-codegen.sh

TBD

## verify-docs.sh

TBD

## verify-forbidden-functions.sh

TBD

## verify-grafana-dashboards.sh

TBD

## verify-kubermatic-chart.sh

TBD

## verify-licenses.sh

TBD

## verify-prometheus-rules.sh

TBD

## verify-swagger.sh

TBD

