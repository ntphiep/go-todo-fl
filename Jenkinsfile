pipeline {
    agent any
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/ntphiep/go-todo-pg', branch: 'main'
            }
        }
        stage('Build') {
            steps {
                sh 'go build -v ./...'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v ./...'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...'
                // Thêm các bước triển khai thực tế vào đây
            }
        }
    }
}
