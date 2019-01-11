pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile'
    }

  }
  stages {
    stage('build') {
      agent {
        dockerfile {
          filename 'Dockerfile'
        }

      }
      steps {
        echo 'building from dockerfile'
      }
    }
    stage('test') {
      agent any
      steps {
        sleep 10
      }
    }
  }
}