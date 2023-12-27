@extends("layout", ["title" => "Delete file"])

@section("content")
    <body class="flex items-center justify-center h-screen bg-gradient-to-r from-violet-300 to-cyan-950">

    <div class="bg-white p-8 rounded shadow-md w-96 text-center">

        <h1 class="text-2xl font-semibold mb-4 text-red-600">Confirm delete</h1>

        <form action="{{route("file.delete", $id)}}" method="post">
            @csrf
            @method("delete")

            <div>
                <label for="code" class="block text-sm font-medium text-gray-600">Verification Code</label>
                <input type="text" id="code" name="code" pattern="[0-9]{6}" maxlength="6"
                       class="mt-1 p-2 w-full rounded-md focus:outline-none focus:ring focus:border-blue-300 border
                   border-gray-600 border-2"
                       placeholder="Enter 6-digit code" required>
                <p class="text-xs text-gray-500 mt-1">Please enter a 6-digit code.</p>
            </div>

            <button class="bg-red-700 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-full focus:outline-none
                focus:shadow-outline-blue active:bg-blue-800 mt-5">
                Delete
            </button>

        </form>
    </div>

    </body>

@endsection
