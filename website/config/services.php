<?php

return [
    "github" => [
        "client_id" => env("GITHUB_CLIENT_ID"),
        "client_secret" => env("GITHUB_CLIENT_SECRET"),
        "redirect" => env("GITHUB_REDIRECT_URI"),
        "scopes" => ["user:email"],
        "state" => env("GITHUB_STATE"),
    ],
];
