@extends("layout", ["title" => "Delete file"])

<body class="flex items-center justify-center h-screen bg-gradient-to-r from-violet-300 to-cyan-950">

@section("content")
    <div class="bg-white p-8 rounded shadow-md w-96 text-center">

        <h1 class="text-2xl font-semibold mb-4 text-red-600">Confirm delete</h1>

        <form action="{{route("file.delete", $id)}}" method="post">
            @csrf
            @method("delete")

            @include("forms.code")

            <button class="bg-red-700 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-full focus:outline-none
                        focus:shadow-outline-blue active:bg-blue-800 mt-5">
                Delete
            </button>

        </form>
    </div>
@endsection

</body>
