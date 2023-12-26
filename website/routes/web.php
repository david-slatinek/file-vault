<?php

use App\Http\Controllers\GithubController;
use App\Http\Controllers\UserController;
use Illuminate\Support\Facades\Route;

Route::get("/", [UserController::class, "index"])->name("/");

Route::get("setup", [UserController::class, "setup"])->name("setup");

//Route::get("login", [UserController::class, "login"]);

Route::get("home", [UserController::class, "home"])->name("home");

Route::get("upload", [UserController::class, "upload"])->name("upload");

Route::get("/github/redirect", [GithubController::class, "redirect"])->name("github.redirect");

Route::get("/github/callback", [GithubController::class, "callback"])->name("github.callback");
