<?php

namespace App\Http\Controllers;

use Illuminate\Contracts\View\View;
use Illuminate\Routing\Controller as BaseController;

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

    public function login(): View
    {
        return view("login");
    }

    public function home(): View
    {
        return view("home");
    }

    public function upload(): View
    {
        return view("upload");
    }
}
