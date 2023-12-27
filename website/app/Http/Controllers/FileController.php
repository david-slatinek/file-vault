<?php

namespace App\Http\Controllers;

use App\Http\Client\File;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\View\View;

class FileController extends BaseController
{
    public function deleteForm($id): View
    {
        return view("delete", ["id" => $id]);
    }

    public function delete($id)
    {
        $code = request()->input("code");

        $err = File::deleteFile($id, $code);

        if ($err !== "") {
            return redirect()->route("error.error")->with("error", $err);
        }

        return redirect()->route("user.file");
    }

    public function uploadForm(): View
    {
        return view("upload");
    }

    public function upload()
    {
        $file = request()->file("file");
        $password = request()->input("password");
        $code = request()->input("code");

        $path = $file->storeAs("images", $file->getClientOriginalName());

        if ($path === false) {
            return redirect()->route("error.error")->with("error", "Failed to upload file");
        }

        $err = File::uploadFile(storage_path("app/" . $path), $password, $code);

        unlink(storage_path("app/" . $path));

        if ($err !== "") {
            return redirect()->route("error.error")->with("error", $err);
        }

        return redirect()->route("user.file");
    }
}
