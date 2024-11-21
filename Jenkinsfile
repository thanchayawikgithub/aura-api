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

    stage('build') {
      steps {
        sh 'docker build -t aura-api:latest .'
      }
    }

    stage('run') {
      steps {
        sh 'docker run -p 8081:8081 aura-api:latest'
      }
    }
  }

  post {
    always {
      cleanWs()
    }
  }
}