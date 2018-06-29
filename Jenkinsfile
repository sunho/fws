#!/usr/bin/env groovy
import hudson.model.*

pipeline {
	agent any

	stages {
		stage("Build") {
			steps {
				sh "${tool name: 'sbt', type: 'org.jvnet.hudson.plugins.SbtPluginBuilder$SbtInstallation'}/bin/sbt compile test"
			}
		}
	}
}
