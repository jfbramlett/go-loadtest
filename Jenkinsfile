pipeline {
    agent {
        label "Docker-Large"
    }
    options {
        skipStagesAfterUnstable()
    }
    stages {
        stage('Build') {
            steps {
                sh './build-docker.sh'
            }
        }
        stage('Run') {
            steps {
                sh './run-dev-scenario.sh'
            }
        }
    }
    post {
        always {
            script {
            sh './clean-local-images.sh'
            }
        }
        success {
            script{
                env.BUILD_RESULT = 'SUCCESS'
            }
        }

        failure {
            script{
                env.BUILD_RESULT = 'FAILURE'
            }
            emailext body: '''$PROJECT_NAME - Build # $BUILD_NUMBER - $BUILD_STATUS:
Check console output at $BUILD_URL to view the results.''',  replyTo: 'cicdpipeline@ninth-wave.com', subject: '$PROJECT_NAME - Build # $BUILD_NUMBER - $BUILD_STATUS!', to: 'jbramlett@ninth-wave.com'

        }
        changed {
            script {
                if (env.BUILD_RESULT != 'FAILURE') {
                    emailext body: '''$PROJECT_NAME - Build # $BUILD_NUMBER - $BUILD_STATUS:
        Check console output at $BUILD_URL to view the results.''',  replyTo: 'cicdpipeline@ninth-wave.com', subject: '$PROJECT_NAME - Build # $BUILD_NUMBER - $BUILD_STATUS!', to: 'jbramlett@ninth-wave.com'
                }
            }
        }
    }
}