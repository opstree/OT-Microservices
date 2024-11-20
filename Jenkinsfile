pipeline {
    agent {
        docker { image 'golang:1.18' } // Use the official Golang Docker image
    }
    environment {
        GO_VERSION = '1.18' // Set the Go version if needed
        GOPATH = '/go' // Define Go workspace path
        DOCKER_IMAGE = 'rohit/employee-app' // Docker image name (you can update this)
        DOCKER_REGISTRY = 'docker.io' // Docker registry, e.g., Docker Hub (can be a private registry too)
        GITHUB_REPO = 'https://github.com/rohitdevops05112024/OT-Microservices.git' // GitHub repo URL
        BRANCH = 'rohit_test' // Branch to checkout
    }
    stages {
        stage('Checkout') {
            steps {
                // Checkout the code from GitHub repository
                git branch: "${BRANCH}", url: "${GITHUB_REPO}"
            }
        }

        stage('Install Dependencies') {
            steps {
                script {
                    // Install Go dependencies using `go mod`
                    sh 'go mod tidy'
                }
            }
        }

        stage('Build') {
            steps {
                script {
                    // Build the Go application
                    sh 'go build -o employee-app ./employee'
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    // Run Go tests
                    sh 'go test -v ./employee'
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    // Build the Docker image using the Dockerfile located in the employee directory
                    sh 'docker build -t ${DOCKER_IMAGE} ./employee'
                }
            }
        }

        stage('Publish Docker Image') {
            steps {
                script {
                    // Push the Docker image to the specified Docker registry (e.g., Docker Hub)
                    sh 'docker push ${DOCKER_IMAGE}'
                }
            }
        }

        stage('Clean Up') {
            steps {
                script {
                    // Clean up generated binaries and Docker images
                    sh 'rm -rf employee-app'
                    sh 'docker rmi ${DOCKER_IMAGE}'
                }
            }
        }
    }
    post {
        always {
            // Actions to be taken after the pipeline execution, such as cleaning the workspace
            cleanWs() // Clean up the workspace
        }
        success {
            // Actions to take on success
            echo 'Pipeline completed successfully.'
        }
        failure {
            // Actions to take on failure
            echo 'Pipeline failed.'
        }
    }
}
