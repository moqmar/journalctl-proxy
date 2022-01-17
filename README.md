# Journalctl proxy

This app exposes your systemd logs to web via web interface.

- Serves as a proxy that reroutes new messages on journalctl to web interface.
- When you load the page all available running services are listed in dropdown.
- You can switch between them and previous logs are still being preserved.
- Once you switch to another service you will stop receiving updates from the previous one.
- **[FORK]** You can now select multiple services
- **[FORK]** ANSI escape sequences also work
- **[FORK]** Docker containers are services as well (when using the journald logging driver)
- **[FORK]** Future roadmap:
  - Extensive filter functionality with autocomplete etc.
  - Views (saved filters), accessible via REST API and in the browser (through WebSockets)

![Screenshot](https://user-images.githubusercontent.com/5559994/149535084-a3c894a9-2abd-40e5-8149-5f2480e56f7f.png)

## Usage

There are two prebuilt binaries available for Linux/ARM and Linux/AMD64 under the [release tab](https://github.com/moqmar/journalctl-proxy/releases).

Once you unzip the downloaded application you can set port on which the server is running at:

```
$ ./journalctl-proxy -help
Usage of ./journalctl-proxy:
  -ap string
        Password for basic auth
  -au string
        Username for basic auth
  -docker
        Add container names for Docker scopes (with journald logging driver)
  -p int
        Server port number (default 8000)
$ ./journalctl-proxy -p 8000
```
