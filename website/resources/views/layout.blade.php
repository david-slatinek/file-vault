<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <script src="https://cdn.tailwindcss.com"></script>

    <title>{{$title}}</title>
</head>

<nav class="bg-slate-500 p-4 fixed w-full top-0">
    <div class="container mx-auto flex justify-between items-center">
        <a href="{{route("user.home")}}" class="text-white text-lg font-bold">Home</a>

        <div class="flex space-x-4">
            <a href="{{route("user.file")}}" class="text-white mr-5">Files</a>
            <a href="{{route("file.upload-form")}}" class="text-white">Upload</a>
        </div>
    </div>
</nav>

@yield("content")

</html>
