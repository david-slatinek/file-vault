@extends("layout", ["title" => "Upload"])

<body class="flex items-center justify-center h-screen bg-gradient-to-r from-blue-300 to-lime-800">

@section("content")
    <div class="p-8 rounded shadow-md container mx-auto max-w-lg bg-white">
        <h1 class="text-2xl font-semibold mb-4">File Upload</h1>

        <form action="{{route("file.upload")}}" method="post" enctype="multipart/form-data">
            @csrf
            @method("post")

            <div class="mb-5">
                <label for="file" class="block text-sm font-medium text-gray-600">Choose a file</label>
                <input type="file" id="file" name="file" class="mt-1 p-2 w-full rounded-md focus:outline-none focus:ring
                        focus:border-blue-300" required>
            </div>

            @include("forms.password")
            @include("forms.code")

            <button type="submit" class="bg-blue-700 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-full
                    focus:outline-none focus:shadow-outline-blue active:bg-blue-800">
                Upload
            </button>
        </form>
    </div>
@endsection

</body>
