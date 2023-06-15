### GOLANG

_1. Minimalistic project structure_

- Create a cmd folder in the root directory (contains main entry points)
- Run go mod init (to initilize a new Go module)
- Run go get to add new dependencies to the project
- Create an api folder in cmd folder (for JSON schema files, protocol definition files)
- Create examples folder in the root directory (for usage examples)

_2. Visual Studio Set Up_

- Cmd+Shift+P and run `Go: Install/Update Tools`, then select all and click ok (all extensions are up)

_3. Enable rootless docker and run docker_

- `dockerd-rootless-setuptool.sh install`

_4. Run docker conatiner_

- `docker-compose up -d`
