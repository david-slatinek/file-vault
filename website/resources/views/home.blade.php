@extends("layout", ["title" => "Home"])

@section("content")
    <body class="h-screen bg-gradient-to-r from-zinc-400 to-stone-500">

    <div class="container mx-auto mt-8 p-8 bg-white rounded shadow-md">

        {{--        <div class="max-w-md mb-5">--}}
        {{--            <label for="code" class="block text-gray-700 text-sm font-bold mb-2">Verification--}}
        {{--                Code</label>--}}
        {{--            <input form="myform" type="text" id="code" name="code" pattern="[0-9]{6}" maxlength="6"--}}
        {{--                   class="mt-1 p-2 w-full rounded-md focus:outline-none focus:ring focus:border-blue-300 border--}}
        {{--                   border-gray-600 border-2"--}}
        {{--                   placeholder="Enter 6-digit code" required>--}}
        {{--            <p class="text-xs text-gray-500 mt-1">Please enter a 6-digit code.</p>--}}
        {{--        </div>--}}

        {{--        <div class="max-w-md mb-10">--}}
        {{--            <label for="password" class="block text-gray-700 text-sm font-bold mb-2">Password</label>--}}
        {{--            <input form="myform" type="password" id="password" name="password" placeholder="Enter your password"--}}
        {{--                   class="mt-1 p-2 w-full rounded-md focus:outline-none focus:ring focus:border-blue-300 border--}}
        {{--                   border-gray-600 border-2" minlength="12"--}}
        {{--                   pattern="(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+])[A-Za-z\d!@#$%^&*()_+]{8,}">--}}
        {{--            <p class="text-xs text-gray-500 mt-1">Must be at least 12 characters, at least 1 uppercase letter,--}}
        {{--                1 lowercase letter, 1 digit, and 1 special character.</p>--}}
        {{--        </div>--}}

        <h1 class="text-2xl font-semibold mb-6">File Management</h1>

        <table class="min-w-full border border-gray-300">
            <thead>
            <tr>
                <th class="py-2 px-4 border-b">ID</th>
                <th class="py-2 px-4 border-b">Filename</th>
                <th class="py-2 px-4 border-b">Created</th>
                <th class="py-2 px-4 border-b">Accessed</th>
                <th class="py-2 px-4 border-b">Action</th>
            </tr>
            </thead>
            <tbody>

            @foreach($files as $file)
                <tr class="text-center hover:bg-gray-300 even:bg-amber-100 odd:bg-sky-100">
                    <td class="py-2 px-4 border-b">{{$loop->iteration}}</td>
                    <td class="py-2 px-4 border-b">{{$file->filename}}</td>
                    <td class="py-2 px-4 border-b">{{$file->createdAt}}</td>
                    <td class="py-2 px-4 border-b">{{$file->accessedAt}}</td>
                    <td class="py-2 px-4 border-b">
                        <button class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-1 px-2 rounded-full
                            focus:outline-none focus:shadow-outline-blue active:bg-blue-800">
                            Download
                        </button>

                        <form action="{{route("code", $file->id)}}" method="post">
                            @csrf
                            @method("delete")
                            <button class="bg-red-500 hover:bg-red-600 text-white font-bold py-1 px-2 rounded-full
                            focus:outline-none focus:shadow-outline-red active:bg-red-800">
                                Delete
                            </button>
                        </form>

{{--                        <a href="{{route("code", $file->id)}}">--}}
{{--                            <button class="bg-red-500 hover:bg-red-600 text-white font-bold py-1 px-2 rounded-full--}}
{{--                            focus:outline-none focus:shadow-outline-red active:bg-red-800">--}}
{{--                                Delete--}}
{{--                            </button>--}}
{{--                        </a>--}}
                    </td>
                </tr>

            @endforeach

            </tbody>
        </table>

    </div>

    </body>
@endsection
