<?php

namespace App\Models;

class Otp
{
    public string $key;
    public string $url;

    public function set($data): void
    {
        foreach ($data as $key => $value) {
            $this->{$key} = $value;
        }
    }
}
