# Journal Viewer

A web-based exploration tools for your systemd-journald logs. Aggregate your logs, the systemd way!

## Features

- View journald logs from a web browser
- Easy scrolling through your timeline
- Advanced search to match field values, e.g. `(source = test_noop_1.docker) (level < warn) !"irrelevant log message"`
- Filter by host & source unit
- Create views for commonly used searches
- Advanced authentication through LDAP or OAuth2 (or just use a text file to get started)
- External Notifiers, works with Nagios & CheckMK
- Fully automatic HTTPS setup with Let's Encrypt
- Server-Sent-Events for great performance

![Screenshot](https://user-images.githubusercontent.com/5559994/149535084-a3c894a9-2abd-40e5-8149-5f2480e56f7f.png)

## Roadmap

- [ ] Create web component for the search field
- [ ] Create web component for a scrolling graph
- [ ] Create basic Vue 3 app
- [ ] Implement basic API with authentication
- [ ] Make journalctl access more efficient
- [ ] Implement API for views (personal & global?)
- [ ] Implement UI for views
- [ ] UI polish
- [ ] Package & distribute
- [ ] Build standalone Docker container
- [ ] Build external notifier

## Why?

There are many log aggregation tools out there - Loki, Greylog, the ELK stack and quite some more.
They all have in common that they require you to set up additional infrastructure dedicated to logging, with each server passing their logs to the log storage through one of many protocols.
This gives you a lot of flexibility and scalability, for the cost of having to learn and set up complex and performance-hungry new things.

Journal Viewer is the opposite: it only works with systemd-journald, and in most cases, all your logs are already there.
You can use systemd-journal-upload to aggregate logs to a different system - it supports TLS-based authentication, as well as even forward secure sealing for audit purposes.

Especially for just a few servers, the goal is to be everything you need, while being as simple to use as possible.

The project is originally a fork of [journalctl-proxy](https://github.com/mitjafelicijan/journalctl-proxy), with lots of UI/UX improvements.

## Installation

```shell
# Debian/Ubuntu-based
add-apt-repository ppa:momar/systemd-journal-viewer
apt install systemd-journal-viewer
$EDITOR /etc/defaults/systemd-journal-viewer
systemctl enable --now systemd-journal-viewer.service

# RHEL/CentOS-based
# TODO: repository?
dnf install systemd-journal-viewer
$EDITOR /etc/defaults/systemd-journal-viewer
systemctl enable --now systemd-journal-viewer.service

# Manual installation
wget -O /usr/local/bin/systemd-journal-viewer https://codeberg.org/momar/systemd-journal-viewer/...
wget -O /etc/systemd/system/systemd-journal-viewer.service https://codeberg.org/momar/systemd-journal-viewer/...
wget -O /etc/default/systemd-journal-viewer https://codeberg.org/momar/systemd-journal-viewer/...
$EDITOR /etc/defaults/systemd-journal-viewer
systemctl enable --now systemd-journal-viewer.service

# As a standalone log aggregation container with Docker
docker run -d -v "/var/lib/journal-aggregator" docker.io/momar/systemd-journal-aggregator
```

The Docker container is intentionally a whole own systemd-journald server (and thus called `systemd-journal-aggregator`) for the sole purpose of having other servers upload their journal there.
As it requires quite low-level access to systemd, the viewer alone can not be used with Docker right now.

## Configuration
All configuration is done through environment variables, which by default are loaded from `/etc/default/systemd-journal-viewer`:

```shell
# Listen Settings
HOST=[::]
PORT=19533

# Let's Encrypt Settings
LETSENCRYPT_DOMAIN=logs.example.org
LETSENCRYPT_ACCEPT_TOS=yes

# Behaviour Settings
TRUST_SOURCE=yes

# Authentication Settings
AUTH=htpassed
AUTH_FILE=/etc/sytemd-journal-viewer.htpasswd
#AUTH=ldap
#TODO
#AUTH=oauth2
#TODO
```

The `AUTH_FILE` is a simple htpasswd file - you can generate a user/password line for that using `docker run --rm -it httpd htpasswd -n USERNAME`.

## Aggregate logs from other systems

TODO

## Using the external notifier

TODO