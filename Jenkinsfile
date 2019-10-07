#!groovy
pipeline {
    options {
        buildDiscarder(logRotator(daysToKeepStr: '30'))
        timestamps()
        ansiColor('xterm')
        timeout(time: 2, unit: 'HOURS')
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
                    sh '''
                    echo "Ensuring container killed."
                    cids=$(docker ps -aq -f name=docker-pr${BUILD_NUMBER}-*)
                    [ -n "$cids" ] && docker rm -vf $cids || true
                    '''

                    sh '''
                    echo "Chowning /workspace to jenkins user"
                    docker run --rm -v "$WORKSPACE:/workspace" busybox chown -R "$(id -u):$(id -g)" /workspace
                    '''

                    junit testResults: 'components/engine/bundles/test-integration/*junit-report.xml', allowEmptyResults: true

                    catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE', message: 'Failed to create bundles.tar.gz') {
                        sh '''
                        bundleName=amd64
                        echo "Creating ${bundleName}-bundles.tar.gz"
                        # exclude overlay2 directories
                        find components/engine/bundles -path '*/root/*overlay2' -prune -o -type f \\( -name '*-report.json' -o -name '*.log' -o -name '*.prof' -o -name '*-report.xml' \\) -print | xargs tar -czf ${bundleName}-bundles.tar.gz
                        '''

                        archiveArtifacts artifacts: '*-bundles.tar.gz', allowEmptyArchive: true
                    }
                }
                cleanup {
                    sh 'make clean'
                    deleteDir()
                }
            }
        }
    }
}
