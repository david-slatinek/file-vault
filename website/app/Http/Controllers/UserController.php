<?php

namespace App\Http\Controllers;

use App\Http\Client\File;
use Illuminate\Contracts\Foundation\Application;
use Illuminate\Contracts\View\View;
use Illuminate\Http\RedirectResponse;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Routing\Redirector;

class UserController extends BaseController
{
    public function index(): View
    {
        return view("index");
    }

    public function setup(): View
    {
        $otp = session("otp");
        return view("setup", ["otp" => $otp]);
    }

//    public function login(): View
//    {
//        return view("login");
//    }

    public function home(): Application|RedirectResponse|Redirector|View
    {
        [$files, $err] = File::getFiles();

        if ($err != null) {
            return redirect("/")->with("error", $err);
        }

        return view("home", ["files" => $files]);
    }

    public function upload(): View
    {
        return view("upload");
    }
}
