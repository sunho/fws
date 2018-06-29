#!/usr/bin/env groovy
import hudson.model.*

pipeline {
	agent any

	stages {
		stage("Build & Test") {
			steps {
				kubernetesDeploy configs: 'deployments/deploy.yaml', kubeconfigId: 'kube-config', enableConfigSubstitution: true
				sh "${tool name: 'sbt', type: 'org.jvnet.hudson.plugins.SbtPluginBuilder$SbtInstallation'}/bin/sbt compile test docker"
			}
		}

		stage("Deploy") {
			steps {
				script {
					docker.withRegistry("", "docker-credentials") {
						sh "docker push ksunhokim/bot-registry:latest"
					}
				}

			}
		}
	}
}
