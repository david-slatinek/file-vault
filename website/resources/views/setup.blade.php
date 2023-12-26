<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <script src="https://cdn.tailwindcss.com"></script>

    <title>2FA Setup</title>
</head>

<body class="flex items-center justify-center h-screen bg-gray-100 bg-gradient-to-r from-cyan-300 to-blue-600">

<div class="bg-white p-8 rounded shadow-md w-96 text-center">

    <img src="image.jpg" alt="QR code" class="mb-4 mx-auto rounded-full">

    <div>
        <h1 class="text-2xl font-semibold mb-6 text-center text-gray-800">asfrtn</h1>
    </div>

    <div>
        <p class="text-gray-700 mb-4">Please scan the QR code with your 2FA app or enter the code manually.</p>
        <p class="text-gray-700 mb-4">As a backup, we recommend you save these codes in a password manager.</p>
    </div>

    <div class="flex justify-between items-center">
        <div></div>
        <button
            class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-full focus:outline-none
            focus:shadow-outline-blue active:bg-blue-800">
            Next
        </button>
    </div>

</div>

</body>

</html>
