pipeline {
    agent {
        kubernetes
    }

    parameters {
        choice(name: 'OS', choices: ['linux', 'darwin', 'windows', 'all'], description: 'Pick OS')
        choice(name: 'ARCH', choices: ['amd64', 'arm64'], description: 'Pick ARCH')
        string(name: 'REGISTRY', defaultValue: 'docker.io/viktordevopscourse', description: 'Enter registry')
    }

    environment {
        TELE_TOKEN = credentials('TELEGRAM_TOKEN')
        DOCKER_HUB_CREDENTIALS = credentials('dockerhub-credentials-id')
    }

    stages {
        stage('test') {
            steps {
                echo "Run unit tests"
                sh 'make test'
            }
        }
        stage('build') {
            steps {
                echo "Build for platform ${params.OS} arch: ${params.ARCH}"
                sh "make image TARGETOS=${OS} TARGETARCH=${ARCH}"
            }
        }
        stage('publish') {
            steps {
                echo "login to docker hub"
                withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'DOCKER_HUB_CREDENTIALS', usernameVariable: 'DOCKER_HUB_USERNAME', passwordVariable: 'DOCKER_HUB_PASSWORD']]) {
                    sh "docker login -u ${DOCKER_HUB_USERNAME} -p ${DOCKER_HUB_PASSWORD}"
                }
                sh "make push REGISTRY=${parems.REGISTRY} TARGETOS=${OS} TARGETARCH=${ARCH}"
            }
        }
    }
}