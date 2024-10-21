pipeline {
  agent any
  stages {
    stage('Checkout Code') {
      steps {
        git(url: 'https://github.com/steveodhiambo/Structy', branch: 'main')
      }
    }

    stage('Structy Tests') {
      steps {
        sh 'go test -v ./...'
      }
    }

  }
}