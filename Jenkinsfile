pipeline {
	agent any
	tools {
		go 'Go'
	}
	environment {
		SONAR_PROJECT_KEY = 'Go-App'
		SONAR_SCANNER_HOME = tool 'SonarQubeScanner'
	}

	stages {
		stage('Checkout Github'){
			steps {
				git([url: 'https://github.com/steveodhiambo/Structy.git', branch: 'main', credentialsId: 'github-cred'])
			}
		}
		
		stage ('Tests'){
            steps{
                sh 'go test -v ./...'
            }
        }
		stage('SonarQube Analysis'){
			steps {
				withCredentials([string(credentialsId: 'node-token', variable: 'SONAR_TOKEN')]) {
				   
					withSonarQubeEnv('SonarQube') {
						sh """
                  				${SONAR_SCANNER_HOME}/bin/sonar-scanner \
                  				-Dsonar.projectKey=${SONAR_PROJECT_KEY} \
                    				-Dsonar.sources=. \
                   				-Dsonar.host.url=http://88.99.190.104:9000/ \
                    				-Dsonar.login=${SONAR_TOKEN}
                    				"""
					}	
				}
			}
		}
	}
	post {
		success {
			echo 'Build completed succesfully!'
		}
		failure {
			echo 'Build failed. Check logs.'
		}
	}
}



