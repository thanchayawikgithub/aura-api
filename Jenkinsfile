pipeline {
  agent any

  tools {
    go '1.22.6'
    dockerTool '27.3.1'
  }

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