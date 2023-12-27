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

    public static function deleteFile(string $id, string $code): string
    {
        $response = Http::withoutVerifying()->withToken(session("token"))->delete(env("BASE_URL")
            . "/delete/" . $id, ["code" => $code]);

        if ($response->noContent()) {
            return "";
        }

        return $response->status() . " - " . $response->body();
    }

    public static function uploadFile(string $file, string $password, string $code): string
    {
        $multipart = [
            [
                "name" => "file",
                "contents" => fopen($file, "r")
            ],
            [
                "name" => "password",
                "contents" => $password
            ],
            [
                "name" => "code",
                "contents" => $code
            ]
        ];

        $response = Http::withoutVerifying()->withToken(session("token"))->attach($multipart)->post(env("BASE_URL") . "/upload");

        if ($response->created()) {
            return "";
        }

        return $response->status() . " - " . $response->body();
    }
}
