package templates

// CIPipeline defined default jenkins pipeline (default)
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

	parameters {
		{{- range $i, $item := .Parameters }}
		string(name: '{{ $item.Name }}', defaultValue: '{{ $item.DefaultValue }}', description: '{{ $item.Description }}')
		{{- end}}
	}

    environment {
        {{- range $i, $item := .EnvVars }}
        def {{ $item.Key }} = '{{ $item.Value }}'
        {{- end }}
    }
    stages {
        {{ .Stages }}
    }
	{{if .CallBack }}
	post {
		always {
			script {
				httpRequest httpMode: 'POST', url: ' {{ .CallBack.URL }}', contentType: 'APPLICATION_JSON', requestBody: '''
                {
					"pid": "${PID}",
					"appid": "${APPID}",
					"branch": "${params.BRANCH_NAME}",
					"image": "${env.SGR_REPOSITORY_NAME}:v${env.BUILD_NUMBER}",
					"buildNumber": "${env.BUILD_NUMBER}",
                    "message": "Pipeline completed",
                    "status": "${currentBuild.currentResult}",
                    "timestamp": "${new Date().toString()}"
                }
                '''
			}
		}
	}
	{{ else }}

	{{ end }}
}
`

// Checkout ..
const Checkout = `
stage('Checkout') {
    {{if .CheckoutItems }}
    stages {
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

const CICD = `
{{- range $i, $item := .pipelineStages }}
		stage('{{ $item.Name }}') {
			steps {
			{{- range $j, $step := $item.Steps }}
			 {{ $step.Command }}
			{{- end }}
			}
		}
{{- end }}
`

// Compile stage
const Compile = `
stage('Compile') {
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
stage('Build') {
    {{if .ImageItems }}
	 parallel {
      {{- range $i, $item := .ImageItems }}
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
            sh "echo 'there was no images items'"
        }
    {{ end }}
}
`

const DeployImage = `
stage('Deploy') {
    {{if .DeployItems }}
	parallel {
      {{- range $i, $item := .DeployItems }}
       stage('{{ $item.Name }}') {
		   steps {
			   {{ $item.Command }}
		   }
       }
      {{- end }}
	}
    {{ else }}
        steps {
            sh "echo 'there was no deploy items'"
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
