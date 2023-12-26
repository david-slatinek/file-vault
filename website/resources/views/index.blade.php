<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <script src="https://cdn.tailwindcss.com"></script>

    <title>File Vault</title>
</head>

<body class="flex items-center justify-center h-screen bg-gradient-to-r from-slate-400 to-sky-400">

<div class="bg-white p-8 rounded shadow-md w-96">
    <h1 class="text-3xl font-semibold mb-6 text-center text-gray-800">Login with Github</h1>

    <div class="flex items-center justify-center">
        <button class="bg-gradient-to-r from-blue-500 to-red-400 hover:from-blue-600 hover:to-teal-500 text-white
            font-bold py-3 px-10 rounded-full focus:outline-none focus:shadow-outline-blue active:from-blue-800
            active:to-teal-700">
            Login
        </button>
    </div>
</div>

</body>

</html>
