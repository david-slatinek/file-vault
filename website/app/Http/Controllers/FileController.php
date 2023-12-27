<?php

namespace App\Http\Controllers;

use App\Http\Client\File;
use Illuminate\Contracts\Routing\ResponseFactory;
use Illuminate\Foundation\Application;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Response;
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

    public function downloadForm($id): View
    {
        return view("download", ["id" => $id]);
    }

    public function download($id): Application|Response|\Illuminate\Contracts\Foundation\Application|RedirectResponse|ResponseFactory
    {
        $password = request()->input("password");
        $code = request()->input("code");

        [$data, $contentType, $contentDisposition, $err] = File::downloadFile($id, $password, $code);

        if ($err != null) {
            return redirect()->route("error.error")->with("error", $err);
        }

        return response($data)->header("Content-Type", $contentType)
            ->header("Content-Disposition", $contentDisposition);
    }
}
