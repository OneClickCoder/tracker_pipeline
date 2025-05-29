pipeline {
    agent any
    // removed the options{} block - count 7

    stages {
        stage('Checkout Code') {
            steps {
                echo "Cloning repository..."
                // Use the SSH credential you configured in Jenkins
                git branch: 'main',
                    credentialsId: 'github-ssh-key', // <-- Make sure this ID matches your Jenkins credential ID
                    url: 'git@github.com:OneClickCoder/tracker_pipeline.git'
                echo "Repository cloned successfully."
                sh "ls -la" // List files to confirm checkout
            }
        }
    }

    post {
        always {
            cleanWs() // Clean up workspace after build
        }
        success {
            echo "Pipeline finished successfully!"
        }
        failure {
            echo "Pipeline failed!"
        }
    }
}