#!groovy
pipeline {
    options {
        buildDiscarder(logRotator(daysToKeepStr: '30'))
        timestamps()
        ansiColor('xterm')
        timeout(time: 3, unit: 'HOURS')
    }
    environment {
        DOCKER_BUILDKIT = '1'
    }
    agent {
        node {
            label 'amd64 && ubuntu-1804 && overlay2'
        }
    }

    stages {
        stage('CLI Integration Test') {
            steps {
                sh 'TEST_SKIP_INTEGRATION=1 make test-integration-cli'
            }
            post {
                always {
                    junit testResults: 'components/engine/bundles/test-integration/*junit-report.xml', allowEmptyResults: true
                }
            }
        }
    }
}
