#!groovy

def branch = env.CHANGE_TARGET ?: env.BRANCH_NAME

test_steps = [
	'deb': { ->
		stage('Ubuntu Xenial Debian Package') {
			wrappedNode(label: 'ubuntu && x86_64', cleanWorkspace: true) {
				try {
					checkout scm
					sh('git clone https://github.com/docker/cli.git')
					sh("git -C cli checkout $branch")
					sh('git clone https://github.com/docker/engine.git')
					sh("git -C engine checkout $branch")
					sh('make -C deb VERSION=0.0.1-dev ENGINE_DIR=$(pwd)/engine CLI_DIR=$(pwd)/cli ubuntu-xenial')
				} finally {
					sh('make ENGINE_DIR=$(pwd)/engine clean-engine')
				}
			}
		}
	},
	'rpm': { ->
		stage('Centos 7 RPM Package') {
			wrappedNode(label: 'ubuntu && x86_64', cleanWorkspace: true) {
				try {
					checkout scm
					sh('git clone https://github.com/docker/cli.git')
					sh("git -C cli checkout $branch")
					sh('git clone https://github.com/docker/engine.git')
					sh("git -C engine checkout $branch")
					sh('make -C rpm VERSION=0.0.1-dev ENGINE_DIR=$(pwd)/engine CLI_DIR=$(pwd)/cli centos-7')
				} finally {
					sh('make ENGINE_DIR=$(pwd)/engine clean-engine')
				}
			}
		}
	},
	'static': { ->
		stage('Static Linux Binaries') {
			wrappedNode(label: 'ubuntu && x86_64', cleanWorkspace: true) {
				try {
					checkout scm
					sh('git clone https://github.com/docker/cli.git')
					sh("git -C cli checkout $branch")
					sh('git clone https://github.com/docker/engine.git')
					sh("git -C engine checkout $branch")
					sh('make VERSION=0.0.1-dev DOCKER_BUILD_PKGS=static-linux ENGINE_DIR=$(pwd)/engine CLI_DIR=$(pwd)/cli static')
				} finally {
					sh('make ENGINE_DIR=$(pwd)/engine clean-engine')
				}
			}
		}
	},
]

parallel(test_steps)
