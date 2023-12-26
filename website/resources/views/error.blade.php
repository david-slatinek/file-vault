@extends("layout", ["title" => "Home"])

@section("content")
    <body class="text-center p-10 bg-gray-100 font-sans">

    <div class="max-w-md mx-auto p-6 bg-white border border-gray-300 rounded-md shadow-md">
        <h1 class="text-red-600 text-2xl font-bold mb-4">Oops! Something went wrong.</h1>
        <p class="text-gray-700">We're sorry, but an error occurred while processing your request.</p>
        <p class="text-gray-700">Please try again later.</p>
        <p class="text-gray-700">Error: {{Session::get("error")}}</p>
    </div>

    </body>
@endsection
