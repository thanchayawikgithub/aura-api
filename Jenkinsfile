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

    stage('Deploy') {
      steps {
        sh 'docker compose down -v'
        sh 'docker compose up --build'
      }
    }


  }

  post {
    always {
      cleanWs()
    }
  }
}