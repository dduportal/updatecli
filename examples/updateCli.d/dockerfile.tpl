---
source:
  name: Get Latest helm release version
  kind: githubRelease
  spec:
    owner: "helm"
    repository: "helm"
    token: {{ requiredEnv .github.token }}
    username: olblak
    version: latest
conditions:
  isENVSet:
    name: Is ENV HELM_VERSION set
    kind: dockerfile
    spec:
      file: docker/Dockerfile
      Instruction: ENV[1][0]
      Value: "HELM_VERSION"
    scm:
      github:
        user: "updatecli"
        email: "updatecli@olblak.com"
        owner: "olblak"
        repository: "charts"
        token: {{ requiredEnv "GITHUB_TOKEN" }}
        username: "olblak"
        branch: "main"
targets:
  updateENVHELMVERSION:
    name: Update HELM_VERSION
    kind: dockerfile
    spec:
      file: docker/Dockerfile
      Instruction: ENV[1][1]
    scm:
      github:
        user: "updatecli"
        email: "updatecli@olblak.com"
        owner: "olblak"
        repository: "charts"
        token: {{ requiredEnv "GITHUB_TOKEN" }}
        username: "olblak"
        branch: "main"
