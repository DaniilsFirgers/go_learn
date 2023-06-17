### GOLANG FULL-STACK MOCK PROJECT

_1. Visual Studio Set Up_

- Cmd+Shift+P and run `Go: Install/Update Tools`, then select all and click ok (all extensions are up)

_2. Enable rootless docker and run docker_

- `dockerd-rootless-setuptool.sh install`

_3. Build container_

- `docker-compose up -d` to run in detached mode
- `docker-compose up --build` to build a container

_4. Start the container_

- `docker-compose start`

_5. Stop the container_

- `docker-compose stop`
