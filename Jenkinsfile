pipeline {
    agent any
    tools {
        go 'Go'
    }
    
    environment{
        PROJECT = 'tm-product'
        PROJECT_IMAGE_TAG = ''
    }

    stages {
        stage('Build Test image'){
            steps{
                sh "docker build --target test -t ${env.PROJECT} ."
            }
        }
        // stage('Test'){
        //     steps{
        //         // 測試環境需使用docker ，使用掛載捲的方式
        //         sh "docker run -v /var/run/docker.sock:/var/run/docker.sock ${env.PROJECT}"
        //         sh 'echo y | docker system prune'
        //     }
        // }
        stage('Build Deployment image'){
            steps{
                echo 'build image'
                sh "docker build . -t ${env.PROJECT}"
            }
        }
        stage('Push image'){
            steps{
                echo 'docker login'
                withCredentials([usernamePassword(credentialsId: 'dockerHub', passwordVariable: 'dockerHubPassword', usernameVariable: 'dockerHubUser')]) {
                    sh "docker login -u ${env.dockerHubUser} -p ${env.dockerHubPassword}"
                    echo 'docker login successful'
                    // sh "PROJECT_IMAGE_TAG=\$(git log -1 --pretty=%h)"
                    sh "docker tag ${env.PROJECT} ${env.dockerHubUser}/${env.PROJECT}:${BUILD_NUMBER}"
                    sh "docker push ${env.dockerHubUser}/${env.PROJECT}:${BUILD_NUMBER}"
                }
            }
        }
        stage('Execute Ansible'){
            steps{
                ansiblePlaybook installation: 'Ansible',  inventory: './ansible/inventory', playbook: './ansible/playbook.yaml'
                // sh "ansible-playbook playbook.yaml"
            }
        }
        
        
    }
}