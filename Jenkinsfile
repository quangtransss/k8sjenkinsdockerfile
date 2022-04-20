// import groovy.json.JsonOutput
// def COLER_MAP = [
//     'SUCESS' : 'good',
//     'FAILURE': 'danger',
// ]
// def getBuildUser() {
//     return currentBuild.rawBuild.getCause(Cause.UserIdCause).getUserId()
// }
pipeline{
    agent any
    environment {
        DOCKER_IMAGE = "quangno129/quang"
        APP_NAME="golangbank"
        DOCKER_TAG="${GIT_BRANCH.tokenize('/').pop()}-${GIT_COMMIT.substring(0,7)}"
        // BUILD_USER= ''
    }
    stages{
        stage("build"){
            // slackSend color: '#BADA55', message: 'Hello, Wrld!'
            // environment {
            // DOCKER_TAG="${GIT_BRANCH.tokenize('/').pop()}-${GIT_COMMIT.substring(0,7)}"

            // }
            steps{
                echo "run dockerfile"
                sh "docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG}${BUILD_NUMBER} . "
                // sh "docker tag ${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_IMAGE}:latest"
                sh "docker image ls | grep ${DOCKER_IMAGE}"
                withCredentials([usernamePassword(credentialsId: '7f7df69d-adad-46ff-85dd-8b3ba4ac805a', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                    sh 'echo $DOCKER_PASSWORD | docker login --username $DOCKER_USERNAME --password-stdin'
                    sh "docker push ${DOCKER_IMAGE}:${DOCKER_TAG}${BUILD_NUMBER}"
                }
             //clean to save disk
            sh "docker image rm ${DOCKER_IMAGE}:${DOCKER_TAG}${BUILD_NUMBER}"
            // sh "docker image rm ${DOCKER_IMAGE}:latest"
            }
        }
        stage("deployment"){
            agent { label 'docker-deploy'}
            steps{
                withCredentials([usernamePassword(credentialsId: '	15d40a63-6cfc-49c1-b810-01f5c2b1d1a7', usernameVariable: 'ARGOCD_USERNAME', passwordVariable: 'ARGOCD_PASSWORD')]) {
                    sh "argocd login argocd.medpro.com.vn --username ${ARGOCD_USERNAME} --password ${ARGOCD_PASSWORD}"
                    sh "argocd app set ${APP_NAME}  --kustomize-image  ${DOCKER_IMAGE}:${DOCKER_TAG}${BUILD_NUMBER}"
                    sh "argocd app sync ${APP_NAME}"
                }
            }
        }
    }
    post {
        always {
            //The script step takes a block of Scripted Pipeline and executes that in the Declarative Pipeline. 
            //For most use-cases, the script step should be unnecessary in Declarative Pipelines, but it can provide
            //a useful "escape hatch." script blocks of non-trivial size and/or complexity should be moved into Shared Libraries instead.
            // script {
            //     BUILD_USER = getBuildUser()
            // }
            
            slackSend channel: '#spsdevops',
                // color: COLOR_MAP[currentBuild.currentResult],
                message: "*${currentBuild.currentResult}:* Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}"
        }
    }
}
