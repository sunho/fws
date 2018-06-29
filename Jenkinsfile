#!/usr/bin/env groovy
import hudson.model.*

pipeline {
	agent any

	stages {
		stage("Build") {
			steps {
				script {
					withKubeConfig([credentialsId: 'kube-token', serverUrl: 'https://kubernetes.default']) {
						sh 'kubectl get pods'
					}
				}
			}
		}
	}
}

//sh "${tool name: 'sbt', type: 'org.jvnet.hudson.plugins.SbtPluginBuilder$SbtInstallation'}/bin/sbt compile test docker"
