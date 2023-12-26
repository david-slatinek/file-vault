@extends("layout", ["title" => "File Vault"])

@section("content")
    <body class="flex flex-col items-center justify-center h-screen bg-gradient-to-r from-slate-400 to-sky-400">

    @if (Session::has("error"))
        <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4">
            <h3 class="block sm:inline">{{ Session::get("error") }}</h3>
        </div>
    @endif

    <div class="bg-white p-8 rounded shadow-md w-96">
        <h1 class="text-3xl font-semibold mb-6 text-center text-gray-800">Login with Github</h1>

        <div class="flex items-center justify-center">
            <a href="/github/redirect">
                <button class="bg-gradient-to-r from-blue-500 to-red-400 hover:from-blue-600 hover:to-teal-500 text-white
            font-bold py-3 px-10 rounded-full focus:outline-none focus:shadow-outline-blue active:from-blue-800
            active:to-teal-700">
                    Login
                </button>
            </a>
        </div>
    </div>

    </body>

@endsection
