@extends("layout", ["title" => "Home"])

<body class="flex flex-col text-center mt-20 bg-gradient-to-r from-slate-200 to-red-700 font-sans">

@section("content")
    <div class="max-w-full mx-auto p-20 bg-white border border-gray-300 rounded-md shadow-md">
        <h1 class="text-red-600 text-2xl font-bold mb-4">Oops! Something went wrong.</h1>
        <p class="text-gray-700">We're sorry, but an error occurred while processing your request.</p>
        <p class="text-gray-700">Please try again later.</p>

        @if (session("error"))
            <div class="alert alert-danger">
                <p class="text-gray-700">Error: {{session("error")}}</p>
            </div>
        @endif

    </div>
@endsection

</body>
