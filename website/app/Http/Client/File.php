<?php

namespace App\Http\Client;

use App\Models\FileModel;
use Illuminate\Support\Facades\Http;

class File
{
    public static function getFiles(): array
    {
        $response = Http::withoutVerifying()->withToken(session("token"))->get(env("BASE_URL") . "/files");

        if ($response->ok()) {
            $data = $response->json()["files"];

            $files = [];

            foreach ($data as $file) {
                $fileModel = new FileModel();
                $fileModel->set($file);
                $files[] = $fileModel;
            }

            return [$files, null];
        }

        if ($response->noContent()) {
            return [[], null];
        }

        return [[], $response->status() . " - " . $response->body()];
    }
}
