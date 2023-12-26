<?php

use App\Http\Controllers\UserController;
use Illuminate\Support\Facades\Route;

Route::get("/", [UserController::class, "index"]);

Route::get("setup", [UserController::class, "setup"]);

Route::get("login", [UserController::class, "login"]);

Route::get("home", [UserController::class, "home"]);

Route::get("upload", [UserController::class, "upload"]);
