#!groovy

test_steps = [
	'deb': { ->
		stage('Ubuntu Xenial Debian Package') {
			wrappedNode(label: 'docker-edge && x86_64', cleanWorkspace: true) {
				checkout scm
				sh("git clone https://github.com/docker/cli.git")
				sh("git clone https://github.com/moby/moby.git")
				sh("make DOCKER_BUILD_PKGS=ubuntu-xenial ENGINE_DIR=moby CLI_DIR=cli deb")
			}
		}
	},
	'rpm': { ->
		stage('Centos 7 RPM Package') {
			wrappedNode(label: 'docker-edge && x86_64', cleanWorkspace: true) {
				checkout scm
				sh("git clone https://github.com/docker/cli.git")
				sh("git clone https://github.com/moby/moby.git")
				sh("make DOCKER_BUILD_PKGS=centos-7 ENGINE_DIR=moby CLI_DIR=cli rpm")
			}
		}
	},
	'static': { ->
		stage('Static Linux Binaries') {
			wrappedNode(label: 'docker-edge && x86_64', cleanWorkspace: true) {
				checkout scm
				sh("git clone https://github.com/docker/cli.git")
				sh("git clone https://github.com/moby/moby.git")
				sh("make DOCKER_BUILD_PKGS=static-linux ENGINE_DIR=moby CLI_DIR=cli static")
			}
		}
	},
]

parallel(test_steps)
