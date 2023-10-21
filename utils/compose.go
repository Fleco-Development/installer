package utils

import "os/exec"

func GetCompose() string {

	err := exec.Command("docker-compose", "--version").Run()
	if err != nil {

		err = exec.Command("docker", "compose", "--version").Run()
		if err != nil {
			return ""
		}

		return "docker compose"

	}

	return "docker-compose"

}
