pipeline {
	agent { label "dev" }
	tools { go "go-1.27.1" }
	stages {
		stage("Download dependency") {
			steps {
				sh "go mod tidy"
			}
		}
		stage("Test") {
			steps {
				sh "go test -count=1 -coverprofile -v ./..."
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
