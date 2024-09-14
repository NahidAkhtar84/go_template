 # Run Project:
 
- Run `go mod tidy` in project rood directory.
- Create `wait-for-it.sh` file if not present. 
  - Command: 
  ```
  curl -o wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
  chmod +x wait-for-it.sh
  ```
- Then Run `docker compose up --build`


# Description:



# Package and Script Description:

## Wait for it sh file:

This shell script (wait-for-it.sh) is used to test if a TCP host and port are available before proceeding with additional commands. Here's a short breakdown of what it does:

Host and Port Check: The script accepts a host:port combination and waits until the specified TCP port on the given host becomes available (i.e., the server is ready to accept connections).

Timeout: You can specify a timeout for how long to wait. If the port doesn't become available within the specified time, it exits with an error.

Optional Command Execution: Once the host and port are available, the script can optionally execute a provided command and its arguments. If strict mode (-s) is enabled, the command only runs if the host and port are successfully reached.

Quiet Mode: The script supports quiet mode (-q), where it suppresses status messages.

Handling Busybox: It includes logic to check whether it's running on a BusyBox system and adjusts the timeout flag accordingly.

Use Case:
It is commonly used in Docker environments to ensure that dependent services (like a database) are up and running before starting another service or executing further commands.

## go mod tidy
The go mod tidy command is used to clean up the go.mod and go.sum files in your Go project. Specifically, it:

Removes Unused Dependencies: It removes any dependencies that are listed in go.mod but are not actually used in the code.
Adds Missing Dependencies: It ensures that all the necessary dependencies for the current code are present in go.mod and go.sum.
Updates go.sum: It ensures that the go.sum file contains checksums for the dependencies used in the project, which ensures the integrity of the modules you're using.


## go.mod: Module Definition File
The go.mod file defines the module (project) and tracks its dependencies. It is like a 
configuration file for managing the versions of external Go modules your project depends on.

## go.sum: Dependency Checksum File
The go.sum file ensures integrity and security by storing checksums (hashes) of the module 
versions used by your project. It prevents unwanted changes to dependencies (e.g., malicious 
updates or code tampering).

## Extras:

- "--": The double dash (--) is used to indicate that the arguments following it are for the command that will be run after the script finishes waiting. This tells the script that everything after -- is a separate command.
