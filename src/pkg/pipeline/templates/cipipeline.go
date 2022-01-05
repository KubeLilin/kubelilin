package templates

// CIPipeline defined default jenkins pipeline
const CIPipeline = `
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
        {{ .Stages }}

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

// Checkout ..
const Checkout = `
stage('Checkout') {
    {{if .CheckoutItems }}
    parallel {
        {{- range $i, $item := .CheckoutItems }}
        stage('{{ $item.Name }}') {
            steps {
                {{ $item.Command }}
            }
        }
        {{- end }}
    }
    {{ else }}
        steps {
            sh "echo 'there was no checkout items'"
        }
    {{ end }}
}
`

// Compile stage
const Compile = `
stage('Builds') {
    {{if .BuildItems }}
    parallel {
        {{- range $i, $item := .BuildItems }}
        stage('{{ $item.Name }}') {
            steps {
                container('{{ $item.ContainerName }}') {
                    {{ $item.Command }}
                }
            }
        }
        {{- end }}
    }
    {{ else }}
        steps {
            sh "echo 'there was no build items'"
        }
    {{ end }}
}
`

// BuildImage stage
const BuildImage = `
stage('Images') {
    {{if .ImageItems }}
    parallel {
        {{- range $i, $item := .ImageItems }}
        stage('{{ $item.Name }}') {
            steps {
                container("kaniko") {
                    sh "[ -d $DOCKER_CONFIG ] || mkdir -pv $DOCKER_CONFIG"

                    sh """
                    echo '{"auths": {"'$REGISTRY_ADDR'": {"auth": "'$DOCKER_AUTH'"}}}' > $DOCKER_CONFIG/config.json
                    """
                    {{ $item.Command }}
                }
            }
        }
        {{- end }}
    }
    {{ else }}
        steps {
            sh "echo 'there was no images items'"
        }
    {{ end }}
}
`

// CustomScript stage
const CustomScript = `
stage({{ .CustomScriptItem.Name }}) {
    steps {
        {{ .CustomScriptItem.Command }}
    }
}
`
