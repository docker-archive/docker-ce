#!groovy

def branch = env.CHANGE_TARGET ?: env.BRANCH_NAME

test_steps = [
	'deb': { ->
		stage('Ubuntu Xenial and Focal Package') {
			wrappedNode(label: 'ubuntu && x86_64', cleanWorkspace: true) {
				try {
					checkout scm
					sh "make REF=$branch checkout"
					sh "make -C deb ubuntu-xenial ubuntu-focal"
				} finally {
					sh "make clean"
				}
			}
		}
	},
	'rpm': { ->
		stage('Centos 7 and 8 RPM Packages') {
			wrappedNode(label: 'ubuntu && x86_64', cleanWorkspace: true) {
				try {
					checkout scm
					sh "make REF=$branch checkout"
					sh "make -C rpm centos-7 centos-8"
				} finally {
					sh "make clean"
				}
			}
		}
	},
	'static': { ->
		stage('Static Linux Binaries') {
			wrappedNode(label: 'ubuntu && x86_64', cleanWorkspace: true) {
				try {
					checkout scm
					sh "make REF=$branch checkout"
					sh "make REF=$branch DOCKER_BUILD_PKGS=static-linux static"
				} finally {
					sh "make clean"
				}
			}
		}
	},
]

parallel(test_steps)
