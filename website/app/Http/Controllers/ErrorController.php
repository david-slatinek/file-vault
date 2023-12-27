<?php

namespace App\Http\Controllers;

use Illuminate\Routing\Controller as BaseController;
use Illuminate\View\View;

class ErrorController extends BaseController
{
    public function error(): View
    {
        return view("error");
    }
}
