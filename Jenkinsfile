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

    stage('Verify') {
      steps {
        sh 'docker version'
        sh 'docker info'
        sh 'docker compose version'
      }
    }

    stage('Test') {
      steps {
        sh 'go test ./...'
      }
    }
  }

  post {
    always {
      cleanWs()
    }
  }
}