<div align="center">
    <img alt="Go" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"/>
    <img alt="Laravel" src="https://img.shields.io/badge/LAravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white"/>
    <img alt="Docker" src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white"/>
    <img alt="Amazon S3" src="https://img.shields.io/badge/amazon_s3-569A31?style=for-the-badge&logo=amazons3&logoColor=white"/>
    <img alt="PostgreSQL" src="https://img.shields.io/badge/postgresql-4169E1?style=for-the-badge&logo=postgresql&logoColor=white"/>
    <img alt="GNU Bash" src="https://img.shields.io/badge/gnu_bash-4EAA25?style=for-the-badge&logo=gnubash&logoColor=white"/>
    <img alt="tailwindcss" src="https://img.shields.io/badge/tailwind_css-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white"/>
</div>

## Description

File Vault is a simple, secure file storage service that allows you to upload, download, and delete files.

## API

[API](api) was made using Go, along with the PostgreSQL [database](api/db). Files are stored
in [Localstack S3](api/storage) and are encrypted using AES-256-GCM. For deriving the encryption key from the user
password, Argon2 is used. We additionally added [2FA](api/otp) using TOTP. The 2FA secret key is encrypted using
[RSA-4096](api/pki). We also added our own [TLS certificate](api/cert).

## Website

The [website](website) was made using Laravel. It is a simple web interface for the File Vault API. For login, we used
Github OAuth.

<div align="center">
  <img src="images/home.png" alt="Home page">
  <br/>
  <i>Home page.</i>
</div>

For additional images, see the [images](images) folder.
