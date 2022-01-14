# Journalctl proxy

This app exposes your systemd logs to web via web interface.

- Serves as a proxy that reroutes new messages on journalctl to web interface.
- When you load the page all available running services are listed in dropdown.
- You can switch between them and previous logs are still being preserved.
- Once you switch to another service you will stop receiving updates from the previous one.
- **[FORK]** You can now select multiple services
- **[FORK]** ANSI escape sequences also work
- **[FORK]** Docker containers are services as well (when using the journald logging driver)

![Screenshot](https://user-images.githubusercontent.com/5559994/149535084-a3c894a9-2abd-40e5-8149-5f2480e56f7f.png)

## Usage

There are two prebuild binaries available for ARM and AMD64 under [release tab](https://github.com/mitjafelicijan/journalctl-proxy/releases).

Once you unzip downloaded application you can set port on which the server is running at'

```sh
$ ./journalctl-proxy -help
Usage of ./journalctl-proxy:
  -p int
    	Server port number (default 8000)

$ ./journalctl-proxy -p 8000

```
