package templates

// DeployPipeline defined default jenkins pipeline
const DeployPipeline = `
pipeline {
    agent {
        kubernetes {
            defaultContainer 'jnlp'
            yaml """
apiVersion: v1
kind: Pod
metadata:
  namespace: {{ .Namespace }}
spec:
  containers:
  {{- range $i, $item := .ContainerTemplates }}
  - name: {{ $item.Name }}
    image: {{ $item.Image }}
    workingDir: {{ $item.WorkingDir }}
    command:
    {{- range $cmd := $item.CommandArr }}
    - {{ $cmd }}
    {{- end }}
    args:
    {{- range $arg := $item.ArgsArr }}
    - {{ $arg }}
    {{- end }}
    tty: true
  {{- end }}
""" 
        }
    }
    environment {
        {{- range $i, $item := .EnvVars }}
        def {{ $item.Key }} = '{{ $item.Value }}'
        {{- end }}
    }
    stages {
        stage('HealthCheck') {
            {{ if .HealthCheckItems }}
            parallel {
                {{- range $i, $item := .HealthCheckItems }}
                stage('{{ $item.Name }}') {
                    steps {
                        {{ $item.Command }}
                    }
                }
                {{- end }}
            }
            {{ else }}
                steps {
                    sh "echo 'there was no health ckeck items'"
                }
            {{ end }}
        }
        stage('Callback') {
            steps {
                retry(count: 5) {
                    httpRequest acceptType: 'APPLICATION_JSON', contentType: 'APPLICATION_JSON', customHeaders: [[maskValue: true, name: 'Authorization', value: 'Bearer {{ .CallBack.Token }}']], httpMode: 'POST', requestBody: '''{{ .CallBack.Body }}''', responseHandle: 'NONE', timeout: 10, url: '{{ .CallBack.URL }}'
                }
            }
        }
    }
}
`
