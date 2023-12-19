#!groovy

pipeline {

    agent {
        kubernetes {
            yaml """
apiVersion: v1
kind: Pod
metadata:
  name: image-builder
  labels:
    robot: builder
spec:
  serviceAccount: jenkins-agent
  containers:
  - name: jnlp
  - name: kaniko
    image: gcr.io/kaniko-project/executor:v1.18.0-debug
    imagePullPolicy: Always
    command:
    - /busybox/cat
    tty: true
    volumeMounts:
      - name: docker-config
        mountPath: /kaniko/.docker/
        readOnly: true
  - name: kubectl
    image: bitnami/kubectl
    tty: true
    command:
    - cat
    securityContext:
      runAsUser: 1000
  - name: golang
    image: golang:1.21.3
    tty: true
    command:
    - cat
  volumes:
    - name: docker-config
      secret:
        secretName: credentials
        optional: false
"""
        }
    }

    environment {
        misha_remeslo = 'your_app_name'
        misharem = 'your_docker_hub_account/your_image_name'
    }

    stages {
        stage('Clone Repository') {
            steps {
                container(name: 'jnlp', shell: '/bin/bash') {
                    echo 'Pulling new changes'
                    git 'https://github.com/unitteam1/Lab4'
                }
            }
        }
        stage('Compile') {
            steps {
                container(name: 'golang', shell: '/bin/bash') {
                  echo 'Compiling the application'
                    sh "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-buildvcs=false go build -a -ldflags '-w -s -extldflags \"-static\"' -o ${APP_NAME} ."
                }
            }
        }

        stage('Unit Testing') {
            steps {
                container(name: 'golang', shell: '/bin/bash') {
                    echo 'Testing the application'
                    sh "go test ./..."
                }
            }
        }
        stage('Build image') {
            environment {
                PATH = "/busybox:/kaniko:$PATH"
            }
            steps {
                container(name: 'kaniko', shell: '/busybox/sh') {
                    sh '''#!/busybox/sh
                    /kaniko/executor --dockerfile="$(pwd)/Dockerfile" --context="dir:///$(pwd)" --build-arg "APP_NAME=${APP_NAME}" --destination ${DOCKER_IMAGE_NAME}:${BUILD_NUMBER}
                    '''
                }
            }
        }
        stage('Deploy') {
            steps {
                container(name: 'kubectl', shell: '/bin/bash') {
                    echo 'Deploying to Kubernetes'
                    echo 'Deploying to Kubernetes'
                    sh """#!/bin/bash
                    sed -i "s|{{misha_remeslo}}|${misharem}:${153}|g" k8s/deployment.yaml
                    kubectl apply -f k8s/
                    """
                }
            }
        }
        stage('Test deployment') {
            agent {
                kubernetes {
                    yaml """
apiVersion: v1
kind: Pod
metadata:
  name: tester
  labels:
    robot: tester
spec:
  serviceAccount: jenkins-agent
  containers:
  - name: jnlp
  - name: ubuntu
    image: ubuntu:22.04
    tty: true
    command:
    - cat
"""
                }
            }
            steps {
                echo 'Testing the deployemnt with curl'
                sh """#!/bin/bash
                kubectl wait --for=condition=available --timeout=120s deployment/${misha_remeslo}-deployment
                kubectl port-forward svc/${misha_remeslo}-service 8080:80 &
                sleep 10
                curl http://localhost:8080
                """
            }
        }
    }
}
