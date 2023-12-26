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
            return redirect("/")->with("error", "You must allow the application to access your GitHub account.");
        }

        $token = $user->token;
        session(["token" => $token]);

        [$otp, $err] = User::register();

        if ($otp != null) {
            return redirect()->route("setup")->with("otp", $otp);
        }

        if ($err == "home") {
            return redirect()->route("home");
        }

        return redirect("/")->with("error", $err);
    }
}
