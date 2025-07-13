# Pass-Man 🔐

A secure command-line password manager built with Go that stores encrypted passwords locally using BoltDB.

## Features

- **Secure Storage**: Passwords are encrypted using AES encryption before storage
- **Domain-based Organization**: Passwords are organized by domain extracted from email addresses
- **Clipboard Integration**: Automatically copies passwords to clipboard for easy use
- **Simple CLI Interface**: Easy-to-use command-line interface with intuitive commands
- **Local Database**: All data stored locally in an encrypted BoltDB database

## Prerequisites

- `xclip` (for Linux clipboard support)

## Installation

You just need to run this command and enjoy! 

```
curl -sL https://raw.githubusercontent.com/PandaX185/pass-man/refs/heads/master/install.sh | bash
```

## Usage

### Add a Password

Add or update a password for a specific email address:

```bash
passman add <email> <password>
passman a <email> <password>  # Short alias
```

**Example:**
```bash
passman add john@example.com mySecurePassword123
```

### Retrieve a Password

Get the password for a specific email address (copies to clipboard):

```bash
passman get <email>
passman g <email>  # Short alias
```

**Example:**
```bash
passman get john@example.com
```

### Get All Passwords for a Domain

Retrieve all passwords for a specific domain (copies to clipboard):

```bash
passman get-all <domain>
passman ga <domain>  # Short alias
```

**Example:**
```bash
passman get-all example
```

### Help

Display help information:

```bash
passman --help
passman <command> --help
```