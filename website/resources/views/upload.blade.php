@extends("layout", ["title" => "Upload"])

<body class="flex items-center justify-center h-screen bg-gradient-to-r from-yellow-100 to-lime-800">

@section("content")
    <div class="p-8 bg-white rounded shadow-md container mx-auto max-w-lg">
        <h1 class="text-2xl font-semibold mb-4">File Upload</h1>

        <form action="{{route("file.upload")}}" method="post" enctype="multipart/form-data">
            @csrf
            @method("post")

            <div class="mb-5">
                <label for="file" class="block text-sm font-medium text-gray-600">Choose a file</label>
                <input type="file" id="file" name="file" class="mt-1 p-2 w-full rounded-md focus:outline-none focus:ring
                        focus:border-blue-300" required>
            </div>

            <div class="max-w-md mb-5">
                <label for="password" class="block text-gray-700 text-sm font-bold mb-2">Password</label>
                <input type="password" id="password" name="password" placeholder="Enter your password"
                       class="mt-1 p-2 w-full rounded-md focus:outline-none focus:ring focus:border-blue-300 border
                           border-gray-600 border-2" minlength="12" required>
                <p class="text-xs text-gray-500 mt-1">Must be at least 12 characters, at least 1 uppercase letter,
                    1 lowercase letter, 1 digit, and 1 special character.</p>
                <label for="show">Show password</label>
                <input id="show" name="show" type="checkbox" onclick="(() => {
                    let x = document.getElementById('password');
                    if (x.type === 'password') {
                        x.type = 'text';
                    } else {
                        x.type = 'password';
                    }
                })();">
            </div>

            <div class="mb-5">
                <label for="code" class="block text-gray-700 text-sm font-bold mb-2">Verification Code</label>
                <input type="text" id="code" name="code" pattern="[0-9]{6}" maxlength="6" minlength="6"
                       class="mt-1 p-2 w-full rounded-md focus:outline-none focus:ring focus:border-blue-300 border
                            border-gray-600 border-2"
                       placeholder="Enter 6-digit code" required>
                <p class="text-xs text-gray-500 mt-1">Please enter a 6-digit code.</p>
            </div>

            <button type="submit" class="bg-blue-700 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-full focus:outline-none
                        focus:shadow-outline-blue active:bg-blue-800">
                Upload
            </button>

        </form>
    </div>
@endsection

</body>
