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
pass-man add <email> <password>
pass-man a <email> <password>  # Short alias
```

**Example:**
```bash
pass-man add john@example.com mySecurePassword123
```

### Retrieve a Password

Get the password for a specific email address (copies to clipboard):

```bash
pass-man get <email>
pass-man g <email>  # Short alias
```

**Example:**
```bash
pass-man get john@example.com
```

### Get All Passwords for a Domain

Retrieve all passwords for a specific domain (copies to clipboard):

```bash
pass-man get-all <domain>
pass-man ga <domain>  # Short alias
```

**Example:**
```bash
pass-man get-all example
```

### Help

Display help information:

```bash
pass-man --help
pass-man <command> --help
```