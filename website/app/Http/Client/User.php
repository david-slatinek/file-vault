<?php

namespace App\Http\Client;

use App\Models\Otp;
use Illuminate\Support\Facades\Http;

class User
{
    public static function register(): array
    {
        $response = Http::withoutVerifying()->withToken(session("token"))->post(env("BASE_URL") . "/register");

        if ($response->created()) {
            $otp = new Otp();
            $otp->set($response->json());
            return [$otp, ""];
        }

        if ($response->badRequest()) {
            return [null, "home"];
        }

        return [null, $response->status() . " - " . $response->body()];
    }
}
