def git_url = 'https://github.com/oversampling/jenkins-demo.git'
def swr_login = 'docker login -u ${SWR_REGION}@$SWR_ACCESS_KEY -p $SWR_SECRET_KEY swr.${SWR_REGION}.myhuaweicloud.com'
def build_name = 'jenkins-demo'

pipeline {
    agent any
    environment {
        SWR_ACCESS_KEY = credentials("SWR_ACCESS_KEY")
        SWR_SECRET_KEY = credentials("SWR_SECRET_KEY")
        ORGANIZATION = credentials("ORGANIZATION")
        SWR_REGION = credentials("SWR_REGION")
    }
    stages {
        stage('Clone') { 
            steps{
                echo "1.Clone Stage" 
                git url: git_url
                script { 
                    build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim() 
                } 
            }
        } 
        stage('Test') { 
            steps{
                echo "2.Test Stage" 
            }
        } 
        stage('Build') { 
            steps{
                echo "3.Build Docker Image Stage" 
                sh "docker build -t swr.${SWR_REGION}.myhuaweicloud.com/${ORGANIZATION}/${build_name}:${BUILD_NUMBER} ." 
            }
        } 
        stage('Push') { 
            steps{
                echo "4.Push Docker Image Stage" 
                sh swr_login
                sh "docker push swr.${SWR_REGION}.myhuaweicloud.com/${ORGANIZATION}/${build_name}:${BUILD_NUMBER}" 
            }
        }
        stage('Update Deployment Image') {
            steps {
                script {
                    sh "sed -i 's|image: .*|image: swr.${SWR_REGION}.myhuaweicloud.com/${ORGANIZATION}/${build_name}:${BUILD_NUMBER}|' k8s.yaml"
                }
            }
        }
        stage('Deploy') {
            steps{
                echo "5. Deploy Stage"
                echo "This is a deploy step to test"                
                sh "cat k8s.yaml"
                sh "kubectl apply -f k8s.yaml -n jenkins"
            }
        }
    }
}
