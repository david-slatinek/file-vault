@extends("layout", ["title" => "Download"])

<body class="flex items-center justify-center h-screen bg-gradient-to-r from-zinc-400 to-emerald-700">

<div class="p-8 bg-white rounded shadow-md container mx-auto max-w-lg">
    <h1 class="text-2xl font-semibold mb-4">File Download</h1>

    <form action="{{route("file.download", $id)}}" method="get">
        @csrf
        @method("get")

        @include("forms.password")
        @include("forms.code")

        <button type="submit" class="bg-blue-700 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded-full
                focus:outline-none focus:shadow-outline-blue active:bg-blue-800">
            Download
        </button>
    </form>
</div>

</body>
