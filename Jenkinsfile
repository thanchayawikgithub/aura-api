pipeline {
  agent any

  tools {
    go '1.22.6'
    dockerTool '27.3.1'
  }

  stages {
    stage('Verify') {
      steps {
        sh 'docker version'
        sh 'docker info'
        sh 'docker compose version'
      }
    }

    stage('Checkout') {
      steps {
        git branch: 'main', url: 'https://github.com/thanchayawikgithub/aura-api.git'
      }
    }

    stage('Test') {
      steps {
        sh 'go test ./...'
      }
    }

    
    // stage('Deploy Local') {
    //   steps {
    //     sh 'docker compose version'
    //     sh 'docker compose down'
    //     sh 'docker compose up -d --build'
    //   }
    // }
  }

  post {
    always {
      cleanWs()
    }
  }
}