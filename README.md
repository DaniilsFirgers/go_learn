### GOLANG

1. Minimalistic project structure

- Create a cmd folder in the root directory (contains main entry points)
- Run go mod init (to initilize a new Go module)
- Run go get to add new dependencies to the project
- Create an api folder in cmd folder (for JSON schema files, protocol definition files)
- Create examples folder in the root directory (for usage examples)

2. Visual Studio SetUp

- Cmd+Shift+P and run Go: Install/Update Tools, then select all and click ok (all extensions are up)

3. Enable rootless docker and run docker

- `dockerd-rootless-setuptool.sh install`
