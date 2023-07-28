# Konveyor Addon - Move2Kube
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkonveyor%2Faddon-move2kube.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkonveyor%2Faddon-move2kube?ref=badge_shield)


This project exposes Konveyor's Move2Kube tool https://github.com/konveyor/move2kube as an addon for Tackle 2 Hub https://github.com/konveyor/tackle2-hub , enabling applications to have container images and Kubernetes manifests built for them.

For further documentation see https://move2kube.konveyor.io/

## Instructions

1. Install OLM, the Tackle 2 Hub operator and run a instance of Tackle. https://konveyor.github.io/tackle/installation/

2. Apply the CR to install the Move2Kube addon.

```console
$ kubectl apply -f addon.yaml
```

3. Open the Tackle UI, create an application, add the repo URL and credentials for the repo. Call the API to get the application ID.

```console
curl -i localhost/hub/applications
```

4. Call the tasks API to run the addon on an application.

```console
curl -i -X POST localhost/hub/tasks -d \
'{
    "name":"move2kube",
    "state": "Ready",
    "locator": "move2kube",
    "addon": "move2kube",
    "application": {"id": 1},
    "data": {}
}'
```

5. To provide a config encode the YAML as base64 for the key `config-base64` in the request data.

```console
$ curl -i -X POST localhost/hub/tasks -d \
'{
    "name":"move2kube",
    "state": "Ready",
    "locator": "move2kube",
    "addon": "move2kube",
    "application": {"id": 1},
    "data": {"config-base64": "bW92ZTJrdWJlOgogIG1pbnJlcGxpY2FzOiAiMiIKICBzZXJ2aWNlczoKICAgIGphdmEtZ3JhZGxlOgogICAgICAiOTA4MCI6CiAgICAgICAgc2VydmljZXR5cGU6IEluZ3Jlc3MKICAgICAgICB1cmxwYXRoOiAvamF2YS1ncmFkbGUtMTIzNAogICAgICBjaGlsZE1vZHVsZXM6CiAgICAgICAgamF2YS1ncmFkbGU6CiAgICAgICAgICBwb3J0OiAiODA4MCIKICAgICAgZG9ja2VyZmlsZVR5cGU6IGJ1aWxkIHN0YWdlIGluIGJhc2UgaW1hZ2UKICAgICAgZW5hYmxlOiB0cnVlCiAgICAgIHdhcnRyYW5zZm9ybWVyOiBMaWJlcnR5CiAgc3Bhd25jb250YWluZXJzOiBmYWxzZQogIHRhcmdldDoKICAgIGRlZmF1bHQ6CiAgICAgIGNsdXN0ZXJ0eXBlOiBLdWJlcm5ldGVzCiAgICAgIGluZ3Jlc3M6CiAgICAgICAgaG9zdDogbXlwcm9qZWN0LTEyMzQuY29tCiAgICAgICAgaW5ncmVzc2NsYXNzbmFtZTogIiIKICAgICAgICB0bHM6ICIiCiAgICBpbWFnZXJlZ2lzdHJ5OgogICAgICBuYW1lc3BhY2U6IG15cHJvamVjdAogICAgICBxdWF5LmlvOgogICAgICAgIGxvZ2ludHlwZTogbm8gYXV0aGVudGljYXRpb24KICAgICAgdXJsOiBxdWF5LmlvCiAgdHJhbnNmb3JtZXJzOgogICAgdHlwZXM6CiAgICAgIC0gTWF2ZW4KICAgICAgLSBSdXN0LURvY2tlcmZpbGUKICAgICAgLSBSZWFkTWVHZW5lcmF0b3IKICAgICAgLSBDbG91ZEZvdW5kcnkKICAgICAgLSBKYXIKICAgICAgLSBLdWJlcm5ldGVzVmVyc2lvbkNoYW5nZXIKICAgICAgLSBadXVsQW5hbHlzZXIKICAgICAgLSBMaWJlcnR5CiAgICAgIC0gR29sYW5nLURvY2tlcmZpbGUKICAgICAgLSBDb21wb3NlQW5hbHlzZXIKICAgICAgLSBQSFAtRG9ja2VyZmlsZQogICAgICAtIFdhckFuYWx5c2VyCiAgICAgIC0gQXJnb0NECiAgICAgIC0gUHl0aG9uLURvY2tlcmZpbGUKICAgICAgLSBEb3ROZXRDb3JlLURvY2tlcmZpbGUKICAgICAgLSBFYXJBbmFseXNlcgogICAgICAtIE5vZGVqcy1Eb2NrZXJmaWxlCiAgICAgIC0gSmJvc3MKICAgICAgLSBEb2NrZXJmaWxlUGFyc2VyCiAgICAgIC0gR3JhZGxlCiAgICAgIC0gUnVieS1Eb2NrZXJmaWxlCiAgICAgIC0gVG9tY2F0CiAgICAgIC0gQ2x1c3RlclNlbGVjdG9yCiAgICAgIC0gRG9ja2VyZmlsZUltYWdlQnVpbGRTY3JpcHQKICAgICAgLSBLdWJlcm5ldGVzCiAgICAgIC0gT3BlcmF0b3JzRnJvbVRDQQogICAgICAtIFRla3RvbgogICAgICAtIFdhclJvdXRlcgogICAgICAtIENvbXBvc2VHZW5lcmF0b3IKICAgICAgLSBDTkJDb250YWluZXJpemVyCiAgICAgIC0gQnVpbGRjb25maWcKICAgICAgLSBLbmF0aXZlCiAgICAgIC0gV2luV2ViQXBwLURvY2tlcmZpbGUKICAgICAgLSBEb2NrZXJmaWxlRGV0ZWN0b3IKICAgICAgLSBDb250YWluZXJJbWFnZXNQdXNoU2NyaXB0R2VuZXJhdG9yCiAgICAgIC0gT3BlcmF0b3JUcmFuc2Zvcm1lcgogICAgICAtIFBhcmFtZXRlcml6ZXIKICAgICAgLSBFYXJSb3V0ZXIKICB0cmFuc2Zvcm1lcnNlbGVjdG9yOiAiIgo="}
}'
```

6. You can also provide the config as JSON using the key `config` in the request data.

```
$ curl -i -X POST localhost/hub/tasks -d \
'{
    "name":"move2kube",
    "state": "Ready",
    "locator": "move2kube",
    "addon": "move2kube",
    "application": {"id": 1},
    "data": {"config": {"move2kube": {"target": {"default": {"ingress": {"host": "myproject-1234.com"} } } } } }
}'
```

## Options

```
bool                   dont-copy-config-to-output
string                 commit-message
string                 output-branch
string                 output-dir
string                 config-base64
map[string]interface{} config
```


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkonveyor%2Faddon-move2kube.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkonveyor%2Faddon-move2kube?ref=badge_large)