pipeline{
    agent any DOUBT really ?? 
    sshagent(['ec2-ssh-key']){
        sh """
        ssh -i ${SSH_KEY} ${REMOTE_USER}@${INSTANCE_IP} << EOF 
        docker login -u ${DOCkER_HUB_CREDENTIALS_USR} -p 
        """
    }
    stages{
        stage("Code"){
            steps{
        git url: "https://github.com/YogeshDharya/MessageMate.git", branch:"main"
            }
        }
        stage("Build"){
            steps{
                sh "docker build . -t messageMate"
            }
        }
        stage("Push 2 DockerHub"){
            steps{
                withCredentials([usernamePassword(crednetialsId: "dockerHub",passwordVariable:"dockerHubPassword",usernameVariable:"dockerHubUsername")]){
                    sh "docker login -u ${env.dockerHubUsername} -p ${env.dockerHubPassword}"
                    sh "docker tag messageMate ${env.dockerHubUsername"}/messageMate:latest"
                    sh "docker push ${env.dockerHubUsername}/messageMate:latest"
                }
            }
        }
        stage("Deploy"){
            steps{
                sh "docker-compose down"
            }
        }
    }
}