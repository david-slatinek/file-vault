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
    public function home(): View
    {
        return view("home");
    }

    public function setup(): View
    {
        $otp = session("otp");
        return view("setup", ["otp" => $otp]);
    }

    public function file(): Application|RedirectResponse|Redirector|View
    {
        [$files, $err] = File::getFiles();

        if ($err != null) {
            return redirect()->route("error.error")->with("error", $err);
        }

        return view("file", ["files" => $files]);
    }

    public function upload(): View
    {
        return view("upload");
    }
}
