pipeline {
  agent any

  tools {
    go '1.22.6'
    dockerTool '27.3.1'
  }

  stages {
    stage('Checkout') {
      steps {
        git branch: 'main', url: 'https://github.com/thanchayawikgithub/aura-api.git'
      }
    }

    // stage('Build') {
    //   steps {
    //     sh 'go build -o ./bin/aura ./cmd/aura/main.go'
    //   }
    // }

    stage('Test') {
      steps {
        sh 'go test ./...'
      }
    }

    // stage('Docker Build') {
    //   steps {
    //     sh 'docker build -t aura-app .'
    //   }
    // }

    stage('Deploy Local') {
      steps {
        sh 'docker compose down'
        sh 'docker compose up -d --build'
      }
    }
  }

  post {
    always {
      cleanWs()
    }
  }
}