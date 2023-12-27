<?php

use App\Http\Controllers\ErrorController;
use App\Http\Controllers\FileController;
use App\Http\Controllers\GithubController;
use App\Http\Controllers\UserController;
use Illuminate\Support\Facades\Route;

Route::get("/", [UserController::class, "home"])->name("user.home");
Route::get("setup", [UserController::class, "setup"])->name("user.setup");
Route::get("file", [UserController::class, "file"])->name("user.file");
Route::get("upload", [UserController::class, "upload"])->name("user.upload");

Route::delete("delete-form/{id}", [FileController::class, "code"])->name("file.delete-form");
Route::delete("delete/{id}", [FileController::class, "delete"])->name("file.delete");

Route::get("/github/redirect", [GithubController::class, "redirect"])->name("github.redirect");
Route::get("/github/callback", [GithubController::class, "callback"])->name("github.callback");

Route::get("error", [ErrorController::class, "error"])->name("error.error");
