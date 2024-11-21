pipeline {
  agent any

  stages {
     stage('Test') {
      steps {
        sh 'go test ./...'
      }
    }

    stage('Build & Deploy') {
      steps {
        sh 'docker compose down'
        sh 'docker compose up'
      }
    }
  }

  post {
    always {
      cleanWs()
    }
  }
}