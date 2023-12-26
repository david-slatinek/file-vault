<?php

namespace App\Http\Controllers;

use App\Http\Client\File;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\View\View;

class FileController extends BaseController
{
    public function code($id): View
    {
        return view("delete", ["id" => $id]);
    }

    public function delete($id)
    {
        $code = request()->input("code");

        $err = File::deleteFile($id, $code);

        if ($err !== "") {
            return redirect("error")->with("error", $err);

//            return view("error", ["error" => $err]);
        }

        return redirect("/home");
    }
}
