pipeline {
	agent { label "dev" }
	tools { go "go-1.27.1" }
	stages {
		stage("Package Download") {
			steps {
				sh "go mod tidy"
			}
		}
		stage("Test") {
			steps {
				sh "go test -cover ./..."
			}
		}
		stage("Build") {
			steps {
				sh "go build ."
			}
		}
	}
	post {
		success {
			archiveArtifacts artifacts: "cms",fingerprint: true
		}
	}
}
