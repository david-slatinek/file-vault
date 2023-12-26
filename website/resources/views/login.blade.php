<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <script src="https://cdn.tailwindcss.com"></script>

    <title>Login</title>
</head>

<body class="flex items-center justify-center h-screen bg-gradient-to-r from-violet-300 to-cyan-950">

<div class="bg-white p-8 rounded shadow-md w-96 text-center">

    <h1 class="text-2xl font-semibold mb-4">Login</h1>

    <div>
        <label for="verificationCode" class="block text-sm font-medium text-gray-600">Verification Code</label>
        <input type="text" id="verificationCode" name="verificationCode" pattern="[0-9]{6}" maxlength="6"
               class="mt-1 p-2 w-full rounded-md focus:outline-none focus:ring focus:border-blue-300 border
                   border-gray-600 border-2"
               placeholder="Enter 6-digit code" required>
        <p class="text-xs text-gray-500 mt-1">Please enter a 6-digit code.</p>
    </div>

    <button class="bg-cyan-700 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-full focus:outline-none
                focus:shadow-outline-blue active:bg-blue-800 mt-5">
        Login
    </button>

</div>

</body>

</html>
