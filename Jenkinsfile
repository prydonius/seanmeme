pipeline {
  environment {
    IMAGE_NAME = 'prydonius/seanmeme'
  }

  agent any

  stages {
    stage('Build') {
      steps {
        checkout scm
        sh '''
          docker run --rm -v "${PWD}":/go/src/seanmeme -w /go/src/seanmeme -e CGO_ENABLED=0 golang:1.8 go build
          docker build -t $IMAGE_NAME:$BUILD_ID .
        '''
      }
    }
    stage('Test') {
      steps {
        echo 'TODO: add tests'
      }
    }
    stage('Image Release') {
      when {
        expression { env.BRANCH_NAME == 'master' }
      }

      steps {
        withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'dockerhub',
          usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD']]) {
          sh '''
            docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
            docker push $IMAGE_NAME:$BUILD_ID
          '''
        }
      }
    }
    stage('Staging Deployment') {
      when {
        expression { env.BRANCH_NAME == 'master' }
      }

      environment {
        RELEASE_NAME = 'seanmeme-staging'
        SERVER_HOST = 'staging.seanmeme.k8s.prydoni.us'
      }

      steps {
        sh '''
          . ./helm/helm-init.sh
          helm upgrade --install --namespace staging $RELEASE_NAME ./helm/seanmeme --set image.tag=$BUILD_ID,ingress.host=$SERVER_HOST
        '''
      }
    }
    stage('Deploy to Production?') {
      when {
        expression { env.BRANCH_NAME == 'master' }
      }

      steps {
        // Prevent any older builds from deploying to production
        milestone(1)
        input 'Deploy to Production?'
        milestone(2)
      }
    }
    stage('Production Deployment') {
      when {
        expression { env.BRANCH_NAME == 'master' }
      }

      environment {
        RELEASE_NAME = 'seanmeme-production'
        SERVER_HOST = 'seanmeme.k8s.prydoni.us'
      }

      steps {
        sh '''
          . ./helm/helm-init.sh
          helm upgrade --install --namespace production $RELEASE_NAME ./helm/seanmeme --set image.tag=$BUILD_ID,ingress.host=$SERVER_HOST
        '''
      }
    }
  }
}
