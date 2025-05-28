pipeline {
    agent any

    /*Actions{        -> Not actions its Options
        onGitPush()   -> githubPush()
    }*/

    Options{
        // to trigger the webhook
        githubPush()
    }

    stages{
        // steps{                                             -> It's stage() and need to name the stage inside the braces
        stage('Checkout code - Cloning the github repo'){
            //steps("Cloning the main repo"){                 -> It's just steps{} no comment  
            steps{
                //Git.repo("https://github.com/OneClickCoder/tracker_pipeline")        -> WRONG Syntax
                //need to give ssh credentials - using the cred ID created in Jenkins   
                git branch: 'main'
                    credentialsId:'github-ssh-key'
                    url: 'git@github.com/OneClickCoder/tracker_pipeline.git' // need to mention .git in the end of the URL
                
                echo "GitHub repo cloned Sucessfully"
                sh "ls -la"

            }
        }
    }

    post{
        //this will run even if the above stages pass or fail
        // so can place actions need to be taken after the build is completed and app is deployed
        always {
            cleanWs() // clean Workspace this will clean the Working Dir after an pipeline is run 
        }
        success {
            echo "+++++++ Mission Passed : Repect+ \n Pipeline finished successfully +++++++"
        }
        failure {
            echo "Pipeline failed!"
        }
    }
}