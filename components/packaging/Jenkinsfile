#!groovy

def stageWithGithubNotify(String name, Closure cl) {
	def notify = { String status ->
		safeGithubNotify([context: name, targetUrl: "${BUILD_URL}/flowGraphTable"] + [status: status])
	}

	stage(name) {
		notify 'PENDING'
		try {
			cl()
		} catch (err) {
			addFailedStage(name)
			notify 'FAILURE'
			throw err
		}
		notify 'SUCCESS'
	}
}

parallel {
	stageWithGithubNotify('Ubuntu Xenial Debian Package') {
		wrappedNode(label: 'docker-edge && x86_64', cleanWorkspace: true) {
			checkout scm
			sh("git clone https://github.com/docker/cli.git")
			sh("git clone https://github.com/moby/moby.git")
			sh("make DOCKER_BUILD_PKGS=ubuntu-xenial ENGINE_DIR=moby CLI_DIR=cli deb")
		}
	},
	stageWithGithubNotify('Centos 7 RPM Package') {
		wrappedNode(label: 'docker-edge && x86_64', cleanWorkspace: true) {
			checkout scm
			sh("git clone https://github.com/docker/cli.git")
			sh("git clone https://github.com/moby/moby.git")
			sh("make DOCKER_BUILD_PKGS=centos-7 ENGINE_DIR=moby CLI_DIR=cli rpm")
		}
	}
	stageWithGithubNotify('Static Linux Binaries') {
		wrappedNode(label: 'docker-edge && x86_64', cleanWorkspace: true) {
			checkout scm
			sh("git clone https://github.com/docker/cli.git")
			sh("git clone https://github.com/moby/moby.git")
			sh("make DOCKER_BUILD_PKGS=static-linux ENGINE_DIR=moby CLI_DIR=cli static")
		}
	}
}
