<?php

namespace App\Http\Controllers;

use App\Http\Client\User;
use Exception;
use Illuminate\Routing\Controller as BaseController;
use Laravel\Socialite\Facades\Socialite;
use Symfony\Component\HttpFoundation\RedirectResponse;

class GithubController extends BaseController
{
    public function redirect(): RedirectResponse
    {
        return Socialite::driver("github")->redirect();
    }

    public function callback()
    {
        try {
            $user = Socialite::driver("github")->user();
        } catch (Exception) {
            return redirect()->route("error.error")->with("error", "You must allow the application to access your
                GitHub account.");
        }

        $token = $user->token;
        session(["token" => $token]);

        [$otp, $err] = User::register();

        if ($err == "setup") {
            return redirect()->route("user.setup")->with("otp", $otp);
        }

        if ($err == "file") {
            return redirect()->route("user.file");
        }

        return redirect()->route("error.error")->with("error", $err);
    }
}
