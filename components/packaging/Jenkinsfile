#!groovy


def genBranch(String arch) {
	return [
		"${arch}": { ->
			stage("Build engine image on ${arch}") {
				wrappedNode(label: "linux&&${arch}", cleanWorkspace: true) {
					try {
						checkout scm
						sh("git clone https://github.com/docker/engine.git engine")
						sh('make ENGINE_DIR=$(pwd)/engine image')
					} finally {
						sh('make ENGINE_DIR=$(pwd)/engine clean-image clean-engine')
					}
				}
		}
	}]
}

def branch = env.CHANGE_TARGET ?: env.BRANCH_NAME

test_steps = [
	'deb': { ->
		stage('Ubuntu Xenial Debian Package') {
			wrappedNode(label: 'ubuntu && x86_64', cleanWorkspace: true) {
				checkout scm
				sh('git clone https://github.com/docker/cli.git')
				sh("git -C cli checkout $branch")
				sh('git clone https://github.com/docker/engine.git')
				sh("git -C engine checkout $branch")
				sh('make -C deb VERSION=0.0.1-dev ENGINE_DIR=$(pwd)/engine CLI_DIR=$(pwd)/cli ubuntu-xenial')
			}
		}
	},
	'rpm': { ->
		stage('Centos 7 RPM Package') {
			wrappedNode(label: 'ubuntu && x86_64', cleanWorkspace: true) {
				checkout scm
				sh('git clone https://github.com/docker/cli.git')
				sh("git -C cli checkout $branch")
				sh('git clone https://github.com/docker/engine.git')
				sh("git -C engine checkout $branch")
				sh('make -C rpm VERSION=0.0.1-dev ENGINE_DIR=$(pwd)/engine CLI_DIR=$(pwd)/cli centos-7')
			}
		}
	},
	'static': { ->
		stage('Static Linux Binaries') {
			wrappedNode(label: 'ubuntu && x86_64', cleanWorkspace: true) {
				checkout scm
				sh('git clone https://github.com/docker/cli.git')
				sh("git -C cli checkout $branch")
				sh('git clone https://github.com/docker/engine.git')
				sh("git -C engine checkout $branch")
				sh('make VERSION=0.0.1-dev DOCKER_BUILD_PKGS=static-linux ENGINE_DIR=$(pwd)/engine CLI_DIR=$(pwd)/cli static')
			}
		}
	},
]

arches = [
	"x86_64",
	"aarch64",
	"armhf"
]

arches.each {
	test_steps << genBranch(it)
}

parallel(test_steps)
