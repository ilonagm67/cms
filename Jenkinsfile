pipeline {
	agent { label "dev" }
	tools { go "go-1.27.1" }
	stages {
		stage("Test") {
			steps {
				sh "go vet ."
				sh "go test ./..."
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
