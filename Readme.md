# Vaulta

Vaulta is a secure terminal based secret manager. It encrypts your data using AES-256-GCM and stores it to a local JSON file on your computer.

## Installation

### Using Go Install

If you have Go installed, you can install vaulta directly:

```bash
go install github.com/azizrmadi/vaulta@latest
```

### Download Binary from GitHub Releases

You can download pre-built binaries from the [GitHub Releases](https://github.com/azizrmadi/vaulta/releases) page.

1. Go to the [releases page](https://github.com/azizrmadi/vaulta/releases)
2. Download the appropriate binary for your operating system and architecture
3. Extract the archive and move the binary to a directory in your `PATH`

### Using curl

#### macOS (Apple Silicon)

```bash
curl -Lo vaulta.tar.gz https://github.com/azizrmadi/vaulta/releases/latest/download/vaulta_Darwin_arm64.tar.gz
tar -xzf vaulta.tar.gz
sudo mv vaulta /usr/local/bin/
```

#### macOS (Intel)

```bash
curl -Lo vaulta.tar.gz https://github.com/azizrmadi/vaulta/releases/latest/download/vaulta_Darwin_x86_64.tar.gz
tar -xzf vaulta.tar.gz
sudo mv vaulta /usr/local/bin/
```

#### Linux (amd64)

```bash
curl -Lo vaulta.tar.gz https://github.com/azizrmadi/vaulta/releases/latest/download/vaulta_Linux_x86_64.tar.gz
tar -xzf vaulta.tar.gz
sudo mv vaulta /usr/local/bin/
```

#### Linux (arm64)

```bash
curl -Lo vaulta.tar.gz https://github.com/azizrmadi/vaulta/releases/latest/download/vaulta_Linux_arm64.tar.gz
tar -xzf vaulta.tar.gz
sudo mv vaulta /usr/local/bin/
```

## Usage

To start using vaulta, run the following command.

```shell
vaulta init
```

Follow the prompts to initialize your password.

### Available Commands

Below is a list of the currently available commands in vaulta (more features to come hopefully)

#### Add Entries

To add a new entry in the vault, run:

```bash
vaulta add
```

Then follow the prompts to create the entry.

#### Delete Entries

To delete an entry in the vault, run:

```bash
vaulta delete <entry>
```

#### List Entries

To list all entries in the vault, run:

```bash
vaulta list
```

#### Get Entries

To get an entry from the vault, run:

```bash
vaulta get <entry>
```
