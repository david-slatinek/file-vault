<?php

namespace App\Models;

class FileModel
{
    public string $id;
    public string $filename;
    public string $createdAt;
    public string $accessedAt;

    public function set($data): void
    {
        foreach ($data as $key => $value) {
            $this->{$key} = $value;
        }
    }
}
