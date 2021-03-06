source:
  kind: helmChart
  name: "Get latest jenkins helm chart version"
  spec:
    url: https://charts.jenkins.io
    name: jenkins

conditions:
  exist:
    name: "Is Jenkins helm chart available on Registry?"
    kind: helmChart
    spec:
      url: https://kubernetes-charts.storage.googleapis.com
      name: jenkins
  chartDependencyIsJenkins:
    name: "Is Jenkins dependency correclty set?"
    kind: yaml
    spec:
      file: "charts/jenkins/requirements.yaml"
      key: "dependencies[0].name"
      value: "jenkins"
    scm:
      github:
        user: "{{ .github.user }}"
        email: "{{ .github.email }}"
        owner: "{{ .github.owner }}"
        repository: "{{ .github.repository }}"
        token: "{{ requiredEnv .github.token }}"
        username: "{{ .github.username }}"
        branch: "{{ .github.branch }}"

targets:
  imageTag:
    name: "Update required Jenkins helm chart version"
    kind: yaml
    spec:
      file: "charts/jenkins/requirements.yaml"
      key: "dependencies[0].version"
    scm:
      github:
        user: "{{ .github.user }}"
        email: "{{ .github.email }}"
        owner: "{{ .github.owner }}"
        repository: "{{ .github.repository }}"
        token: "{{ requiredEnv .github.token }}"
        username: "{{ .github.username }}"
        branch: "{{ .github.branch }}"
