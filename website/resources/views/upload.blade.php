@extends("layout", ["title" => "Upload"])

@section("content")
    <body class="flex items-center justify-center h-screen bg-gradient-to-r from-yellow-100 to-lime-800">

    <div class="p-8 bg-white rounded shadow-md container mx-auto max-w-lg">
        <h1 class="text-2xl font-semibold mb-6">File Upload</h1>

        <div>
            <label for="file" class="block text-sm font-medium text-gray-600">Choose a file</label>
            <input type="file" id="file" name="file" class="mt-1 p-2 w-full rounded-md focus:outline-none focus:ring
                focus:border-blue-300">
        </div>

        <button class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-full focus:outline-none
            focus:shadow-outline-blue active:bg-blue-800 mt-7">
            Upload File
        </button>
    </div>

    </body>

@endsection
